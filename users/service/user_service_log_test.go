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
	}
	mockUserRepo := new(mocks.UserRepoMock)
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	idExpected := 1

	t.Run("success", func(t *testing.T) {
		tempMockUser := mockUser
		tempMockUser.ID = idExpected
		mockUserRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("domain.Users")).
			Return(idExpected, nil).Once()

		u := NewUserService(mockUserRepo)
		u = NewUserServiceLogging(logger, u)

		idGot, err := u.CreateUser(context.TODO(), tempMockUser)

		assert.Equal(t, idExpected, idGot)
		assert.NoError(t, err)
		assert.Equal(t, mockUser.Name, tempMockUser.Name)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		errorWant := errors.New("test error")
		tempMockUser := mockUser
		tempMockUser.ID = 1
		mockUserRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("domain.Users")).
			Return(0, errorWant).Once()

		u := NewUserService(mockUserRepo)
		u = NewUserServiceLogging(logger, u)

		_, errGot := u.CreateUser(context.TODO(), tempMockUser)

		assert.EqualError(t, errGot, errorWant.Error())
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserServiceLogging_GetUsers(t *testing.T) {
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
	mockUserRepo := new(mocks.UserRepoMock)
	limit, offset := 5, 1

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)

	t.Run("success", func(t *testing.T) {
		//tempMockUser := mockUsers
		mockUserRepo.On("GetUsers", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(mockUsers, nil).Once()

		u := NewUserService(mockUserRepo)
		u = NewUserServiceLogging(logger, u)

		usersGot, err := u.GetUsers(context.TODO(), limit, offset)

		assert.NoError(t, err)
		assert.Equal(t, mockUsers[0].ID, usersGot[0].ID)
		assert.Equal(t, mockUsers[1].ID, usersGot[1].ID)

		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		errorWant := errors.New("test error")
		mockUserRepo.On("GetUsers", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]domain.Users{}, errorWant).Once()

		u := NewUserService(mockUserRepo)
		u = NewUserServiceLogging(logger, u)

		_, errGot := u.GetUsers(context.TODO(), limit, offset)

		assert.EqualError(t, errGot, errorWant.Error())
		mockUserRepo.AssertExpectations(t)
	})
}
