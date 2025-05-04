package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBEnv = initConfig()

type DBConfig struct {
	PublicHost string
	Port       string
	User       string
	Password   string
	Name       string
	SSLMode    string
}

func initConfig() DBConfig {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default database configuration")
	}

	return DBConfig{
		PublicHost: getEnv("DB_HOST", "localhost"),
		Port:       getEnv("DB_PORT", "5432"),
		User:       getEnv("DB_USER", "root"),
		Password:   getEnv("DB_PASSWORD", "1234"),
		Name:       getEnv("DB_NAME", "pharmacy"),
		SSLMode:    getEnv("DB_SSL_MODE", "disable"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
