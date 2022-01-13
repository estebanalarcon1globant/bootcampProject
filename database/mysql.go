package database

import (
	"bootcampProject/users/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_USERNAME = "root"
	DB_PASSWORD = "globant12345"
	DB_NAME     = "globant_db"
	DB_HOST     = "127.0.0.1"
	DB_PORT     = "3306"
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
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func GetConnection() DBHandler {
	return dbHandler
}
