package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/takeshi-arihori/movie-api/internal/config"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	fmt.Printf("Starting Movie API server on port %s\n", cfg.Server.Port)
	fmt.Printf("Environment: %s\n", cfg.Server.Environment)
	fmt.Printf("Log Level: %s\n", cfg.Logging.Level)
	fmt.Printf("Cache Enabled: %v\n", cfg.Cache.Enabled)

	// Simple health check endpoint for now
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","service":"movie-api"}`))
	})

	// Start server
	addr := ":" + cfg.Server.Port
	fmt.Printf("Server listening on %s\n", addr)
	
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}