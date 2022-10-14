package transport

import (
	"bootcampProject/users/domain"
	"bootcampProject/utils"
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
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

/*
func makeCreateAdditionalInfo(s domain.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		message, err := s.CreateAdditionalInfo(ctx)
		if err != nil {
			return CreateUserResponse{Err: err}, err
		}
	}
}*/

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
		return CreateUserResponse{Err: utils.NewErrBadRequest()}, utils.NewErrBadRequest()
	}
}

func makeGetUsersGRPCEndpoint(s domain.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(GetUsersRequest); ok {
			req.SetDefault()
			users, err := s.GetUsers(ctx, req.limit, req.offset)
			return GetUsersResponse{Users: users, Err: err}, err
		}
		return CreateUserResponse{Err: utils.NewErrBadRequest()}, utils.NewErrBadRequest()
	}
}

func makeAuthenticateGRPCEndpoint(s domain.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var authResp AuthResponse
		if req, ok := request.(domain.Auth); ok {
			token, err := s.Authenticate(ctx, req)
			if err != nil {
				authResp.Err = err.Error()
				setStatus(ctx, err)
			}
			authResp.Token = token
			return authResp, nil
		}
		setStatus(ctx, utils.NewErrBadRequest())
		return AuthResponse{Err: utils.NewErrBadRequest().Error()}, nil
	}
}

func MakeHTTPResponseModifier() func(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	return func(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
		md, ok := runtime.ServerMetadataFromContext(ctx)
		if !ok {
			return nil
		}

		vals := md.HeaderMD.Get("x-http-code")

		if len(vals) > 0 {
			code, err := strconv.Atoi(vals[0])
			if err != nil {
				return err
			}

			delete(md.HeaderMD, "x-http-code")
			delete(w.Header(), "Grpc-Metadata-X-Http-Code")
			w.WriteHeader(code)
		}
		return nil
	}
}

func setStatus(ctx context.Context, e error) {

	switch e.(type) {
	case *utils.ErrBadRequest:
		grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "400"))
	case *utils.ErrInvalidCredentials:
		grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "403"))
	default:
		grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "500"))
	}
}
