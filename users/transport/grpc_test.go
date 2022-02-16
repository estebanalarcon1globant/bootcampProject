package transport

import (
	pb "bootcampProject/proto"
	"bootcampProject/users/domain"
	"bootcampProject/users/mocks"
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

var (
	ErrTesting = errors.New("error testing")
)

func TestDecodeAuthenticateGRPCRequest(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		emailTest := "test@test.com"
		passwordTest := "test"

		decodeIn := &pb.AuthReq{
			Email:    emailTest,
			Password: passwordTest,
		}

		resWant := &domain.Auth{
			Email:    emailTest,
			Password: passwordTest,
		}

		resGot, errGot := decodeAuthenticateGRPCRequest(context.TODO(), decodeIn)
		assert.NoError(t, errGot)
		assert.Equal(t, resWant, resGot)
	})

	t.Run("error in request: doesn't match interface", func(t *testing.T) {
		decodeError := &pb.User{
			Name: "nameTest",
			Age:  20,
		}
		_, errGot := decodeAuthenticateGRPCRequest(context.TODO(), decodeError)
		assert.EqualError(t, errGot, ErrBadRequest.Error())
	})
}

func TestEncodeAuthenticateGRPCResponse(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		token := "token"
		encodeIn := AuthResponse{Token: token}
		resWant := &pb.AuthResp{Token: token}

		resGot, errGot := encodeAuthenticateGRPCResponse(context.TODO(), encodeIn)
		assert.NoError(t, errGot)
		assert.Equal(t, resWant, resGot)
	})

	t.Run("error in request", func(t *testing.T) {
		encodeError := &pb.CreateUserResp{
			Id:    1,
			Email: "test@test.com",
			Error: "",
		}
		_, errGot := encodeAuthenticateGRPCResponse(context.TODO(), encodeError)
		assert.EqualError(t, errGot, ErrBadRequest.Error())
	})
}

func TestDecodeCreateUserGRPCRequest(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		decodeIn := &pb.CreateUserReq{
			PwdHash: "test",
			Name:    "nameTest",
			Age:     20,
		}
		resWant := CreateUserRequest{User: domain.Users{
			PwdHash: "test",
			Name:    "nameTest",
			Age:     20,
		}}
		resGot, errGot := decodeCreateUserGRPCRequest(context.TODO(), decodeIn)
		assert.NoError(t, errGot)
		assert.Equal(t, resWant, resGot)
	})

	t.Run("error on request", func(t *testing.T) {
		decodeError := &pb.User{
			Name: "nameTest",
			Age:  20,
		}
		_, errGot := decodeCreateUserGRPCRequest(context.TODO(), decodeError)
		assert.EqualError(t, errGot, ErrBadRequest.Error())
	})
}

func TestEncodeCreateUserGRPCResponse(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedId := 1
		encodeIn := CreateUserResponse{
			ID:  expectedId,
			Err: nil,
		}
		resWant := &pb.CreateUserResp{
			Id: int32(expectedId),
		}
		resGot, errGot := encodeCreateUserGRPCResponse(context.TODO(), encodeIn)
		assert.NoError(t, errGot)
		assert.Equal(t, resWant, resGot)
	})

	t.Run("error on request", func(t *testing.T) {
		encodeError := &pb.CreateUserResp{
			Id:    1,
			Email: "test@test.com",
			Error: "",
		}
		_, errGot := encodeCreateUserGRPCResponse(context.TODO(), encodeError)
		assert.EqualError(t, errGot, ErrBadRequest.Error())
	})
}

func TestGRPCServer_CreateUser(t *testing.T) {
	userSvcMock := new(mocks.UserServiceMock)
	endpointsGRPC := MakeEndpointsGRPC(userSvcMock)

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	grpcServer := NewUserGRPCServer(endpointsGRPC, logger)

	userMock := &pb.CreateUserReq{
		PwdHash: "test",
		Name:    "nameTest",
		Age:     20,
		Email:   "test@test.com",
	}

	t.Run("success", func(t *testing.T) {
		tempUserMock := userMock
		idExpected := 1
		resWant := &pb.CreateUserResp{
			Id: int32(idExpected),
		}
		userSvcMock.On("CreateUser", mock.Anything, mock.AnythingOfType("domain.Users")).
			Return(idExpected, nil).Once()

		resGot, err := grpcServer.CreateUser(context.TODO(), tempUserMock)
		assert.NoError(t, err)
		assert.Equal(t, resWant, resGot)
		userSvcMock.AssertExpectations(t)
		//mock := new(mocks.UserGrpcMock)
	})

	t.Run("error on serve", func(t *testing.T) {
		tempUserMock := userMock
		userSvcMock.On("CreateUser", mock.Anything, mock.AnythingOfType("domain.Users")).
			Return(0, ErrTesting).Once()

		_, errGot := grpcServer.CreateUser(context.TODO(), tempUserMock)
		assert.EqualError(t, errGot, ErrTesting.Error())
		userSvcMock.AssertExpectations(t)
	})
}

func TestDecodeGetUsersRequest(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		limit, offset := 10, 0
		decodeIn := &pb.GetUsersReq{
			Limit:  int32(limit),
			Offset: int32(offset),
		}
		resWant := GetUsersRequest{
			limit:  limit,
			offset: offset,
		}

		resGot, errGot := decodeGetUsersGRPCRequest(context.TODO(), decodeIn)
		assert.NoError(t, errGot)
		assert.Equal(t, resWant, resGot)
	})

	t.Run("error on request", func(t *testing.T) {
		decodeError := &pb.User{}
		_, errGot := decodeGetUsersGRPCRequest(context.TODO(), decodeError)
		assert.EqualError(t, errGot, ErrBadRequest.Error())
	})
}

func TestEncodeGetUsersResponse(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		encodeIn := GetUsersResponse{
			Users: []domain.Users{
				{
					ID:      1,
					PwdHash: "test1",
					Name:    "nameTest1",
					Age:     10,
				},
				{
					ID:      2,
					PwdHash: "test2",
					Name:    "nameTest2",
					Age:     10,
				},
			},
			Err: nil,
		}
		resWant := &pb.GetUsersResp{
			Users: []*pb.User{
				{
					Id:   int32(encodeIn.Users[0].ID),
					Name: encodeIn.Users[0].Name,
					Age:  int32(encodeIn.Users[0].Age),
				},
				{
					Id:   int32(encodeIn.Users[1].ID),
					Name: encodeIn.Users[1].Name,
					Age:  int32(encodeIn.Users[1].Age),
				},
			},
		}

		resGot, errGot := encodeGetUsersGRPCResponse(context.TODO(), encodeIn)
		//fmt.Println(resGot)
		assert.NoError(t, errGot)
		assert.ObjectsAreEqualValues(resWant, resGot)
	})

	t.Run("error on request", func(t *testing.T) {
		encodeError := &pb.User{
			Name: "nameTest",
			Age:  20,
		}
		_, errGot := encodeGetUsersGRPCResponse(context.TODO(), encodeError)
		assert.EqualError(t, errGot, ErrBadRequest.Error())
	})
}
