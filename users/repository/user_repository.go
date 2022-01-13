package repository

import (
	"bootcampProject/database"
	"bootcampProject/users/domain"
	"context"
	"github.com/go-kit/kit/log"
)

type userRepository struct {
	dbHandler database.DBHandler
	logger    log.Logger
}

func NewUserRepository(dbHandler database.DBHandler, logger log.Logger) domain.UserRepository {
	return &userRepository{
		dbHandler: dbHandler,
		logger:    log.With(logger, "userRepository", "sql_db"),
	}
}

func (rep *userRepository) CreateUser(ctx context.Context, user domain.Users) (int, error) {
	err := rep.dbHandler.Conn.Create(&user).Error
	return user.ID, err
}
