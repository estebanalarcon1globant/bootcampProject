package transport

import (
	pb "bootcampProject/grpc"
	"bootcampProject/users/domain"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type UserEndpointsHTTP struct {
	CreateUser endpoint.Endpoint
	GetUsers   endpoint.Endpoint
}

// MakeEndpointsHTTP initializes all Go kit endpoints for the Order service.
func MakeEndpointsHTTP(grpcClient pb.UserServiceClient) UserEndpointsHTTP {
	return UserEndpointsHTTP{
		CreateUser: makeCreateUserHTTPEndpoint(grpcClient),
		GetUsers:   makeGetUsersHTTPEndpoint(grpcClient),
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

func makeGetUsersHTTPEndpoint(grpcClient pb.UserServiceClient) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//TODO: Handle type assertion's error
		req := request.(GetUsersRequest)
		r, err := grpcClient.GetUsers(ctx, &pb.GetUsersParams{
			Limit:  int32(req.limit),
			Offset: int32(req.offset),
		})
		return GetUsersResponse{
			Users: func(users []*pb.User) []domain.Users {
				var res []domain.Users
				for _, user := range users {
					temp := domain.Users{
						ID:      int(user.Id),
						PwdHash: user.PwdHash,
						Name:    user.Name,
						Age:     int(user.Age),
					}
					res = append(res, temp)
				}
				return res
			}(r.Users),
			Err: err,
		}, err
	}
}
