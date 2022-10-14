package repository

import (
	"bootcampProject/additional_information/internal/domain"
	"bootcampProject/database"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userAdditionalInfoRepo struct {
	noSqlDBHandler database.NoSQLHandler
}

func NewUserAdditionalInfoRepo(noSqlDB database.NoSQLHandler) domain.AdditionalInfoRepository {
	return &userAdditionalInfoRepo{
		noSqlDBHandler: noSqlDB,
	}
}

func (r *userAdditionalInfoRepo) CreateAdditionalInfo(ctx context.Context, info domain.UserAdditionalInfo) (string, error) {
	user := bson.D{{"user_id", info.UserId}, {"additional_info", info.AdditionalInfo}}
	res, err := r.noSqlDBHandler.Collection.InsertOne(ctx, user)
	fmt.Println("RESULT MONGODB: ", res.InsertedID)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}
