package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	HTTP_PORT   string
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %q", err)
	}
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	HTTP_PORT = os.Getenv("HTTP_PORT")
}
