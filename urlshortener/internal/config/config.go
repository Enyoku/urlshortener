package config

import (
	"os"
)

type Config struct {
	DB   MySQLConfig
	Port string
}

type MySQLConfig struct {
	Username string
	Password string
	Name     string
}

func New() *Config {
	return &Config{
		DB: MySQLConfig{
			Username: getEnv("MYSQL_USER", ""),
			Password: getEnv("MYSQL_PWD", ""),
			Name:     getEnv("MYSQL_NAME", ""),
		},
		Port: getEnv("port", ":8080"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
