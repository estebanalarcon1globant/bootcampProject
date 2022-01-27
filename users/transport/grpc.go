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
		//TODO: endpoint getUsers
		getUsers: gt.NewServer(
			svcEndpoints.GetUsers,
			decodeGetUsersGRPCRequest,
			encodeGetUsersGRPCResponse,
			opts...,
		),
	}
}

//func (s *gRPCServer) mustEmbedUnimplementedUserServiceServer(){}

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

func (s *gRPCServer) GetUsers(ctx context.Context, req *pb.GetUsersParams) (*pb.UserList, error) {
	_, resp, err := s.getUsers.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UserList), nil
}

func decodeGetUsersGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	if req, ok := request.(*pb.GetUsersParams); ok {
		return GetUsersRequest{
			limit:  int(req.Limit),
			offset: int(req.Offset),
		}, nil
	}
	return GetUsersRequest{}, ErrBadRequest
}

func encodeGetUsersGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	if resp, ok := response.(GetUsersResponse); ok {
		return &pb.UserList{
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
