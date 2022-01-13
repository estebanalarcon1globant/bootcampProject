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

func TestUserService_CreateUser(t *testing.T) {
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

		u := NewUserService(mockUserRepo, logger)

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

		u := NewUserService(mockUserRepo, logger)

		_, errGot := u.CreateUser(context.TODO(), tempMockUser)

		assert.EqualError(t, errGot, errorWant.Error())
		mockUserRepo.AssertExpectations(t)
	})
}
