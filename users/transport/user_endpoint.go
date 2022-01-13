package transport

import (
	pb "bootcampProject/grpc"
	"bootcampProject/users/domain"
	"context"
	"github.com/go-kit/kit/endpoint"
)

// UserEndpointsGRPC holds all Go kit endpoints for the User service.
type UserEndpointsGRPC struct {
	CreateUser endpoint.Endpoint
}

type UserEndpointsHTTP struct {
	CreateUser endpoint.Endpoint
}

// MakeEndpointsGRPC initializes all Go kit endpoints for the Order service.
func MakeEndpointsGRPC(s domain.UserService) UserEndpointsGRPC {
	return UserEndpointsGRPC{
		CreateUser: makeCreateUserEndpoint(s),
	}
}

// MakeEndpointsHTTP initializes all Go kit endpoints for the Order service.
func MakeEndpointsHTTP(grpcClient pb.UserServiceClient) UserEndpointsHTTP {
	return UserEndpointsHTTP{
		CreateUser: makeCreateUserHTTPEndpoint(grpcClient),
	}
}

func makeCreateUserEndpoint(s domain.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//TODO: Handle type assertion's error
		req := request.(CreateUserRequest)
		id, err := s.CreateUser(ctx, req.User)
		return CreateUserResponse{ID: id, Err: err}, nil
	}
}

func makeCreateUserHTTPEndpoint(grpcClient pb.UserServiceClient) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//TODO: Handle type assertion's error
		req := request.(CreateUserRequest)
		r, err := grpcClient.CreateUser(ctx, &pb.NewUser{
			PwdHash: req.User.PwdHash,
			Name:    req.User.Name,
			Age:     int32(req.User.Age),
		})
		return CreateUserResponse{ID: int(r.Id), Err: err}, nil
	}
}
