package main

import (
	"bootcampProject/config"
	"bootcampProject/database"
	pb "bootcampProject/grpc"
	"bootcampProject/grpc/server"
	"bootcampProject/users/domain"
	"bootcampProject/users/repository"
	"bootcampProject/users/service"
	"bootcampProject/users/transport"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultHttpPort = ":8080"
	defaultGrpcPort = ":50051"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", defaultHttpPort, "HTTP listen address")
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
	userSvc = service.NewUserService(userRepo)
	userSvc = service.NewUserServiceLogging(log.With(logger, "component", "users"), userSvc)

	//GRPC SERVER
	endpointsGRPC := transport.MakeEndpointsGRPC(userSvc)
	grpcServer := transport.NewUserGRPCServer(endpointsGRPC, logger)

	go func() {
		userGrpcServer := server.NewUserServer()
		if err = userGrpcServer.Run(grpcServer, defaultGrpcPort); err != nil {
			logger.Log("during", "gRPC serve", "err", err)
			os.Exit(1)
		}
		level.Info(logger).Log("transport", "gRPC", "addr", *grpcAddr)
	}()

	//HTTP SERVER
	grpcClient := pb.NewGrpcClient()
	endpointsHTTP := transport.MakeEndpointsHTTP(grpcClient)
	httpServer := transport.NewUserHTTPServer(endpointsHTTP, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		httpServer := &http.Server{
			Addr:    *httpAddr,
			Handler: httpServer,
		}
		errs <- httpServer.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)
}
