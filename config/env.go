package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	MongoDB_URI   string
	Database_Name string
}

type EnvironmentConfigMissingError struct {
	MissingKey string
}

func (m *EnvironmentConfigMissingError) Error() string {
	return m.MissingKey + " is missing from .env file"
}

func Config() (*EnvironmentConfig, error) {
	if err := godotenv.Load(); err != nil {
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

	return &EnvironmentConfig{
		MongoDB_URI:   mongodb_uri,
		Database_Name: database_name,
	}, nil
}
