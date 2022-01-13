package domain

import "context"

type UserService interface {
	CreateUser(ctx context.Context, user Users) (int, error)
}
