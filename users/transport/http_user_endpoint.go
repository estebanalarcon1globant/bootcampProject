package transport

import (
	pb "bootcampProject/grpc"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type UserEndpointsHTTP struct {
	CreateUser endpoint.Endpoint
}

// MakeEndpointsHTTP initializes all Go kit endpoints for the Order service.
func MakeEndpointsHTTP(grpcClient pb.UserServiceClient) UserEndpointsHTTP {
	return UserEndpointsHTTP{
		CreateUser: makeCreateUserHTTPEndpoint(grpcClient),
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
		return CreateUserResponse{ID: int(r.Id), Err: err}, err
	}
}
