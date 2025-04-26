package config

import (
	"os"
)

type Config struct {
	Port         string
	DatabaseURL  string
	RedisURL     string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port:         getEnv("PORT", "8000"),
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://tasker_user:strongpassword@localhost:5432/tasker_db?sslmode=disable"),
		RedisURL:     getEnv("REDIS_URL", "redis://localhost:6379"),
	}
	return cfg, nil
}

// Helper: récupère une variable d'environnement, sinon valeur par défaut
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
