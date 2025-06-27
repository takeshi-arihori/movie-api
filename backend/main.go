package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/takeshi-arihori/movie-api/internal/config"
	"github.com/takeshi-arihori/movie-api/internal/handlers"
	"github.com/takeshi-arihori/movie-api/internal/services"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}
	
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

	// Initialize services
	tmdbClient := services.NewTMDbClient(cfg)
	movieHandler := handlers.NewMovieHandler(tmdbClient)

	// Setup router
	router := setupRouter(movieHandler)

	// Start server
	addr := ":" + cfg.Server.Port
	fmt.Printf("Server listening on %s\n", addr)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET /api/v1/health            - Health check")
	fmt.Println("  GET /api/v1/movies/{id}       - Movie details")
	fmt.Println("  GET /api/v1/movies/{id}/credits - Movie credits")
	fmt.Println("  GET /api/v1/movies/{id}/reviews - Movie reviews")
	fmt.Println("  GET /health                   - Simple health check")
	
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// setupRouter configures and returns the HTTP router
func setupRouter(movieHandler *handlers.MovieHandler) *mux.Router {
	router := mux.NewRouter()

	// API v1 routes
	api := router.PathPrefix("/api/v1").Subrouter()

	// Movie endpoints
	api.HandleFunc("/movies/{id:[0-9]+}", movieHandler.GetMovieDetails).Methods("GET", "OPTIONS")
	api.HandleFunc("/movies/{id:[0-9]+}/credits", movieHandler.GetMovieCredits).Methods("GET", "OPTIONS")
	api.HandleFunc("/movies/{id:[0-9]+}/reviews", movieHandler.GetMovieReviews).Methods("GET", "OPTIONS")

	// Health check endpoint
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy","service":"movie-api","version":"1.0.0"}`))
	}).Methods("GET", "OPTIONS")

	// Legacy health check endpoint (for compatibility)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","service":"movie-api","version":"1.0.0"}`))
	}).Methods("GET", "OPTIONS")

	// Add CORS middleware
	router.Use(corsMiddleware)

	// Add logging middleware
	router.Use(loggingMiddleware)

	return router
}

// corsMiddleware adds CORS headers to all responses
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// loggingMiddleware logs HTTP requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
