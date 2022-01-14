package transport

import (
	"bootcampProject/users/domain"
	"context"
	"github.com/go-kit/kit/endpoint"
)

// UserEndpointsGRPC holds all Go kit endpoints for the User service.
type UserEndpointsGRPC struct {
	CreateUser endpoint.Endpoint
}

// MakeEndpointsGRPC initializes all Go kit endpoints for the Order service.
func MakeEndpointsGRPC(s domain.UserService) UserEndpointsGRPC {
	return UserEndpointsGRPC{
		CreateUser: makeCreateUserGRPCEndpoint(s),
	}
}

func makeCreateUserGRPCEndpoint(s domain.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//TODO: Handle type assertion's error
		req := request.(CreateUserRequest)
		id, err := s.CreateUser(ctx, req.User)
		return CreateUserResponse{ID: id, Err: err}, err
	}
}
