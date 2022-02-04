package service

import (
	"bootcampProject/users/domain"
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"time"
)

type userServiceLogging struct {
	logger log.Logger
	next   domain.UserService
}

// NewUserServiceLogging returns a new instance of a logging Service.
func NewUserServiceLogging(logger log.Logger, svc domain.UserService) domain.UserService {
	return &userServiceLogging{logger, svc}
}

func (mw *userServiceLogging) CreateUser(ctx context.Context, user domain.Users) (output int, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "CreateUser",
			"input", getStringFromStruct(user),
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.CreateUser(ctx, user)
	return
}

func (mw *userServiceLogging) GetUsers(ctx context.Context, limit int, offset int) (output []domain.Users, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "CreateUser",
			"limit", limit,
			"offset", offset,
			"output", getStringFromStruct(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.GetUsers(ctx, limit, offset)
	return
}

func (mw *userServiceLogging) GetUserByEmail(ctx context.Context, email string) (output domain.Users, err error) {
	return domain.Users{}, nil
}

func (mw *userServiceLogging) Authenticate(ctx context.Context, auth domain.Auth) (output string, err error) {
	return "", nil
}

func getStringFromStruct(entity interface{}) string {
	return fmt.Sprintf("%#v", entity)
}
