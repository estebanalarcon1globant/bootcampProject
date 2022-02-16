package proto

import (
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

func NewGrpcClient() UserServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return NewUserServiceClient(conn)
}
