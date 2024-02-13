package config

import (
	"os"
	"strconv"
)

var Config *ConfigModel

type ConfigModel struct {
	RateLimit        int
	RateLimitTime    int
	JWTKey           string
	DatabaseHost     string
	DatabaseUsername string
	DatabasePassword string
	DatabasePort     int
	DatabaseName     string
}

func LoadConfig() {
	// Load values from environment variables
	Config = &ConfigModel{
		RateLimit:        getEnvAsInt("RATE_LIMIT", 2),
		RateLimitTime:    getEnvAsInt("RATE_LIMIT_TIME", 1),
		JWTKey:           getEnv("JWT_KEY", "mykey"),
		DatabaseHost:     getEnv("POSTGRES_HOST", ""),
		DatabaseUsername: getEnv("POSTGRES_USERNAME", ""),
		DatabasePassword: getEnv("POSTGRES_PASSWORD", ""),
		DatabasePort:     getEnvAsInt("POSTGRES_PORT", 5432), // Default port is 5432
		DatabaseName:     getEnv("POSTGRES_DBNAME", ""),
	}
}

// retrieves the environment value variable as string
func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// retrieves the environment value variable as integer
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return valueInt
}
