package config

import (
	"os"
)

//AppConfig Application configuration
type AppConfig struct {
	Port int
	Url  string
}

//GetConfig Initiatilize config in singleton way
func GetConfig() *AppConfig {
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	return &AppConfig{
		Port: 8080,
		Url:  connectionString,
	}
}
