package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Server   ServerConfig
	TMDb     TMDbConfig
	Database DatabaseConfig
	Security SecurityConfig
}

type ServerConfig struct {
	Port        string
	Environment string
	CORSOrigins []string
}

type TMDbConfig struct {
	APIKey  string
	BaseURL string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type SecurityConfig struct {
	JWTSecret string
}

func Load() (*Config, error) {
	config := &Config{
		Server: ServerConfig{
			Port:        getEnv("PORT", "8080"),
			Environment: getEnv("ENV", "development"),
			CORSOrigins: strings.Split(getEnv("CORS_ORIGINS", "http://localhost:3000,http://localhost:3005"), ","),
		},
		TMDb: TMDbConfig{
			APIKey:  getEnv("TMDB_API_KEY", ""),
			BaseURL: getEnv("TMDB_BASE_URL", "https://api.themoviedb.org/3"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnvAsInt("POSTGRES_PORT", 5432),
			User:     getEnv("POSTGRES_USER", "developer"),
			Password: getEnv("POSTGRES_PASSWORD", "password"),
			DBName:   getEnv("POSTGRES_DB", "movieapi"),
		},
		Security: SecurityConfig{
			JWTSecret: getEnv("JWT_SECRET", "your-secret-key"),
		},
	}

	return config, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}