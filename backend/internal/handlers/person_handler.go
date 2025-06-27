// Package handlers provides HTTP handlers for person-related endpoints.
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

// PersonClient defines the interface for person-related TMDb operations
type PersonClient interface {
	GetPersonDetails(ctx context.Context, personID int) (*models.PersonDetails, error)
	GetPersonMovieCredits(ctx context.Context, personID int) (*models.PersonMovieCredits, error)
	GetPersonTVCredits(ctx context.Context, personID int) (*models.PersonTVCredits, error)
	GetPersonCombinedCredits(ctx context.Context, personID int) (*models.PersonCombinedCredits, error)
}

// PersonHandler handles person-related HTTP requests
type PersonHandler struct {
	tmdbClient PersonClient
}

// NewPersonHandler creates a new PersonHandler instance
func NewPersonHandler(tmdbClient PersonClient) *PersonHandler {
	return &PersonHandler{
		tmdbClient: tmdbClient,
	}
}

// GetPersonDetails handles GET /api/v1/people/{id} requests
func (h *PersonHandler) GetPersonDetails(w http.ResponseWriter, r *http.Request) {
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

	// Extract person ID from URL path
	vars := mux.Vars(r)
	personIDStr, exists := vars["id"]
	if !exists {
		writeErrorResponse(w, http.StatusBadRequest, "missing_parameter", "Person ID is required")
		return
	}

	// Parse person ID
	personID, err := strconv.Atoi(personIDStr)
	if err != nil || personID <= 0 {
		writeErrorResponse(w, http.StatusBadRequest, "invalid_parameter", "Person ID must be a positive integer")
		return
	}

	log.Printf("Fetching person details for ID: %d", personID)

	// Get person details from TMDb API
	personDetails, err := h.tmdbClient.GetPersonDetails(r.Context(), personID)
	if err != nil {
		log.Printf("Failed to get person details for ID %d: %v", personID, err)
		
		// Check if it's a TMDb API error
		if tmdbErr, ok := err.(*services.TMDbError); ok {
			if tmdbErr.StatusCode == 404 {
				writeErrorResponse(w, http.StatusNotFound, "person_not_found", fmt.Sprintf("Person with ID %d not found", personID))
				return
			}
		}
		
		writeErrorResponse(w, http.StatusInternalServerError, "api_error", "Failed to retrieve person details")
		return
	}

	log.Printf("Successfully retrieved person details: %s (ID: %d)", personDetails.Name, personDetails.ID)

	// Return person details
	writeJSONResponse(w, http.StatusOK, personDetails)
}

// GetPersonMovieCredits handles GET /api/v1/people/{id}/movie_credits requests
func (h *PersonHandler) GetPersonMovieCredits(w http.ResponseWriter, r *http.Request) {
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

	// Extract person ID from URL path
	vars := mux.Vars(r)
	personIDStr, exists := vars["id"]
	if !exists {
		writeErrorResponse(w, http.StatusBadRequest, "missing_parameter", "Person ID is required")
		return
	}

	// Parse person ID
	personID, err := strconv.Atoi(personIDStr)
	if err != nil || personID <= 0 {
		writeErrorResponse(w, http.StatusBadRequest, "invalid_parameter", "Person ID must be a positive integer")
		return
	}

	log.Printf("Fetching person movie credits for ID: %d", personID)

	// Get person movie credits from TMDb API
	movieCredits, err := h.tmdbClient.GetPersonMovieCredits(r.Context(), personID)
	if err != nil {
		log.Printf("Failed to get person movie credits for ID %d: %v", personID, err)
		
		// Check if it's a TMDb API error
		if tmdbErr, ok := err.(*services.TMDbError); ok {
			if tmdbErr.StatusCode == 404 {
				writeErrorResponse(w, http.StatusNotFound, "person_not_found", fmt.Sprintf("Person with ID %d not found", personID))
				return
			}
		}
		
		writeErrorResponse(w, http.StatusInternalServerError, "api_error", "Failed to retrieve person movie credits")
		return
	}

	log.Printf("Successfully retrieved person movie credits for person ID %d: %d cast, %d crew", 
		movieCredits.ID, len(movieCredits.Cast), len(movieCredits.Crew))

	// Return person movie credits
	writeJSONResponse(w, http.StatusOK, movieCredits)
}

// GetPersonTVCredits handles GET /api/v1/people/{id}/tv_credits requests
func (h *PersonHandler) GetPersonTVCredits(w http.ResponseWriter, r *http.Request) {
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

	// Extract person ID from URL path
	vars := mux.Vars(r)
	personIDStr, exists := vars["id"]
	if !exists {
		writeErrorResponse(w, http.StatusBadRequest, "missing_parameter", "Person ID is required")
		return
	}

	// Parse person ID
	personID, err := strconv.Atoi(personIDStr)
	if err != nil || personID <= 0 {
		writeErrorResponse(w, http.StatusBadRequest, "invalid_parameter", "Person ID must be a positive integer")
		return
	}

	log.Printf("Fetching person TV credits for ID: %d", personID)

	// Get person TV credits from TMDb API
	tvCredits, err := h.tmdbClient.GetPersonTVCredits(r.Context(), personID)
	if err != nil {
		log.Printf("Failed to get person TV credits for ID %d: %v", personID, err)
		
		// Check if it's a TMDb API error
		if tmdbErr, ok := err.(*services.TMDbError); ok {
			if tmdbErr.StatusCode == 404 {
				writeErrorResponse(w, http.StatusNotFound, "person_not_found", fmt.Sprintf("Person with ID %d not found", personID))
				return
			}
		}
		
		writeErrorResponse(w, http.StatusInternalServerError, "api_error", "Failed to retrieve person TV credits")
		return
	}

	log.Printf("Successfully retrieved person TV credits for person ID %d: %d cast, %d crew", 
		tvCredits.ID, len(tvCredits.Cast), len(tvCredits.Crew))

	// Return person TV credits
	writeJSONResponse(w, http.StatusOK, tvCredits)
}

// GetPersonCombinedCredits handles GET /api/v1/people/{id}/combined_credits requests
func (h *PersonHandler) GetPersonCombinedCredits(w http.ResponseWriter, r *http.Request) {
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

	// Extract person ID from URL path
	vars := mux.Vars(r)
	personIDStr, exists := vars["id"]
	if !exists {
		writeErrorResponse(w, http.StatusBadRequest, "missing_parameter", "Person ID is required")
		return
	}

	// Parse person ID
	personID, err := strconv.Atoi(personIDStr)
	if err != nil || personID <= 0 {
		writeErrorResponse(w, http.StatusBadRequest, "invalid_parameter", "Person ID must be a positive integer")
		return
	}

	log.Printf("Fetching person combined credits for ID: %d", personID)

	// Get person combined credits from TMDb API
	combinedCredits, err := h.tmdbClient.GetPersonCombinedCredits(r.Context(), personID)
	if err != nil {
		log.Printf("Failed to get person combined credits for ID %d: %v", personID, err)
		
		// Check if it's a TMDb API error
		if tmdbErr, ok := err.(*services.TMDbError); ok {
			if tmdbErr.StatusCode == 404 {
				writeErrorResponse(w, http.StatusNotFound, "person_not_found", fmt.Sprintf("Person with ID %d not found", personID))
				return
			}
		}
		
		writeErrorResponse(w, http.StatusInternalServerError, "api_error", "Failed to retrieve person combined credits")
		return
	}

	log.Printf("Successfully retrieved person combined credits for person ID %d: %d cast, %d crew", 
		combinedCredits.ID, len(combinedCredits.Cast), len(combinedCredits.Crew))

	// Return person combined credits
	writeJSONResponse(w, http.StatusOK, combinedCredits)
}