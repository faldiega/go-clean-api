package config

import (
	"errors"
	"go-clean-api/internal/utils"
	"os"
	"strings"
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
			// File tidak ada = normal di container; pakai environment variable dari OS
			if errors.Is(err, os.ErrNotExist) {
				log.Warnf("%s not found, using OS environment variables", file)
				return
			}
			// File ada tapi isinya rusak = error sungguhan
			log.Fatalf("Error loading %s file: %v", file, err)
		}
	})
}

func Env(key string, defaultValue string) string {

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func EnvAsInt(key string, defaultValue int) int {

	strVal := Env(key, "")
	if strVal != "" {
		return utils.StringToInt(strVal)
	}

	return defaultValue
}

func EnvAsBool(key string, defaultValue bool) bool {

	strVal := Env(key, "")
	if strVal != "" {
		return utils.StringToBool(strVal)
	}

	return defaultValue
}

// pisahkan string "a,b,c" menjadi []string{"a", "b", "c"}
// dan trim spasi di tiap URL
func ParseWhitelistURLs(raw string) []string {
	if raw == "" {
		return []string{}
	}

	parts := strings.Split(raw, ",")
	result := make([]string, 0, len(parts))

	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}
