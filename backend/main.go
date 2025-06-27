package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gorilla/mux"
	"github.com/takeshi-arihori/movie-api/internal/config"
	"github.com/takeshi-arihori/movie-api/internal/handlers"
	"github.com/takeshi-arihori/movie-api/internal/services"
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

	// Initialize services
	tmdbClient := services.NewTMDbClient(cfg)
	searchHandler := handlers.NewSearchHandler(tmdbClient)
	movieHandler := handlers.NewMovieHandler(tmdbClient)
	reviewHandler := handlers.NewReviewHandler(tmdbClient)
	personHandler := handlers.NewPersonHandler(tmdbClient)

	// Setup router
	router := setupRouter(searchHandler, movieHandler, reviewHandler, personHandler)

	// Start server
	addr := ":" + cfg.Server.Port
	fmt.Printf("Server listening on %s\n", addr)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET /api/v1/health            - Health check")
	fmt.Println("  GET /api/v1/search            - Multi search (movies, TV shows, people)")
	fmt.Println("  GET /api/v1/movies/{id}       - Movie details")
	fmt.Println("  GET /api/v1/movies/{id}/credits - Movie credits")
	fmt.Println("  GET /api/v1/movies/{id}/reviews - Movie reviews")
	fmt.Println("  GET /api/v1/tv/{id}/reviews   - TV show reviews")
	fmt.Println("  GET /api/v1/people/{id}       - Person details")
	fmt.Println("  GET /api/v1/people/{id}/movie_credits - Person movie credits")
	fmt.Println("  GET /api/v1/people/{id}/tv_credits - Person TV credits")
	fmt.Println("  GET /api/v1/people/{id}/combined_credits - Person combined credits")
	fmt.Println("  GET /health                   - Simple health check")
	
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// setupRouter configures and returns the HTTP router
func setupRouter(searchHandler *handlers.SearchHandler, movieHandler *handlers.MovieHandler, reviewHandler *handlers.ReviewHandler, personHandler *handlers.PersonHandler) *mux.Router {
	router := mux.NewRouter()

	// API v1 routes
	api := router.PathPrefix("/api/v1").Subrouter()

	// Search endpoints
	api.HandleFunc("/search", searchHandler.Search).Methods("GET", "OPTIONS")
	api.HandleFunc("/health", searchHandler.HealthCheck).Methods("GET", "OPTIONS")
	api.HandleFunc("/search/suggestions", searchHandler.GetSearchSuggestions).Methods("GET", "OPTIONS")

	// Movie endpoints
	api.HandleFunc("/movies/{id:[0-9]+}", movieHandler.GetMovieDetails).Methods("GET", "OPTIONS")
	api.HandleFunc("/movies/{id:[0-9]+}/credits", movieHandler.GetMovieCredits).Methods("GET", "OPTIONS")
	api.HandleFunc("/movies/{id:[0-9]+}/reviews", reviewHandler.GetMovieReviews).Methods("GET", "OPTIONS")

	// TV show endpoints
	api.HandleFunc("/tv/{id:[0-9]+}/reviews", reviewHandler.GetTVReviews).Methods("GET", "OPTIONS")

	// Person endpoints
	api.HandleFunc("/people/{id:[0-9]+}", personHandler.GetPersonDetails).Methods("GET", "OPTIONS")
	api.HandleFunc("/people/{id:[0-9]+}/movie_credits", personHandler.GetPersonMovieCredits).Methods("GET", "OPTIONS")
	api.HandleFunc("/people/{id:[0-9]+}/tv_credits", personHandler.GetPersonTVCredits).Methods("GET", "OPTIONS")
	api.HandleFunc("/people/{id:[0-9]+}/combined_credits", personHandler.GetPersonCombinedCredits).Methods("GET", "OPTIONS")

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
