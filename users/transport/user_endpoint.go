package transport

import (
	"bootcampProject/users/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

// UserEndpoints holds all Go kit endpoints for the User service.
type UserEndpoints struct {
	CreateUser endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the Order service.
func MakeEndpoints(s service.UserService) UserEndpoints {
	return UserEndpoints{
		CreateUser: makeCreateUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//TODO: Handle type assertion's error
		req := request.(CreateUserRequest)
		id, err := s.CreateUser(ctx, req.User)
		return CreateUserResponse{ID: id, Err: err}, nil
	}
}
