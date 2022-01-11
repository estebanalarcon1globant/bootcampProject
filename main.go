package main

import (
	"bootcampProject/database"
	pb "bootcampProject/grpc"
	"bootcampProject/logging"
	"bootcampProject/users/implementation"
	"bootcampProject/users/repository"
	"bootcampProject/users/transport"
	transportGrpc "bootcampProject/users/transport/grpc"
	transportHttp "bootcampProject/users/transport/http"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "order",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	sqlDB, err := database.SetupDB()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(sqlDB, logger)
	userSvc := implementation.NewUserService(userRepo, logger)
	middleware := logging.NewMiddleware(logger, userSvc)
	userSvc = middleware
	endpoints := transport.MakeEndpoints(userSvc)
	httpServer := transportHttp.NewUserHandler(endpoints, logger)
	grpcServer := transportGrpc.NewUserGRPCServer(endpoints, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:    *httpAddr,
			Handler: httpServer,
		}
		errs <- server.ListenAndServe()
	}()

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterUserServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)

}
