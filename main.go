package main

import (
	"bootcampProject/config"
	"bootcampProject/database"
	pb "bootcampProject/proto"
	"bootcampProject/users/domain"
	"bootcampProject/users/repository"
	"bootcampProject/users/service"
	"bootcampProject/users/transport"
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultGrpcGatewayPort = ":8090"
	defaultGrpcPort        = ":8080"
)

func main() {

	var (
		grpcGwAddr = flag.String("grpc.gw.addr", defaultGrpcGatewayPort, "gRPC Gateway listen address")
		grpcAddr   = flag.String("grpc.addr", defaultGrpcPort, "gRPC listen address")
	)

	//LOGGER
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "user",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	//LOAD .ENV CONFIGURATION
	err := config.LoadConfiguration()
	if err != nil {
		level.Info(logger).Log("during", "Load .env", "err", err)
		os.Exit(1)
	}

	//SET SQL DATABASE
	err = database.SetupDB()
	if err != nil {
		level.Info(logger).Log("during", "Setup DB", "err", err)
		os.Exit(1)
	}

	sqlDB := database.GetConnection()
	userRepo := repository.NewUserRepository(sqlDB)

	var userSvc domain.UserService

	addInfoClient, err := service.NewAdditionalInformationClient("localhost:9080")
	if err != nil {
		level.Info(logger).Log("during", "creating Additional Info Client", "err", err)
		os.Exit(1)
	}
	defer addInfoClient.Close()

	tokenGenerator := service.NewTokenGenerator()
	userSvc = service.NewUserService(userRepo, tokenGenerator, addInfoClient)
	userSvc = service.NewUserServiceLogging(log.With(logger, "component", "users"), userSvc)
	//create additionalInfo client

	//GRPC SERVER
	endpointsGRPC := transport.MakeEndpointsGRPC(userSvc)
	grpcServer := transport.NewUserGRPCServer(endpointsGRPC, logger)

	errs := make(chan error)

	// Create a listener on TCP port
	listener, err := net.Listen("tcp", defaultGrpcPort)
	if err != nil {
		level.Info(logger).Log("during", "Listen tcp", "err", err)
		os.Exit(1)
	}

	baseServer := grpc.NewServer()
	pb.RegisterUserServiceServer(baseServer, grpcServer)
	// Serve gRPC server

	go func() {
		level.Info(logger).Log("transport", "gRPC server", "addr", *grpcAddr)
		errs <- baseServer.Serve(listener)
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0"+defaultGrpcPort,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		level.Info(logger).Log("during", "Dial context", "err", err)
		os.Exit(1)
	}

	responseModifier := transport.MakeHTTPResponseModifier()
	gwMux := runtime.NewServeMux(
		runtime.WithForwardResponseOption(responseModifier),
	)
	// Register User Handler
	err = pb.RegisterUserServiceHandler(context.Background(), gwMux, conn)
	if err != nil {
		level.Info(logger).Log("during", "Setup RegisterUserServiceHandler", "err", err)
		os.Exit(1)
	}

	gwServer := &http.Server{
		Addr:    *grpcGwAddr,
		Handler: gwMux,
	}

	go func() {
		level.Info(logger).Log("transport", "gRPC Gateway", "addr", *grpcGwAddr)
		errs <- gwServer.ListenAndServe()
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	level.Error(logger).Log("exit", <-errs)
}

//Comment only for rebase
