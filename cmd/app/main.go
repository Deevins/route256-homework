package main

import (
	"context"
	"github.com/spf13/viper"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/cli_commands"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/handlers"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/repository"
	postgresqlRepository "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/postgresql"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/server"
	service "gitlab.ozon.dev/daker255/homework-8/internal/app/services"
	"gitlab.ozon.dev/daker255/homework-8/internal/pb"
	database "gitlab.ozon.dev/daker255/homework-8/pkg/database/clients"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := initConfig(); err != nil {
		log.Fatalf("can not read config file %s", err.Error())
	}

	httpServerPort := viper.GetString("server.port")

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

	// init handlers for our end-points
	coreHandler := handlers.NewRootHandler(ctx, coreService)

	lsn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceV1Server(grpcServer, NewUserImplementation(db))
	pb.RegisterOrderServiceV1Server(grpcServer, NewOrderImplementation(db))

	go func() {
		log.Printf("starting gRPC server on %s", lsn.Addr().String())
		if err := grpcServer.Serve(lsn); err != nil {
			log.Fatal(err)
		}
	}()

	// start http server
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
