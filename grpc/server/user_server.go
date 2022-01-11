package server

import (
	pb "bootcampProject/grpc"
	"context"
	"log"
)

const (
	port = ":50051"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.User{
		Name: in.GetName(), Age: in.GetAge(), Id: 11,
	}, nil
}
