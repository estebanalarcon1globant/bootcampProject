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

func TestMakeGetUsersGRPCEndpoint(t *testing.T) {
	mockUserService := new(mocks.UserServiceMock)
	mockUsers := []domain.Users{
		{ID: 1,
			PwdHash: "pass",
			Name:    "test1",
			Age:     24,
		},
		{ID: 2,
			PwdHash: "pass",
			Name:    "test2",
			Age:     30,
		},
	}

	respWant := GetUsersResponse{
		Users: mockUsers,
		Err:   nil,
	}

	t.Run("success", func(t *testing.T) {
		mockUserService.On("GetUsers", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(mockUsers, nil).Once()

		getUsers := makeGetUsersGRPCEndpoint(mockUserService)
		respGot, err := getUsers(context.TODO(), GetUsersRequest{})
		assert.NoError(t, err)
		assert.Equal(t, respWant, respGot)
		mockUserService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		errorWant := errors.New("test error")
		mockUserService.On("GetUsers", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]domain.Users{}, errorWant).Once()

		getUsers := makeGetUsersGRPCEndpoint(mockUserService)
		_, errGot := getUsers(context.TODO(), GetUsersRequest{})

		assert.EqualError(t, errGot, errorWant.Error())
		mockUserService.AssertExpectations(t)
	})
}
