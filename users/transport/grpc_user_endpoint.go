package transport

import (
	"bootcampProject/users/domain"
	"bootcampProject/utils"
	"context"
	"github.com/go-kit/kit/endpoint"
)

// UserEndpointsGRPC holds all Go kit endpoints for the User service.
type UserEndpointsGRPC struct {
	CreateUser   endpoint.Endpoint
	GetUsers     endpoint.Endpoint
	Authenticate endpoint.Endpoint
}

// MakeEndpointsGRPC initializes all Go kit endpoints for the Order service.
func MakeEndpointsGRPC(s domain.UserService) UserEndpointsGRPC {
	return UserEndpointsGRPC{
		CreateUser:   makeCreateUserGRPCEndpoint(s),
		GetUsers:     makeGetUsersGRPCEndpoint(s),
		Authenticate: makeAuthenticateGRPCEndpoint(s),
	}
}

func makeCreateUserGRPCEndpoint(s domain.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(CreateUserRequest); ok {
			err := req.Validate()
			if err != nil {
				return CreateUserResponse{Err: err}, err
			}
			id, err := s.CreateUser(ctx, req.User)
			return CreateUserResponse{ID: id, Err: err}, err
		}
		return CreateUserResponse{Err: utils.ErrBadRequest}, utils.ErrBadRequest
	}
}

func makeGetUsersGRPCEndpoint(s domain.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(GetUsersRequest); ok {
			req.SetDefault()
			users, err := s.GetUsers(ctx, req.limit, req.offset)
			return GetUsersResponse{Users: users, Err: err}, err
		}
		return CreateUserResponse{Err: utils.ErrBadRequest}, utils.ErrBadRequest
	}
}

func makeAuthenticateGRPCEndpoint(s domain.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(domain.Auth); ok {
			token, err := s.Authenticate(ctx, req)
			return AuthResponse{
				Token: token,
			}, err
		}
		return AuthResponse{}, utils.ErrBadRequest
	}
}
