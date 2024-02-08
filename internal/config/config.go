package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
