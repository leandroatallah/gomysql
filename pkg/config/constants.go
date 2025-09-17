package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	name, user, pass, host, port string
}

func NewDBConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file", err)
	}

	return &DBConfig{
		name: os.Getenv("DB_NAME"),
		user: os.Getenv("DB_USER"),
		pass: os.Getenv("DB_PASS"),
		host: os.Getenv("DB_HOST"),
		port: os.Getenv("DB_PORT"),
	}
}
