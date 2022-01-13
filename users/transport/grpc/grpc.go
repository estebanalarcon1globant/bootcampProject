package grpc

import (
	pb "bootcampProject/grpc"
	"bootcampProject/users/domain"
	"bootcampProject/users/transport"
	"context"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	createUser gt.Handler
}

// NewUserGRPCServer initializes a new gRPC server
func NewUserGRPCServer(svcEndpoints transport.UserEndpointsGRPC, logger log.Logger) pb.UserServiceServer {
	return &gRPCServer{
		createUser: gt.NewServer(
			svcEndpoints.CreateUser,
			decodeCreateUserRequest,
			encodeCreateUserResponse,
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

func decodeCreateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.NewUser)
	return transport.CreateUserRequest{User: domain.Users{
		PwdHash: req.GetPwdHash(),
		Name:    req.GetName(),
		Age:     int(req.GetAge()),
	}}, nil
}

func encodeCreateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(transport.CreateUserResponse)
	return &pb.User{
		Id: int32(resp.ID),
	}, nil
}
