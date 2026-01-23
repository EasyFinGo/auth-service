package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresURL string
}

func Load() *Config {
	_ = godotenv.Load()
	return &Config{
		PostgresURL: getEnv("DB_URL"),
	}
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return ""
}
