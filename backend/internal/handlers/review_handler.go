// Package handlers provides HTTP handlers for review-related endpoints.
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

// ReviewClient defines the interface for review-related TMDb operations
type ReviewClient interface {
	GetMovieReviews(ctx context.Context, movieID int, page int) (*models.MovieReviews, error)
	GetTVShowReviews(ctx context.Context, tvID int, page int) (*models.TVReviews, error)
}

// ReviewHandler handles review-related HTTP requests
type ReviewHandler struct {
	tmdbClient ReviewClient
}

// NewReviewHandler creates a new ReviewHandler instance
func NewReviewHandler(tmdbClient ReviewClient) *ReviewHandler {
	return &ReviewHandler{
		tmdbClient: tmdbClient,
	}
}

// GetMovieReviews handles GET /api/v1/movies/{id}/reviews requests
func (h *ReviewHandler) GetMovieReviews(w http.ResponseWriter, r *http.Request) {
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

// GetTVReviews handles GET /api/v1/tv/{id}/reviews requests
func (h *ReviewHandler) GetTVReviews(w http.ResponseWriter, r *http.Request) {
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

	// Extract TV show ID from URL path
	vars := mux.Vars(r)
	tvIDStr, exists := vars["id"]
	if !exists {
		writeErrorResponse(w, http.StatusBadRequest, "missing_parameter", "TV show ID is required")
		return
	}

	// Parse TV show ID
	tvID, err := strconv.Atoi(tvIDStr)
	if err != nil || tvID <= 0 {
		writeErrorResponse(w, http.StatusBadRequest, "invalid_parameter", "TV show ID must be a positive integer")
		return
	}

	// Parse page parameter
	page := 1
	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if parsedPage, err := strconv.Atoi(pageStr); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	log.Printf("Fetching TV show reviews for ID: %d, page: %d", tvID, page)

	// Get TV show reviews from TMDb API
	tvReviews, err := h.tmdbClient.GetTVShowReviews(r.Context(), tvID, page)
	if err != nil {
		log.Printf("Failed to get TV show reviews for ID %d: %v", tvID, err)
		
		// Check if it's a TMDb API error
		if tmdbErr, ok := err.(*services.TMDbError); ok {
			if tmdbErr.StatusCode == 404 {
				writeErrorResponse(w, http.StatusNotFound, "tv_not_found", fmt.Sprintf("TV show with ID %d not found", tvID))
				return
			}
		}
		
		writeErrorResponse(w, http.StatusInternalServerError, "api_error", "Failed to retrieve TV show reviews")
		return
	}

	log.Printf("Successfully retrieved TV show reviews for TV ID %d: %d reviews (page %d/%d)", 
		tvReviews.ID, len(tvReviews.Results), tvReviews.Page, tvReviews.TotalPages)

	// Return TV show reviews
	writeJSONResponse(w, http.StatusOK, tvReviews)
}