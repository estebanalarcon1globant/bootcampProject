package service

import (
	"bootcampProject/users/domain"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, user domain.Users) (int, error)
}
