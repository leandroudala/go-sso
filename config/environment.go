package config

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type ENVIRONMENT struct {
	HOST string
	PORT string
}

var Environment = ENVIRONMENT{
	HOST: "localhost",
	PORT: "8080",
}

func ReadEnv() error {
	log.Println("Reading environment variables")
	if err := godotenv.Load(); err != nil {
		return errors.New("File .env not found. Aborting...")
	}

	var port = os.Getenv("PORT")
	if port != "" {
		Environment.PORT = port
	}

	var host = os.Getenv("HOST")
	if host != "" {
		Environment.HOST = host
	}

	return nil
}

func GetAddr() string {
	return Environment.HOST + ":" + Environment.PORT
}

func GetDatabaseDSN() string {
	// loading environment variables for database
	host := os.Getenv("DATABASE_HOST")
	username := os.Getenv("DATABASE_USERNAME")
	password := url.QueryEscape(os.Getenv("DATABASE_PASSWORD"))
	database := os.Getenv("DATABASE_SCHEMA")

	return fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s",
		username,
		password,
		host,
		database,
	)
}

func IsAutoMigrateEnabled() bool {
	DATABASE_AUTOMIGRATE := strings.ToLower(os.Getenv("DATABASE_AUTOMIGRATE"))
	return DATABASE_AUTOMIGRATE == "true" || DATABASE_AUTOMIGRATE == "1"
}

func GetMySQLDatabaseDSN() string {
	host := os.Getenv("DATABASE_HOST")
	username := os.Getenv("DATABASE_USERNAME")
	password := url.QueryEscape(os.Getenv("DATABASE_PASSWORD"))
	database := os.Getenv("DATABASE_SCHEMA")

	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		database,
	)
}
