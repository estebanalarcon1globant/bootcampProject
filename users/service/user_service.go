package service

import (
	"bootcampProject/users/domain"
	"context"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUserService(rep domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: rep,
	}
	//return userServiceLogging{logger, userSvc}
}

// CreateUser Create User persistent
func (s *userService) CreateUser(ctx context.Context, user domain.Users) (int, error) {
	//logger := log.With(s.logger, "method", "CreateUser")
	return s.userRepository.CreateUser(ctx, user)
	//if err != nil {
	//	level.Error(s.logger).Log("err", err)
	//}
	//return id, err
}

func (s *userService) GetUsers(ctx context.Context, limit int, offset int) ([]domain.Users, error) {
	return s.userRepository.GetUsers(ctx, limit, offset)
}
