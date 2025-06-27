// Package handlers provides HTTP handlers for the Movie API service.
package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/takeshi-arihori/movie-api/internal/models"
	"github.com/takeshi-arihori/movie-api/internal/services"
)

// SearchHandler handles search-related HTTP requests
type SearchHandler struct {
	tmdbClient *services.TMDbClient
}

// NewSearchHandler creates a new SearchHandler instance
func NewSearchHandler(tmdbClient *services.TMDbClient) *SearchHandler {
	return &SearchHandler{
		tmdbClient: tmdbClient,
	}
}


// Search handles multi-search requests
// GET /api/v1/search?query=<query>&type=<type>&page=<page>&language=<language>
func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
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

	// Parse query parameters
	searchReq, err := h.parseSearchRequest(r)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "invalid_request", err.Error())
		return
	}

	// Validate request
	if err := searchReq.Validate(); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "validation_error", err.Error())
		return
	}

	// Set defaults
	searchReq.SetDefaults()

	log.Printf("Search request: query=%s, type=%s, page=%d, language=%s", 
		searchReq.Query, searchReq.Type, searchReq.Page, searchReq.Language)

	// Perform search based on type
	var result *models.MultiSearchResponse
	if searchReq.Type == "all" {
		result, err = h.tmdbClient.MultiSearch(r.Context(), searchReq.Query, searchReq.Page, searchReq.Language)
	} else {
		result, err = h.tmdbClient.SearchByType(r.Context(), searchReq.Type, searchReq.Query, searchReq.Page, searchReq.Language)
	}

	if err != nil {
		log.Printf("Search failed: %v", err)
		writeErrorResponse(w, http.StatusInternalServerError, "search_error", "Failed to perform search")
		return
	}

	// Build response
	response := models.APISearchResponse{
		Query:        searchReq.Query,
		Type:         searchReq.Type,
		Page:         result.Page,
		TotalPages:   result.TotalPages,
		TotalResults: result.TotalResults,
		Results:      result.Results,
		Language:     searchReq.Language,
	}

	log.Printf("Search completed: found %d results (page %d/%d)", 
		len(response.Results), response.Page, response.TotalPages)

	writeJSONResponse(w, http.StatusOK, response)
}

// parseSearchRequest parses HTTP request parameters into SearchRequest
func (h *SearchHandler) parseSearchRequest(r *http.Request) (*models.SearchRequest, error) {
	query := r.URL.Query()

	searchReq := &models.SearchRequest{
		Query:    strings.TrimSpace(query.Get("query")),
		Type:     strings.ToLower(strings.TrimSpace(query.Get("type"))),
		Language: strings.TrimSpace(query.Get("language")),
	}

	// Parse page parameter
	if pageStr := query.Get("page"); pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			return nil, fmt.Errorf("invalid page parameter: must be a positive integer")
		}
		searchReq.Page = page
	}

	// Parse year parameter (for movies)
	if yearStr := query.Get("year"); yearStr != "" {
		year, err := strconv.Atoi(yearStr)
		if err != nil || year < 1900 || year > 2100 {
			return nil, fmt.Errorf("invalid year parameter: must be between 1900 and 2100")
		}
		searchReq.Year = year
	}

	return searchReq, nil
}

// HealthCheck provides a health check endpoint
func (h *SearchHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	healthStatus := map[string]interface{}{
		"status":    "healthy",
		"service":   "movie-api",
		"version":   "1.0.0",
		"timestamp": "2023-01-01T00:00:00Z", // In production, use actual timestamp
		"endpoints": map[string]string{
			"search": "/api/v1/search",
		},
	}

	writeJSONResponse(w, http.StatusOK, healthStatus)
}

// GetSearchSuggestions provides search suggestions (placeholder for future implementation)
func (h *SearchHandler) GetSearchSuggestions(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// For now, return a placeholder response
	suggestions := map[string]interface{}{
		"suggestions": []string{},
		"message":     "Search suggestions feature coming soon",
	}

	writeJSONResponse(w, http.StatusOK, suggestions)
}