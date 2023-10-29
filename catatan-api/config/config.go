package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// DatabaseURL holds the database connection string.
	DatabaseURL string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DatabaseURL = os.Getenv("DATABASE_URL")
	if DatabaseURL == "" {
		log.Fatal("DATABASE_URL not found in .env file")
	}
}
