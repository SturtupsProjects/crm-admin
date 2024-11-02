package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB_NAME string
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
}

func NewConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{}

	config.DB_NAME = os.Getenv("DB_NAME")
	config.DB_USER = os.Getenv("DB_USER")
	config.DB_PASS = os.Getenv("DB_PASS")
	config.DB_NAME = os.Getenv("DB_NAME")
	config.DB_PORT = os.Getenv("DB_PORT")

	return config
}
