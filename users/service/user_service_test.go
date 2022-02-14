package service

import (
	"bootcampProject/users/domain"
	"bootcampProject/users/mocks"
	"bootcampProject/utils"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	mockUser := domain.Users{
		ID:      1,
		PwdHash: "pass",
		Name:    "test",
		Age:     24,
		Email:   "test@test.com",
	}
	mockUserRepo := new(mocks.UserRepoMock)
	mockTokenGen := new(mocks.AuthMock)
	idExpected := 1

	t.Run("success", func(t *testing.T) {
		tempMockUser := mockUser
		tempMockUser.ID = idExpected

		//User doesnt exist
		mockUserRepo.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).
			Return(domain.Users{}, nil).Once()

		mockUserRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("domain.Users")).
			Return(idExpected, nil).Once()

		u := NewUserService(mockUserRepo, mockTokenGen)

		idGot, err := u.CreateUser(context.TODO(), tempMockUser)

		assert.Equal(t, idExpected, idGot)
		assert.NoError(t, err)
		assert.Equal(t, mockUser.Name, tempMockUser.Name)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Error: create User repository", func(t *testing.T) {
		errorWant := errors.New("test error")
		tempMockUser := mockUser
		tempMockUser.ID = 1

		mockUserRepo.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).
			Return(domain.Users{}, nil).Once()

		mockUserRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("domain.Users")).
			Return(0, errorWant).Once()

		u := NewUserService(mockUserRepo, mockTokenGen)

		_, errGot := u.CreateUser(context.TODO(), tempMockUser)

		assert.EqualError(t, errGot, errorWant.Error())
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Error: user already exists", func(t *testing.T) {
		errorWant := ErrUserAlreadyExists
		tempMockUser := mockUser
		tempMockUser.ID = 1

		mockUserRepo.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).
			Return(tempMockUser, nil).Once()

		u := NewUserService(mockUserRepo, mockTokenGen)

		_, errGot := u.CreateUser(context.TODO(), tempMockUser)

		assert.EqualError(t, errGot, errorWant.Error())
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_GetUsers(t *testing.T) {
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
	mockTokenGen := new(mocks.AuthMock)
	limit, offset := 5, 1

	t.Run("success", func(t *testing.T) {
		//tempMockUser := mockUsers
		mockUserRepo.On("GetUsers", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(mockUsers, nil).Once()

		u := NewUserService(mockUserRepo, mockTokenGen)

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

		u := NewUserService(mockUserRepo, mockTokenGen)

		_, errGot := u.GetUsers(context.TODO(), limit, offset)

		assert.EqualError(t, errGot, errorWant.Error())
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_GetUserByEmail(t *testing.T) {

	mockUserRepo := new(mocks.UserRepoMock)
	mockTokenGen := new(mocks.AuthMock)
	emailTest := "test@test.com"
	userMock := domain.Users{
		ID:      10,
		PwdHash: "pass",
		Name:    "test1",
		Age:     24,
		Email:   emailTest,
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).
			Return(userMock, nil).Once()

		u := NewUserService(mockUserRepo, mockTokenGen)
		usersGot, err := u.GetUserByEmail(context.TODO(), emailTest)

		assert.NoError(t, err)
		assert.Equal(t, userMock.ID, usersGot.ID)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Error: user not found", func(t *testing.T) {
		errWant := errors.New("error not found")
		mockUserRepo.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).
			Return(domain.Users{}, errWant).Once()

		u := NewUserService(mockUserRepo, mockTokenGen)
		_, errGot := u.GetUserByEmail(context.TODO(), emailTest)

		assert.EqualError(t, errGot, errWant.Error())
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_Authenticate(t *testing.T) {

	passwordTesting := "test"
	passwordTestingHash := utils.HashSHA256("test")

	mockUserRepo := new(mocks.UserRepoMock)

	emailTest := "test@test.com"
	userMock := domain.Users{
		ID:      10,
		PwdHash: passwordTestingHash,
		Name:    "test1",
		Age:     24,
		Email:   emailTest,
	}

	auth := domain.Auth{
		Email:    "test@test.com",
		Password: passwordTesting,
	}

	mockAuthToken := new(mocks.AuthMock)
	tokenExpected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
		"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ." +
		"SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	t.Run("success", func(t *testing.T) {

		mockAuthToken.On("GenerateToken", mock.AnythingOfType("string")).
			Return(tokenExpected, nil).Once()

		mockUserRepo.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).
			Return(userMock, nil).Once()
		u := NewUserService(mockUserRepo, mockAuthToken)

		tokenGot, err := u.Authenticate(context.TODO(), auth)
		assert.NoError(t, err)
		assert.Equal(t, tokenExpected, tokenGot)
		mockAuthToken.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Error: invalid fields", func(t *testing.T) {
		authEmpty := domain.Auth{
			Email:    "",
			Password: "",
		}

		errWant := errors.New("invalid fields")
		u := NewUserService(mockUserRepo, mockAuthToken)
		_, errGot := u.Authenticate(context.TODO(), authEmpty)
		assert.EqualError(t, errWant, errGot.Error())
	})

	t.Run("Error: Error getUserByEmail", func(t *testing.T) {
		errWant := errors.New("error getUserByEmail")

		mockUserRepo.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).
			Return(domain.Users{}, errWant).Once()

		u := NewUserService(mockUserRepo, mockAuthToken)
		_, errGot := u.Authenticate(context.TODO(), auth)
		assert.EqualError(t, errWant, errGot.Error())
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Error: invalid credentials", func(t *testing.T) {
		authInvalid := domain.Auth{
			Email:    "test@test.com",
			Password: "anotherPass",
		}
		errWant := errors.New("invalid credentials")

		mockUserRepo.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).
			Return(userMock, nil).Once()

		u := NewUserService(mockUserRepo, mockAuthToken)

		_, errGot := u.Authenticate(context.TODO(), authInvalid)
		assert.EqualError(t, errWant, errGot.Error())
		mockUserRepo.AssertExpectations(t)
	})
}
