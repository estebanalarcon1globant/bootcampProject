package domain

import "context"

type UserRepository interface {
	CreateUser(ctx context.Context, user Users) (int, error)
}
