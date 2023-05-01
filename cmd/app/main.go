package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/cli_commands"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/handlers"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/repository"
	postgresqlRepository "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/postgresql"
	server "gitlab.ozon.dev/daker255/homework-8/internal/app/server/http"
	service "gitlab.ozon.dev/daker255/homework-8/internal/app/services"
	"gitlab.ozon.dev/daker255/homework-8/internal/metrics"
	"gitlab.ozon.dev/daker255/homework-8/internal/pb"
	database "gitlab.ozon.dev/daker255/homework-8/pkg/database/clients"
	jaegerExporter "go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"google.golang.org/grpc"
)

var (
	// Store a global trace provider variable to clear it before closing
	tracer *tracesdk.TracerProvider
)

func TracerProvider(url, name string) error {
	// Create the Jaeger exporter
	exp, err := jaegerExporter.New(jaegerExporter.WithCollectorEndpoint(jaegerExporter.WithEndpoint(url)))
	if err != nil {
		return err
	}
	tracer = tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(name),
		)),
	)
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	prometheus.MustRegister(metrics.UserCreateCounter)
	prometheus.MustRegister(metrics.UserDeleteCounter)
	prometheus.MustRegister(metrics.RequestGauge)

	if err := initConfig(); err != nil {
		log.Fatalf("can not read config file %s", err.Error())
	}

	err := TracerProvider("http://localhost:14268/api/traces", "Core-GRPC-server")
	if err != nil {
		log.Fatal(err)
	}
	defer func(tracer *tracesdk.TracerProvider, ctx context.Context) {
		err := tracer.Shutdown(ctx)
		if err != nil {
			log.Fatalf("graceful shutdown Jaeger not successful, error: %s", err)
		}
	}(tracer, context.Background())

	db, err := database.NewDB(ctx, models.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Dbname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("can not connect to database intance in main package in method clients.NewDB: %s", err)
	}

	defer db.GetPool(ctx).Close()

	// init repos
	coreRepo := repository.NewPostgresqlRepository(db)
	userRepo := postgresqlRepository.NewPostgresqlUserRepo(db)
	orderRepo := postgresqlRepository.NewPostgresqlOrderRepo(db)

	//init services
	coreService := service.NewCoreService(coreRepo)
	userService := service.NewUserService(userRepo)
	orderService := service.NewOrderService(userRepo, orderRepo)

	// init CLI commands and CLI handler
	userCommand := cli_commands.NewUserCommand(ctx, userService)
	orderCommand := cli_commands.NewOrderCommand(ctx, orderService)
	spellCommand := cli_commands.NewSpellCommand()

	cliHandler := NewCLIHandler()
	//
	cliHandler.Register("order", orderCommand)
	cliHandler.Register("user", userCommand)
	cliHandler.Register("spell", spellCommand)

	// init	 handlers for our end-points
	coreHandler := handlers.NewRootHandler(ctx, coreService)

	grpcServerPort := viper.GetString("grpc_server.port")
	lsn, err := net.Listen("tcp", grpcServerPort)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceV1Server(grpcServer, NewUserImplementation(userService))
	pb.RegisterOrderServiceV1Server(grpcServer, NewOrderImplementation(orderService))

	go func() {
		log.Printf("gRPC server successfully started on port %s", lsn.Addr().String())
		if err := grpcServer.Serve(lsn); err != nil {
			log.Fatal(err)
		}
	}()

	// HTTP exporter для prometheus
	go func() {
		log.Printf("Prometheus server successfully started on port %s", ":9091")
		err := http.ListenAndServe(":9091", promhttp.Handler())
		if err != nil {
			fmt.Println(err)
			log.Fatalf("error while handling Prometheus http server")
		}
	}()

	// start http server
	httpServerPort := viper.GetString("http_server.port")
	srv := server.NewServer()
	if err := srv.Run(httpServerPort, coreHandler.InitRoutes()); err != nil {
		log.Fatalf("could not start server on port %s with err %s", httpServerPort, err)
	}

	//initialize endless loop to wait for commands
	//for {
	//	go cliHandler.CommandProcessing()
	//}
}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
