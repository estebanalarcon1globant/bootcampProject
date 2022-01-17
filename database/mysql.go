package database

import (
	"bootcampProject/config"
	"bootcampProject/users/domain"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type DBHandler struct {
	Conn *gorm.DB
}

var dbHandler DBHandler

func SetupDB() error {
	var err error
	dbHandler.Conn, err = getMySQLConnection()
	if err != nil {
		return err
	}
	err = dbHandler.Conn.AutoMigrate(&domain.Users{})
	if err != nil {
		return err
	}
	return nil
}

func migrateTables() error {
	return dbHandler.Conn.AutoMigrate(&domain.Users{})
}

func getMySQLConnection() (*gorm.DB, error) {
	dsn := config.GetSqlDBUsername() +
		":" + config.GetSqlDBPass() +
		"@tcp" + "(" + config.GetSqlDBHost() +
		":" + strconv.Itoa(config.GetSqlDBPort()) + ")/" +
		config.GetSqlDBDatabase() + "?" +
		"parseTime=true&loc=Local"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func GetConnection() DBHandler {
	return dbHandler
}
