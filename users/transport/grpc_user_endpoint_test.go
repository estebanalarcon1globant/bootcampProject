package transport

import (
	"bootcampProject/users/domain"
	"bootcampProject/users/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestMakeCreateUserGRPCEndpoint(t *testing.T) {
	mockUserReq := CreateUserRequest{User: domain.Users{
		ID:      1,
		PwdHash: "pass",
		Name:    "test",
		Age:     24,
	}}
	mockUserService := new(mocks.UserServiceMock)
	idExpected := 1
	resExpected := CreateUserResponse{
		ID:  idExpected,
		Err: nil,
	}

	t.Run("success", func(t *testing.T) {
		tempMockUserReq := mockUserReq
		mockUserService.On("CreateUser", mock.Anything, mock.AnythingOfType("domain.Users")).
			Return(idExpected, nil).Once()
		grpc := MakeEndpointsGRPC(mockUserService)
		resGot, err := grpc.CreateUser(context.TODO(), tempMockUserReq)
		assert.NoError(t, err)
		assert.Equal(t, resExpected, resGot)
		mockUserService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		errorWant := errors.New("test error")
		tempMockUserReq := mockUserReq
		mockUserService.On("CreateUser", mock.Anything, mock.AnythingOfType("domain.Users")).
			Return(0, errorWant).Once()
		grpc := MakeEndpointsGRPC(mockUserService)
		_, errGot := grpc.CreateUser(context.TODO(), tempMockUserReq)
		assert.EqualError(t, errGot, errorWant.Error())
		mockUserService.AssertExpectations(t)
	})
}
