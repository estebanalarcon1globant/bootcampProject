package transport

import (
	pb "bootcampProject/grpc"
	"bootcampProject/users/domain"
	"bootcampProject/users/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestMakeCreateUserHTTPEndpoint(t *testing.T) {
	mockUserReq := CreateUserRequest{User: domain.Users{
		ID:      1,
		PwdHash: "pass",
		Name:    "test",
		Age:     24,
	}}
	mockUserGrpc := new(mocks.UserGrpcMock)
	idExpected := 1
	resMock := &pb.User{
		Id:      int32(idExpected),
		PwdHash: "test",
		Name:    "nameTest",
		Age:     10,
	}

	resExpected := CreateUserResponse{
		ID:  idExpected,
		Err: nil,
	}
	//CAUTION: mock.AnythingOfType take package's name, and don't take alias of packages
	t.Run("success", func(t *testing.T) {
		tempMockUserReq := mockUserReq
		mockUserGrpc.On("CreateUser", mock.Anything, mock.AnythingOfType("*grpc.NewUser")).
			Return(resMock, nil).Once()
		http := MakeEndpointsHTTP(mockUserGrpc)
		resGot, err := http.CreateUser(context.TODO(), tempMockUserReq)
		assert.NoError(t, err)
		assert.Equal(t, resExpected, resGot)
		mockUserGrpc.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		errorWant := errors.New("test error")
		tempMockUserReq := mockUserReq
		mockUserGrpc.On("CreateUser", mock.Anything, mock.AnythingOfType("*grpc.NewUser")).
			Return(&pb.User{}, errorWant).Once()
		http := MakeEndpointsHTTP(mockUserGrpc)
		_, errGot := http.CreateUser(context.TODO(), tempMockUserReq)
		assert.EqualError(t, errGot, errorWant.Error())
		mockUserGrpc.AssertExpectations(t)
	})
}
