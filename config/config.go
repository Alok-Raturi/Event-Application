package config

import "os"

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

var Port = getEnv("PORT", "8080")

var Loglevel = getEnv("LOG_LEVEL", "INFO")

var DatabaseURL = getEnv("DATABASE_URL", "postgres://postgres:example@localhost:5432/eventdb?sslmode=disable")