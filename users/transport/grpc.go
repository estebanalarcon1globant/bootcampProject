package transport

import (
	pb "bootcampProject/grpc"
	"bootcampProject/users/domain"
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	createUser gt.Handler
}

var (
	ErrBadRequest = errors.New("bad request. Doesn't match interface")
)

// NewUserGRPCServer initializes a new gRPC server
func NewUserGRPCServer(svcEndpoints UserEndpointsGRPC, _ log.Logger) pb.UserServiceServer {
	return &gRPCServer{
		createUser: gt.NewServer(
			svcEndpoints.CreateUser,
			decodeCreateUserGRPCRequest,
			encodeCreateUserGRPCResponse,
		),
	}
}

func (s *gRPCServer) CreateUser(ctx context.Context, req *pb.NewUser) (*pb.User, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.User), nil
}

func decodeCreateUserGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	if req, ok := request.(*pb.NewUser); ok {
		return CreateUserRequest{User: domain.Users{
			PwdHash: req.GetPwdHash(),
			Name:    req.GetName(),
			Age:     int(req.GetAge()),
		}}, nil
	}
	return CreateUserRequest{}, ErrBadRequest
}

func encodeCreateUserGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	if resp, ok := response.(CreateUserResponse); ok {
		return &pb.User{
			Id: int32(resp.ID),
		}, nil
	}
	return &pb.User{}, ErrBadRequest
}
