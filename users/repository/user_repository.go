package repository

import (
	"bootcampProject/database"
	"bootcampProject/users/domain"
	"context"
)

type userRepository struct {
	dbHandler database.DBHandler
}

func NewUserRepository(dbHandler database.DBHandler) domain.UserRepository {
	return &userRepository{
		dbHandler: dbHandler,
	}
}

func (rep *userRepository) CreateUser(_ context.Context, user domain.Users) (int, error) {
	err := rep.dbHandler.Conn.Create(&user).Error
	return user.ID, err
}

func (rep *userRepository) GetUsers(_ context.Context, limit int, offset int) ([]domain.Users, error) {
	var users []domain.Users
	err := rep.dbHandler.Conn.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}
