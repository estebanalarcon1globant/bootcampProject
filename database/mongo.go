package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NoSQLHandler struct {
	Collection *mongo.Collection
}

var noSqlHandler NoSQLHandler

func SetupNoSqlDB() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	noSqlHandler.Collection = client.Database("bootcamp_mongo").Collection("users")
	return nil
}

func GetNoSQLConnection() NoSQLHandler {
	return noSqlHandler
}
