package service

import (
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

func TestUserServiceLogging_CreateUser(t *testing.T) {
	mockUser := domain.Users{
		ID:      1,
		PwdHash: "pass",
		Name:    "test",
		Age:     24,
		Email:   "test@test.com",
	}

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	idExpected := 1

	mockUserSvc := new(mocks.UserServiceMock)

	t.Run("success", func(t *testing.T) {
		tempMockUser := mockUser
		tempMockUser.ID = idExpected
		mockUserSvc.On("CreateUser", mock.Anything, mock.AnythingOfType("domain.Users")).
			Return(idExpected, nil).Once()

		u := NewUserServiceLogging(logger, mockUserSvc)

		idGot, err := u.CreateUser(context.TODO(), tempMockUser)

		assert.Equal(t, idExpected, idGot)
		assert.NoError(t, err)
		assert.Equal(t, mockUser.Name, tempMockUser.Name)
		mockUserSvc.AssertExpectations(t)
	})

	t.Run("error: user service", func(t *testing.T) {
		errorWant := errors.New("test error")
		tempMockUser := mockUser
		tempMockUser.ID = 1
		mockUserSvc.On("CreateUser", mock.Anything, mock.AnythingOfType("domain.Users")).
			Return(0, errorWant).Once()

		u := NewUserServiceLogging(logger, mockUserSvc)

		_, errGot := u.CreateUser(context.TODO(), tempMockUser)

		assert.EqualError(t, errGot, errorWant.Error())
		mockUserSvc.AssertExpectations(t)
	})
}

func TestUserServiceLogging_GetUsers(t *testing.T) {
	mockUsers := []domain.Users{
		{ID: 1,
			PwdHash: "pass",
			Name:    "test1",
			Age:     24,
			Email:   "test@test.com",
		},
		{ID: 2,
			PwdHash: "pass",
			Name:    "test2",
			Age:     30,
			Email:   "test2@test.com",
		},
	}
	mockUserSvc := new(mocks.UserServiceMock)
	limit, offset := 5, 1

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)

	t.Run("success", func(t *testing.T) {
		//tempMockUser := mockUsers
		mockUserSvc.On("GetUsers", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(mockUsers, nil).Once()

		u := NewUserServiceLogging(logger, mockUserSvc)

		usersGot, err := u.GetUsers(context.TODO(), limit, offset)

		assert.NoError(t, err)
		assert.Equal(t, mockUsers[0].ID, usersGot[0].ID)
		assert.Equal(t, mockUsers[1].ID, usersGot[1].ID)
		mockUserSvc.AssertExpectations(t)
	})
	t.Run("error: user service", func(t *testing.T) {
		errorWant := errors.New("test error")
		mockUserSvc.On("GetUsers", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]domain.Users{}, errorWant).Once()

		u := NewUserServiceLogging(logger, mockUserSvc)

		_, errGot := u.GetUsers(context.TODO(), limit, offset)

		assert.EqualError(t, errGot, errorWant.Error())
		mockUserSvc.AssertExpectations(t)
	})
}

func TestUserServiceLogging_Authenticate(t *testing.T) {
	mockUserSvc := new(mocks.UserServiceMock)

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)

	authMock := domain.Auth{
		Email:    "test@test.com",
		Password: "password",
	}

	t.Run("success", func(t *testing.T) {

		tokenExpected := "token"
		mockUserSvc.On("Authenticate", mock.Anything, mock.AnythingOfType("domain.Auth")).
			Return(tokenExpected, nil).Once()

		svcLog := NewUserServiceLogging(logger, mockUserSvc)

		tokenGot, err := svcLog.Authenticate(context.TODO(), authMock)

		assert.NoError(t, err)
		assert.Equal(t, tokenExpected, tokenGot)
		mockUserSvc.AssertExpectations(t)
	})

	t.Run("Error: authenticate service", func(t *testing.T) {
		errWant := errors.New("error in authenticate service")
		mockUserSvc.On("Authenticate", mock.Anything, mock.AnythingOfType("domain.Auth")).
			Return("", errWant).Once()

		svcLog := NewUserServiceLogging(logger, mockUserSvc)
		_, errGot := svcLog.Authenticate(context.TODO(), authMock)
		assert.EqualError(t, errGot, errWant.Error())
		mockUserSvc.AssertExpectations(t)
	})
}
