package main

import (
	"bootcampProject/additional_information/internal/repository"
	"bootcampProject/additional_information/internal/service"
	"bootcampProject/additional_information/internal/transport"
	pb "bootcampProject/additional_information/proto"
	"bootcampProject/database"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultGrpcPort = ":9080"
)

func main() {
	var (
		grpcAddr = flag.String("grpc.addr", defaultGrpcPort, "gRPC listen address")
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

	err := database.SetupNoSqlDB()
	if err != nil {
		level.Info(logger).Log("during", "Setup DB", "err", err)
		os.Exit(1)
	}

	noSqlDB := database.GetNoSQLConnection()
	additionalInfoRepo := repository.NewUserAdditionalInfoRepo(noSqlDB)

	userAdditionalInfo := service.NewUserAdditionalInfoService(additionalInfoRepo)

	//GRPC Server
	endpointsGrpc := transport.MakeEndpointsGRPC(userAdditionalInfo)
	grpcServer := transport.NewAdditionalInfoGRPCServer(endpointsGrpc)

	errs := make(chan error)

	// Create a listener on TCP port
	listener, err := net.Listen("tcp", defaultGrpcPort)
	if err != nil {
		level.Info(logger).Log("during", "Listen tcp", "err", err)
		os.Exit(1)
	}

	baseServer := grpc.NewServer()
	pb.RegisterAdditionalInformationServiceServer(baseServer, grpcServer)
	// Serve gRPC server

	go func() {
		level.Info(logger).Log("transport", "gRPC server", "addr", *grpcAddr)
		errs <- baseServer.Serve(listener)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	level.Error(logger).Log("exit", <-errs)
}
