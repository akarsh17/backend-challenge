package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load() // optional: loads from .env file if present
	if err != nil {
		log.Println("No .env file found, reading config from environment")
	}

	AppConfig = &Config{
		APIKey: os.Getenv("API_KEY"),
	}

	if AppConfig.APIKey == "" {
		log.Fatal("API_KEY must be set in environment or .env file")
	}
}
