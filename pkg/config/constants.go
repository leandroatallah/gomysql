package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBName string
	DBUser string
	DBPass string
	DBHost string
	DBPort string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file", err)
	}
	DBName = os.Getenv("DB_NAME")
	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASS")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
}
