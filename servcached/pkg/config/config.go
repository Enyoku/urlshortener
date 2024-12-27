package config

import (
	"os"
	"strconv"
)

type APIConfig struct {
	Redis Redis
	Port  int
}

type Redis struct {
	Addr     string
	Password string
	DB       int
}

func New() *APIConfig {
	return &APIConfig{
		Redis: Redis{
			Addr:     getEnvString("redis_addr", ""),
			Password: getEnvString("redis_pass", ""),
			DB:       getEnvInt("redis_db", 0),
		},
		Port: getEnvInt("port", 8000),
	}
}

func getEnvString(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return value
	}
	return defaultVal
}
