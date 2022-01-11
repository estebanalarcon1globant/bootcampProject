package repository

import (
	"bootcampProject/users/domain"
	"context"
	"github.com/go-kit/kit/log"
	"gorm.io/gorm"
)

type userRepository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewUserRepository(db *gorm.DB, logger log.Logger) domain.UserRepository {
	return &userRepository{
		db:     db,
		logger: log.With(logger, "userRepository", "sql_db"),
	}
}

func (rep *userRepository) CreateUser(ctx context.Context, user domain.Users) (int, error) {
	err := rep.db.Create(&user).Error
	return user.ID, err
}
