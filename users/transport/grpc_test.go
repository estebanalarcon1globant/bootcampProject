package transport

import (
	pb "bootcampProject/grpc"
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

func TestDecodeCreateUserGRPCRequest(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		decodeIn := &pb.NewUser{
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
			PwdHash: "test",
			Name:    "nameTest",
			Age:     20,
		}
		_, errGot := decodeCreateUserGRPCRequest(context.TODO(), decodeError)
		assert.EqualError(t, errGot, ErrBadRequest.Error())
	})
}

func TestEncodeCreateUserGRPCRequest(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedId := 1
		encodeIn := CreateUserResponse{
			ID:  expectedId,
			Err: nil,
		}
		resWant := &pb.User{
			Id: int32(expectedId),
		}
		resGot, errGot := encodeCreateUserGRPCResponse(context.TODO(), encodeIn)
		assert.NoError(t, errGot)
		assert.Equal(t, resWant, resGot)
	})

	t.Run("error on request", func(t *testing.T) {
		encodeError := &pb.User{
			PwdHash: "test",
			Name:    "nameTest",
			Age:     20,
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

	userMock := &pb.NewUser{
		PwdHash: "test",
		Name:    "nameTest",
		Age:     20,
	}

	t.Run("success", func(t *testing.T) {
		tempUserMock := userMock
		idExpected := 1
		resWant := &pb.User{
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
