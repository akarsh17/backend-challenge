package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey string
}

var AppConfig *Config

func LoadConfig() error {
	err := godotenv.Load() // optional: loads from .env file if present
	if err != nil {
		log.Println("No .env file found, reading config from environment")
		return err
	}

	AppConfig = &Config{
		APIKey: os.Getenv("API_KEY"),
	}

	if AppConfig.APIKey == "" {
		log.Fatal("API_KEY must be set in environment or .env file")
		return fmt.Errorf("API_KEY must be set in environment or .env file")
	}
	return nil
}
