package config

import (
	"fmt"

	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBSSLMode     string
	LogFilename   string
	LogMaxSize    string
	LogMaxBackups string
	LogMaxAge     string
	LogCompress   string
}

func LoadConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Error(fmt.Sprintf("godotenv.Load() | , %s", err.Error()))
		panic(err)
	}

	return &Config{
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnv("DB_PORT", "5432"),
		DBUser:        getEnv("DB_USER", "postgres"),
		DBPassword:    getEnv("DB_PASSWORD", "postgres"),
		DBName:        getEnv("DB_NAME", "gosimpleapi"),
		DBSSLMode:     getEnv("DB_SSLMODE", "disable"),
		LogFilename:   getEnv("LOG_FILENAME", "app"),
		LogMaxSize:    getEnv("LOG_MAX_SIZE", "1"),
		LogMaxBackups: getEnv("LOG_MAX_BACKUPS", "3"),
		LogMaxAge:     getEnv("LOG_MAX_AGE", "1"),
		LogCompress:   getEnv("LOG_COMPRESS", "true"),
	}
}

func getEnv(key string, defaultValue string) string {

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
