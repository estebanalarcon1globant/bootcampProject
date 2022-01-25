package server

import (
	pb "bootcampProject/grpc"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	userList *pb.UserList
}

func NewUserServer() *UserServer {
	return &UserServer{
		userList: &pb.UserList{},
	}
}

func (s *UserServer) Run(grpcServer pb.UserServiceServer, port string) error {
	grpcListener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	baseServer := grpc.NewServer()
	pb.RegisterUserServiceServer(baseServer, grpcServer)
	return baseServer.Serve(grpcListener)
}

func (s *UserServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.User{
		Name: in.GetName(), Age: in.GetAge(), Id: 11,
	}, nil
}
