package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port      string
	DBConn    string
	JWTSecret string
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	return &Config{
		Port:      getEnv("PORT", "8080"),
		DBConn:    getEnv("DB_CONN", "postgres://postgres:secret@localhost:5432/blog?sslmode=disable"),
		JWTSecret: getEnv("JWT_SECRET", "mysecretkey"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
