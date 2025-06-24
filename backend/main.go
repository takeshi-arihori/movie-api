package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

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
	fmt.Println("pprof debugging available at http://localhost:6060/debug/pprof/")

	// Start pprof server for debugging in development
	if cfg.Server.Environment == "development" {
		go func() {
			fmt.Println("Starting pprof server on :6060")
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}

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
