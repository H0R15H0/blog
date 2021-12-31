package config

import (
	"os"
)

type Config struct {
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_NAME string
}

func GetConfig() (*Config) {
	config := &Config{
		DB_USER: os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASS"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_NAME: os.Getenv("DB_NAME"),
	}

	return config
}