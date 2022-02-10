package transport

import (
	pb "bootcampProject/proto"
	"bootcampProject/users/domain"
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	pb.UnimplementedUserServiceServer
	createUser gt.Handler
	getUsers   gt.Handler
}

var (
	ErrBadRequest = errors.New("bad request. Doesn't match interface")
)

// NewUserGRPCServer initializes a new gRPC server
func NewUserGRPCServer(svcEndpoints UserEndpointsGRPC, logger log.Logger) pb.UserServiceServer {
	opts := []gt.ServerOption{
		gt.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		//kithttp.ServerErrorEncoder(encodeError),
	}

	return &gRPCServer{
		createUser: gt.NewServer(
			svcEndpoints.CreateUser,
			decodeCreateUserGRPCRequest,
			encodeCreateUserGRPCResponse,
			opts...,
		),
		getUsers: gt.NewServer(
			svcEndpoints.GetUsers,
			decodeGetUsersGRPCRequest,
			encodeGetUsersGRPCResponse,
			opts...,
		),
	}
}

func (s *gRPCServer) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserResp, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CreateUserResp), nil
}

func decodeCreateUserGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	if req, ok := request.(*pb.CreateUserReq); ok {
		return CreateUserRequest{User: domain.Users{
			PwdHash: req.GetPwdHash(),
			Name:    req.GetName(),
			Age:     int(req.GetAge()),
			Email:   req.GetEmail(),
		}}, nil
	}
	return CreateUserRequest{}, ErrBadRequest
}

func encodeCreateUserGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	if resp, ok := response.(CreateUserResponse); ok {
		return &pb.CreateUserResp{
			Id:    int32(resp.ID),
			Email: resp.Email,
		}, resp.Err
	}
	return &pb.CreateUserResp{Error: ErrBadRequest.Error()}, ErrBadRequest
}

func (s *gRPCServer) GetUsers(ctx context.Context, req *pb.GetUsersReq) (*pb.GetUsersResp, error) {
	_, resp, err := s.getUsers.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetUsersResp), nil
}

func decodeGetUsersGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	if req, ok := request.(*pb.GetUsersReq); ok {
		return GetUsersRequest{
			limit:  int(req.Limit),
			offset: int(req.Offset),
		}, nil
	}
	return GetUsersRequest{}, ErrBadRequest
}

func encodeGetUsersGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	if resp, ok := response.(GetUsersResponse); ok {
		return &pb.GetUsersResp{
			Users: func(users []domain.Users) []*pb.User {
				var res []*pb.User
				for _, user := range users {
					temp := &pb.User{
						Id:   int32(user.ID),
						Name: user.Name,
						Age:  int32(user.Age),
					}
					res = append(res, temp)
				}
				return res
			}(resp.Users),
		}, nil
	}
	return &pb.User{}, ErrBadRequest
}
