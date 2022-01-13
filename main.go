package main

import (
	"bootcampProject/database"
	pb "bootcampProject/grpc"
	"bootcampProject/users/repository"
	"bootcampProject/users/service"
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

	err := database.SetupDB()
	if err != nil {
		panic(err)
	}
	sqlDB := database.GetConnection()
	userRepo := repository.NewUserRepository(sqlDB, logger)
	userSvc := service.NewUserService(userRepo, logger)

	//GRPC SERVER
	endpointsGRPC := transport.MakeEndpointsGRPC(userSvc)
	grpcServer := transportGrpc.NewUserGRPCServer(endpointsGRPC, logger)

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

	//HTTP SERVER
	//middleware := logging.NewMiddleware(logger, userSvc)
	//userSvc = middleware
	grpcClient := pb.NewGrpcClient()
	endpointsHTTP := transport.MakeEndpointsHTTP(grpcClient)
	httpServer := transportHttp.NewUserHTTPServer(endpointsHTTP, logger)

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

	level.Error(logger).Log("exit", <-errs)

}
