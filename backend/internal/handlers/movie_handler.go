// Package handlers provides HTTP handlers for movie-related endpoints.
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/takeshi-arihori/movie-api/internal/models"
	"github.com/takeshi-arihori/movie-api/internal/services"
)

// MovieClient defines the interface for movie-related TMDb operations
type MovieClient interface {
	GetMovieDetails(ctx context.Context, movieID int) (*models.MovieDetails, error)
	GetMovieCredits(ctx context.Context, movieID int) (*models.MovieCredits, error)
	GetMovieReviews(ctx context.Context, movieID int, page int) (*models.MovieReviews, error)
}

// MovieHandler handles movie-related HTTP requests
type MovieHandler struct {
	tmdbClient MovieClient
}

// NewMovieHandler creates a new MovieHandler instance
func NewMovieHandler(tmdbClient MovieClient) *MovieHandler {
	return &MovieHandler{
		tmdbClient: tmdbClient,
	}
}

// GetMovieDetails handles GET /api/v1/movies/{id} requests
func (h *MovieHandler) GetMovieDetails(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only allow GET requests
	if r.Method != http.MethodGet {
		writeErrorResponse(w, http.StatusMethodNotAllowed, "method_not_allowed", "Only GET method is allowed")
		return
	}

	// Extract movie ID from URL path
	vars := mux.Vars(r)
	movieIDStr, exists := vars["id"]
	if !exists {
		writeErrorResponse(w, http.StatusBadRequest, "missing_parameter", "Movie ID is required")
		return
	}

	// Parse movie ID
	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil || movieID <= 0 {
		writeErrorResponse(w, http.StatusBadRequest, "invalid_parameter", "Movie ID must be a positive integer")
		return
	}

	log.Printf("Fetching movie details for ID: %d", movieID)

	// Get movie details from TMDb API
	movieDetails, err := h.tmdbClient.GetMovieDetails(r.Context(), movieID)
	if err != nil {
		log.Printf("Failed to get movie details for ID %d: %v", movieID, err)
		
		// Check if it's a TMDb API error
		if tmdbErr, ok := err.(*services.TMDbError); ok {
			if tmdbErr.StatusCode == 404 {
				writeErrorResponse(w, http.StatusNotFound, "movie_not_found", fmt.Sprintf("Movie with ID %d not found", movieID))
				return
			}
		}
		
		writeErrorResponse(w, http.StatusInternalServerError, "api_error", "Failed to retrieve movie details")
		return
	}

	log.Printf("Successfully retrieved movie details: %s (%d)", movieDetails.Title, movieDetails.ID)

	// Return movie details
	writeJSONResponse(w, http.StatusOK, movieDetails)
}

// GetMovieCredits handles GET /api/v1/movies/{id}/credits requests
func (h *MovieHandler) GetMovieCredits(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only allow GET requests
	if r.Method != http.MethodGet {
		writeErrorResponse(w, http.StatusMethodNotAllowed, "method_not_allowed", "Only GET method is allowed")
		return
	}

	// Extract movie ID from URL path
	vars := mux.Vars(r)
	movieIDStr, exists := vars["id"]
	if !exists {
		writeErrorResponse(w, http.StatusBadRequest, "missing_parameter", "Movie ID is required")
		return
	}

	// Parse movie ID
	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil || movieID <= 0 {
		writeErrorResponse(w, http.StatusBadRequest, "invalid_parameter", "Movie ID must be a positive integer")
		return
	}

	log.Printf("Fetching movie credits for ID: %d", movieID)

	// Get movie credits from TMDb API
	movieCredits, err := h.tmdbClient.GetMovieCredits(r.Context(), movieID)
	if err != nil {
		log.Printf("Failed to get movie credits for ID %d: %v", movieID, err)
		
		// Check if it's a TMDb API error
		if tmdbErr, ok := err.(*services.TMDbError); ok {
			if tmdbErr.StatusCode == 404 {
				writeErrorResponse(w, http.StatusNotFound, "movie_not_found", fmt.Sprintf("Movie with ID %d not found", movieID))
				return
			}
		}
		
		writeErrorResponse(w, http.StatusInternalServerError, "api_error", "Failed to retrieve movie credits")
		return
	}

	log.Printf("Successfully retrieved movie credits for movie ID %d: %d cast, %d crew", 
		movieCredits.ID, len(movieCredits.Cast), len(movieCredits.Crew))

	// Return movie credits
	writeJSONResponse(w, http.StatusOK, movieCredits)
}

// GetMovieReviews handles GET /api/v1/movies/{id}/reviews requests
func (h *MovieHandler) GetMovieReviews(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only allow GET requests
	if r.Method != http.MethodGet {
		writeErrorResponse(w, http.StatusMethodNotAllowed, "method_not_allowed", "Only GET method is allowed")
		return
	}

	// Extract movie ID from URL path
	vars := mux.Vars(r)
	movieIDStr, exists := vars["id"]
	if !exists {
		writeErrorResponse(w, http.StatusBadRequest, "missing_parameter", "Movie ID is required")
		return
	}

	// Parse movie ID
	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil || movieID <= 0 {
		writeErrorResponse(w, http.StatusBadRequest, "invalid_parameter", "Movie ID must be a positive integer")
		return
	}

	// Parse page parameter
	page := 1
	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if parsedPage, err := strconv.Atoi(pageStr); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	log.Printf("Fetching movie reviews for ID: %d, page: %d", movieID, page)

	// Get movie reviews from TMDb API
	movieReviews, err := h.tmdbClient.GetMovieReviews(r.Context(), movieID, page)
	if err != nil {
		log.Printf("Failed to get movie reviews for ID %d: %v", movieID, err)
		
		// Check if it's a TMDb API error
		if tmdbErr, ok := err.(*services.TMDbError); ok {
			if tmdbErr.StatusCode == 404 {
				writeErrorResponse(w, http.StatusNotFound, "movie_not_found", fmt.Sprintf("Movie with ID %d not found", movieID))
				return
			}
		}
		
		writeErrorResponse(w, http.StatusInternalServerError, "api_error", "Failed to retrieve movie reviews")
		return
	}

	log.Printf("Successfully retrieved movie reviews for movie ID %d: %d reviews (page %d/%d)", 
		movieReviews.ID, len(movieReviews.Results), movieReviews.Page, movieReviews.TotalPages)

	// Return movie reviews
	writeJSONResponse(w, http.StatusOK, movieReviews)
}

