package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Port            string
	DatabaseURL     string
	JWTSecret       string
	JWTExpiryHours  int
	Environment     string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists (ignore error in production)
	_ = godotenv.Load()

	jwtExpiryHours := 24
	if val := os.Getenv("JWT_EXPIRY_HOURS"); val != "" {
		if hours, err := strconv.Atoi(val); err == nil {
			jwtExpiryHours = hours
		}
	}

	cfg := &Config{
		Port:           getEnv("PORT", "8080"),
		DatabaseURL:    getEnv("DATABASE_URL", ""),
		JWTSecret:      getEnv("JWT_SECRET", ""),
		JWTExpiryHours: jwtExpiryHours,
		Environment:    getEnv("ENVIRONMENT", "development"),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
