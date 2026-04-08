package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

var once sync.Once

const defaultConfigFile = "config.env"

func LoadEnv(file string) {
	once.Do(func() {
		if file == "" {
			file = defaultConfigFile
		}
		if err := godotenv.Load(file); err != nil {
			log.Fatalf("Error loading %s file", file)
		}
	})
}

func Env(key string, defaultValue string) string {

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
