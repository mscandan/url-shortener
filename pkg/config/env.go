package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	MongoDB_URI   string
	Redis_URL     string
	Database_Name string
	Port          string
}

type EnvironmentConfigMissingError struct {
	MissingKey string
}

func (m *EnvironmentConfigMissingError) Error() string {
	return m.MissingKey + " is missing from .env file"
}

func Config() (*EnvironmentConfig, error) {
	log.Println("Attempting to load env variables")

	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Println("No .env file found")
		return nil, err
	}

	mongodb_uri := os.Getenv("MONGODB_URI")

	if mongodb_uri == "" {
		return nil, &EnvironmentConfigMissingError{MissingKey: "MONGODB_URI"}
	}

	database_name := os.Getenv("DATABASE_NAME")

	if database_name == "" {
		return nil, &EnvironmentConfigMissingError{MissingKey: "DATABASE_NAME"}
	}

	redis_url := os.Getenv("REDIS_URL")

	if redis_url == "" {
		return nil, &EnvironmentConfigMissingError{MissingKey: "REDIS_URL"}
	}

	port := os.Getenv("PORT")

	if port == "" {
		return nil, &EnvironmentConfigMissingError{MissingKey: "PORT"}
	}

	log.Println("Environment variables loaded succesfully")

	return &EnvironmentConfig{
		MongoDB_URI:   mongodb_uri,
		Database_Name: database_name,
		Redis_URL:     redis_url,
		Port:          port,
	}, nil
}
