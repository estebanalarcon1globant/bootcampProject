package service

import (
	"bootcampProject/users/domain"
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type userService struct {
	userRepository domain.UserRepository
	logger         log.Logger
}

func NewUserService(rep domain.UserRepository, logger log.Logger) domain.UserService {
	return &userService{
		userRepository: rep,
		logger:         logger,
	}
}

// CreateUser Create User persistent
func (s *userService) CreateUser(ctx context.Context, user domain.Users) (int, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	id, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return id, err
}
