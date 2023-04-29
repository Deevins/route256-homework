package main

import (
	"context"
	"github.com/spf13/viper"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	postgresqlRepository "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/postgresql"
	service "gitlab.ozon.dev/daker255/homework-8/internal/app/services"
	pb "gitlab.ozon.dev/daker255/homework-8/internal/pb/server"
	database "gitlab.ozon.dev/daker255/homework-8/pkg/database/clients"
	jaegerExporter "go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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

	if err := initConfig(); err != nil {
		log.Fatalf("can not read config file %s", err.Error())
	}

	if err := TracerProvider("http://localhost:14268/api/traces", "GRPC-client"); err != nil {
		log.Fatal(err)
	}
	defer tracer.Shutdown(ctx)

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
	userRepo := postgresqlRepository.NewPostgresqlUserRepo(db)
	orderRepo := postgresqlRepository.NewPostgresqlOrderRepo(db)

	//init services
	userService := service.NewUserService(userRepo)
	orderService := service.NewOrderService(userRepo, orderRepo)

	grpcServerPort := viper.GetString("grpc_client.port")
	// Create a TCP connection
	lsn, err := net.Listen("tcp", grpcServerPort)
	if err != nil {
		log.Fatal(err)
	}
	// Create the GRPC server
	grpcServer := grpc.NewServer()

	// Allows us to use a 'list' call to list all available APIs
	reflection.Register(grpcServer)

	// We register an objects that should implement all the described APIs
	pb.RegisterUserServiceV1Server(grpcServer, NewUserClientImplementation(userService))
	pb.RegisterOrderServiceV1Server(grpcServer, NewOrderClientImplementation(orderService))

	// Serving the GRPC server on a created TCP socket
	log.Printf("gRPC server successfully started on port %s", lsn.Addr().String())
	if err := grpcServer.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
