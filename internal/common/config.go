package common

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var DBEnv = initDBConfig()
var JWTEnv = initJWTConfig()

type DBConfig struct {
	PublicHost string
	Port       string
	User       string
	Password   string
	Name       string
	SSLMode    string
}

type JWTConfig struct {
	JWTSecret     string
	JWTExpiration int64
}

func initDBConfig() DBConfig {
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

func initJWTConfig() JWTConfig {
	return JWTConfig{
		JWTSecret:     getEnv("JWT_SECRET", "secret_jwt_key"),
		JWTExpiration: getEnvAsInt("JWT_EXP", 3600*24),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 2)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
