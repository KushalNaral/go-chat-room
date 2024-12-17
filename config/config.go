package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWT_TOKEN string
	DB_URL    string
}

func LoadConfig() *Config {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("failed to load config", err)
		return nil
	}

	return &Config{
		JWT_TOKEN: os.Getenv("JWT_TOKEN"),
		DB_URL:    os.Getenv("DB_URL"),
	}
}
