package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"
)

func TestMongo(t *testing.T) {
	err := SetupNoSqlDB()
	if err != nil {
		os.Exit(1)
	}

	noSqlDB := GetNoSQLConnection()
	user := bson.D{{"user_id", 14}, {"additional_info", "testing function"}}
	res, err := noSqlDB.Collection.InsertOne(context.Background(), user)
	fmt.Println("RESULT MONGODB: ", res.InsertedID)
}
