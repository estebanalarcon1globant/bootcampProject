package transport

import (
	"bootcampProject/config"
	pb "bootcampProject/proto"
	"bootcampProject/users/domain"
	"bootcampProject/utils"
	"context"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/golang-jwt/jwt/v4"
)

type gRPCServer struct {
	pb.UnimplementedUserServiceServer
	createUser   gt.Handler
	getUsers     gt.Handler
	authenticate gt.Handler
}

// NewUserGRPCServer initializes a new gRPC server
func NewUserGRPCServer(svcEndpoints UserEndpointsGRPC, logger log.Logger) pb.UserServiceServer {

	key := []byte(config.GetJwtSecret())
	keys := func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}

	opts := []gt.ServerOption{
		gt.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		gt.ServerBefore(kitjwt.GRPCToContext()),
		gt.ServerAfter(),
	}

	return &gRPCServer{
		createUser: gt.NewServer(
			svcEndpoints.CreateUser,
			decodeCreateUserGRPCRequest,
			encodeCreateUserGRPCResponse,
			opts...,
		),
		getUsers: gt.NewServer(
			kitjwt.NewParser(keys, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(svcEndpoints.GetUsers),
			decodeGetUsersGRPCRequest,
			encodeGetUsersGRPCResponse,
			opts...,
		),

		authenticate: gt.NewServer(
			svcEndpoints.Authenticate,
			decodeAuthenticateGRPCRequest,
			encodeAuthenticateGRPCResponse,
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
	return CreateUserRequest{}, utils.ErrBadRequest
}

func encodeCreateUserGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	if resp, ok := response.(CreateUserResponse); ok {
		return &pb.CreateUserResp{
			Id:    int32(resp.ID),
			Email: resp.Email,
		}, resp.Err
	}
	return &pb.CreateUserResp{Error: utils.ErrBadRequest.Error()}, utils.ErrBadRequest
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
	return GetUsersRequest{}, utils.ErrBadRequest
}

func encodeGetUsersGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	if resp, ok := response.(GetUsersResponse); ok {
		return &pb.GetUsersResp{
			Users: func(users []domain.Users) []*pb.User {
				var res []*pb.User
				for _, user := range users {
					temp := &pb.User{
						Id:    int32(user.ID),
						Name:  user.Name,
						Age:   int32(user.Age),
						Email: user.Email,
					}
					res = append(res, temp)
				}
				return res
			}(resp.Users),
		}, nil
	}
	return &pb.User{}, utils.ErrBadRequest
}

func (s *gRPCServer) Authenticate(ctx context.Context, req *pb.AuthReq) (*pb.AuthResp, error) {
	_, resp, err := s.authenticate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.AuthResp), nil
}

func decodeAuthenticateGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	if req, ok := request.(*pb.AuthReq); ok {
		return domain.Auth{
			Email:    req.GetEmail(),
			Password: req.GetPassword(),
		}, nil
	}
	return domain.Auth{}, utils.ErrBadRequest
}

func encodeAuthenticateGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	if resp, ok := response.(AuthResponse); ok {
		return &pb.AuthResp{Token: resp.Token}, nil
	}
	return &pb.AuthResp{}, utils.ErrBadRequest
}
