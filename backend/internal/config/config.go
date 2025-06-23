package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Server   ServerConfig
	TMDb     TMDbConfig
	Database DatabaseConfig
	Security SecurityConfig
	Cache    CacheConfig
	Logging  LoggingConfig
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

type CacheConfig struct {
	Enabled bool
	TTL     int // Time to live in seconds
}

type LoggingConfig struct {
	Level string
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
		Cache: CacheConfig{
			Enabled: getEnvAsBool("CACHE_ENABLED", true),
			TTL:     getEnvAsInt("CACHE_TTL", 300), // 5 minutes default
		},
		Logging: LoggingConfig{
			Level: getEnv("LOG_LEVEL", "info"),
		},
	}

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
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

func getEnvAsBool(key string, fallback bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return fallback
}

// Validate validates all configuration values
func (c *Config) Validate() error {
	// TMDb API key is required
	if c.TMDb.APIKey == "" {
		return fmt.Errorf("TMDB_API_KEY is required")
	}

	// Server port validation
	if c.Server.Port == "" {
		return fmt.Errorf("server port cannot be empty")
	}

	// Log level validation
	validLogLevels := []string{"debug", "info", "warn", "error"}
	validLevel := false
	for _, level := range validLogLevels {
		if strings.ToLower(c.Logging.Level) == level {
			validLevel = true
			break
		}
	}
	if !validLevel {
		return fmt.Errorf("invalid log level: %s (valid: debug, info, warn, error)", c.Logging.Level)
	}

	// JWT Secret validation
	if len(c.Security.JWTSecret) < 32 {
		return fmt.Errorf("JWT secret must be at least 32 characters long")
	}

	return nil
}