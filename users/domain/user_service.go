package domain

import "context"

type UserService interface {
	CreateUser(ctx context.Context, user Users) (int, error)
	GetUsers(ctx context.Context, limit int, offset int) ([]Users, error)
}
