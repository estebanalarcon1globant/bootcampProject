package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

type config struct {
	sqlDBHost     string
	sqlDBPort     int
	sqlDBDatabase string
	sqlDBUsername string
	sqlDBPass     string
}

var conf config

func LoadConfiguration() error {

	c := &config{}

	err := godotenv.Load()
	if err != nil {
		log.
			Error().
			Msg("Environment variables not found.")
	}

	errs := make([]string, 0)

	c.sqlDBHost = os.Getenv("DB_HOST")
	if c.sqlDBHost == "" {
		errs = append(errs, "Error variable database.mysql.host from .env")
	}

	c.sqlDBPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		errs = append(errs, "Error variable database.mysql.port from .env")
	}

	c.sqlDBDatabase = os.Getenv("DB_NAME")
	if c.sqlDBDatabase == "" {
		errs = append(errs, "Error variable database.mysql.database from .env")
	}

	c.sqlDBUsername = os.Getenv("DB_USER")
	if c.sqlDBUsername == "" {
		errs = append(errs, "Error variable database.mysql.username from .env")
	}

	c.sqlDBPass = os.Getenv("DB_PASSWORD")
	if c.sqlDBPass == "" {
		errs = append(errs, "Error variable database.mysql.pass from .env")
	}

	if len(errs) > 0 {
		log.Error().
			Interface("errors", errs).
			Msg("Required enviroments not found")
		return errors.New("errors with arguments")
	}
	conf = *c
	return nil
}

func GetSqlDBHost() string {
	return conf.sqlDBHost
}

func GetSqlDBPort() int {
	return conf.sqlDBPort
}
func GetSqlDBDatabase() string {
	return conf.sqlDBDatabase
}
func GetSqlDBUsername() string {
	return conf.sqlDBUsername
}
func GetSqlDBPass() string {
	return conf.sqlDBPass
}
