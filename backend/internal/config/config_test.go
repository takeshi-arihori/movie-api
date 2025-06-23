package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Save original environment
	originalEnv := make(map[string]string)
	envKeys := []string{
		"PORT", "ENV", "CORS_ORIGINS", "TMDB_API_KEY", "TMDB_BASE_URL",
		"POSTGRES_HOST", "POSTGRES_PORT", "POSTGRES_USER", "POSTGRES_PASSWORD", "POSTGRES_DB",
		"JWT_SECRET", "CACHE_ENABLED", "CACHE_TTL", "LOG_LEVEL",
	}
	
	for _, key := range envKeys {
		originalEnv[key] = os.Getenv(key)
	}
	
	// Clean up after test
	defer func() {
		for _, key := range envKeys {
			if originalEnv[key] != "" {
				os.Setenv(key, originalEnv[key])
			} else {
				os.Unsetenv(key)
			}
		}
	}()

	tests := []struct {
		name        string
		envVars     map[string]string
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid configuration",
			envVars: map[string]string{
				"TMDB_API_KEY": "test-api-key-12345",
				"JWT_SECRET":   "this-is-a-very-long-secret-key-for-testing-purposes-32-chars",
				"LOG_LEVEL":    "info",
			},
			expectError: false,
		},
		{
			name: "missing TMDB_API_KEY",
			envVars: map[string]string{
				"JWT_SECRET": "this-is-a-very-long-secret-key-for-testing-purposes-32-chars",
				"LOG_LEVEL":  "info",
			},
			expectError: true,
			errorMsg:    "TMDB_API_KEY is required",
		},
		{
			name: "short JWT secret",
			envVars: map[string]string{
				"TMDB_API_KEY": "test-api-key-12345",
				"JWT_SECRET":   "short",
				"LOG_LEVEL":    "info",
			},
			expectError: true,
			errorMsg:    "JWT secret must be at least 32 characters long",
		},
		{
			name: "invalid log level",
			envVars: map[string]string{
				"TMDB_API_KEY": "test-api-key-12345",
				"JWT_SECRET":   "this-is-a-very-long-secret-key-for-testing-purposes-32-chars",
				"LOG_LEVEL":    "invalid",
			},
			expectError: true,
			errorMsg:    "invalid log level",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear environment
			for _, key := range envKeys {
				os.Unsetenv(key)
			}
			
			// Set test environment variables
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			config, err := Load()

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				} else if tt.errorMsg != "" && err.Error() != tt.errorMsg {
					t.Errorf("expected error message to contain '%s', got '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if config == nil {
					t.Errorf("expected config but got nil")
				}
			}
		})
	}
}

func TestConfigDefaults(t *testing.T) {
	// Clear all environment variables
	envKeys := []string{
		"PORT", "ENV", "CORS_ORIGINS", "TMDB_API_KEY", "TMDB_BASE_URL",
		"POSTGRES_HOST", "POSTGRES_PORT", "POSTGRES_USER", "POSTGRES_PASSWORD", "POSTGRES_DB",
		"JWT_SECRET", "CACHE_ENABLED", "CACHE_TTL", "LOG_LEVEL",
	}
	
	for _, key := range envKeys {
		os.Unsetenv(key)
	}
	
	// Set required values
	os.Setenv("TMDB_API_KEY", "test-api-key-12345")
	os.Setenv("JWT_SECRET", "this-is-a-very-long-secret-key-for-testing-purposes-32-chars")
	
	defer func() {
		for _, key := range envKeys {
			os.Unsetenv(key)
		}
	}()

	config, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default values
	if config.Server.Port != "8080" {
		t.Errorf("expected default port 8080, got %s", config.Server.Port)
	}
	
	if config.Server.Environment != "development" {
		t.Errorf("expected default environment 'development', got %s", config.Server.Environment)
	}
	
	if config.TMDb.BaseURL != "https://api.themoviedb.org/3" {
		t.Errorf("expected default TMDb URL, got %s", config.TMDb.BaseURL)
	}
	
	if !config.Cache.Enabled {
		t.Errorf("expected cache to be enabled by default")
	}
	
	if config.Cache.TTL != 300 {
		t.Errorf("expected default cache TTL 300, got %d", config.Cache.TTL)
	}
	
	if config.Logging.Level != "info" {
		t.Errorf("expected default log level 'info', got %s", config.Logging.Level)
	}
}

func TestGetEnvAsBool(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		fallback bool
		expected bool
	}{
		{"true string", "true", false, true},
		{"false string", "false", true, false},
		{"1 string", "1", false, true},
		{"0 string", "0", true, false},
		{"empty string", "", true, true},
		{"invalid string", "invalid", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != "" {
				os.Setenv("TEST_BOOL", tt.value)
			} else {
				os.Unsetenv("TEST_BOOL")
			}
			
			result := getEnvAsBool("TEST_BOOL", tt.fallback)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			
			os.Unsetenv("TEST_BOOL")
		})
	}
}

func TestGetEnvAsInt(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		fallback int
		expected int
	}{
		{"valid integer", "123", 0, 123},
		{"zero", "0", 5, 0},
		{"negative", "-10", 0, -10},
		{"empty string", "", 42, 42},
		{"invalid string", "invalid", 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != "" {
				os.Setenv("TEST_INT", tt.value)
			} else {
				os.Unsetenv("TEST_INT")
			}
			
			result := getEnvAsInt("TEST_INT", tt.fallback)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
			
			os.Unsetenv("TEST_INT")
		})
	}
}