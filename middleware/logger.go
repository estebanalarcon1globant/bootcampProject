package middleware

import (
	"bootcampProject/users/domain"
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

type middleware struct {
	logger log.Logger
	next   domain.UserService
}

func NewMiddleware(logger log.Logger, svc domain.UserService) *middleware {
	return &middleware{
		logger: logger,
		next:   svc,
	}
}

func (mw middleware) CreateUser(ctx context.Context, user domain.Users) (output int, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "CreateUser",
			"input", user,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.CreateUser(ctx, user)
	return
}