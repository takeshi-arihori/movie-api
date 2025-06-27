// Package handlers provides HTTP handlers for the Movie API service.
package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/takeshi-arihori/movie-api/internal/config"
	"github.com/takeshi-arihori/movie-api/internal/models"
	"github.com/takeshi-arihori/movie-api/internal/services"
)

// createTestSearchHandler creates a SearchHandler configured for testing
func createTestSearchHandler(mockServer *httptest.Server) *SearchHandler {
	cfg := &config.Config{
		TMDb: config.TMDbConfig{
			APIKey:  "test-api-key",
			BaseURL: mockServer.URL,
		},
	}
	tmdbClient := services.NewTMDbClient(cfg)
	return NewSearchHandler(tmdbClient)
}

// createMockTMDbServer creates a test HTTP server with predefined TMDb responses
func createMockTMDbServer(t *testing.T, responses map[string]interface{}) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract endpoint from URL path
		endpoint := r.URL.Path

		// Check if API key is present
		if r.URL.Query().Get("api_key") != "test-api-key" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				StatusCode:    401,
				StatusMessage: "Invalid API key",
				Success:       false,
			})
			return
		}

		// Find response for endpoint
		response, exists := responses[endpoint]
		if !exists {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				StatusCode:    404,
				StatusMessage: "Resource not found",
				Success:       false,
			})
			return
		}

		// Return mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
}

// TestSearchHandler_Search tests the main search functionality
func TestSearchHandler_Search(t *testing.T) {
	// Mock multi-search response
	mockResponse := models.MultiSearchResponse{
		Page: 1,
		Results: []models.MultiSearchResult{
			{
				ID:               550,
				MediaType:        models.SearchItemTypeMovie,
				Popularity:       61.416,
				Title:            stringPtr("Fight Club"),
				OriginalTitle:    stringPtr("Fight Club"),
				OriginalLanguage: stringPtr("en"),
				Overview:         stringPtr("A ticking-time-bomb insomniac..."),
				VoteAverage:      float64Ptr(8.4),
				VoteCount:        intPtr(26280),
			},
		},
		TotalPages:   1,
		TotalResults: 1,
	}

	responses := map[string]interface{}{
		"/search/multi": mockResponse,
	}

	server := createMockTMDbServer(t, responses)
	defer server.Close()

	handler := createTestSearchHandler(server)

	// Test successful search
	req := httptest.NewRequest("GET", "/api/v1/search?query=Fight+Club", nil)
	w := httptest.NewRecorder()

	handler.Search(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response models.APISearchResponse
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Query != "Fight Club" {
		t.Errorf("Expected query 'Fight Club', got '%s'", response.Query)
	}
	if len(response.Results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(response.Results))
	}
	if response.Results[0].MediaType != models.SearchItemTypeMovie {
		t.Errorf("Expected media type movie, got %s", response.Results[0].MediaType)
	}
}

// TestSearchHandler_SearchWithType tests type-specific search
func TestSearchHandler_SearchWithType(t *testing.T) {
	// Mock movie search response
	movieResponse := models.MovieSearchResponse{
		Page: 1,
		Results: []models.Movie{
			{
				ID:               550,
				Title:            "Fight Club",
				OriginalTitle:    "Fight Club",
				OriginalLanguage: "en",
				Overview:         "A ticking-time-bomb insomniac...",
				VoteAverage:      8.4,
				VoteCount:        26280,
			},
		},
		TotalPages:   1,
		TotalResults: 1,
	}

	responses := map[string]interface{}{
		"/search/movie": movieResponse,
	}

	server := createMockTMDbServer(t, responses)
	defer server.Close()

	handler := createTestSearchHandler(server)

	// Test movie-specific search
	req := httptest.NewRequest("GET", "/api/v1/search?query=Fight+Club&type=movie", nil)
	w := httptest.NewRecorder()

	handler.Search(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response models.APISearchResponse
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Type != "movie" {
		t.Errorf("Expected type 'movie', got '%s'", response.Type)
	}
	if len(response.Results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(response.Results))
	}
}

// TestSearchHandler_SearchValidation tests request validation
func TestSearchHandler_SearchValidation(t *testing.T) {
	server := createMockTMDbServer(t, map[string]interface{}{})
	defer server.Close()

	handler := createTestSearchHandler(server)

	testCases := []struct {
		name           string
		query          string
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "Empty query",
			query:          "",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "validation_error",
		},
		{
			name:           "Invalid type",
			query:          "query=test&type=invalid",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "validation_error",
		},
		{
			name:           "Invalid page",
			query:          "query=test&page=invalid",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid_request",
		},
		{
			name:           "Negative page",
			query:          "query=test&page=-1",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid_request",
		},
		{
			name:           "Invalid year",
			query:          "query=test&year=1800",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid_request",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/v1/search?"+tc.query, nil)
			w := httptest.NewRecorder()

			handler.Search(w, req)

			if w.Code != tc.expectedStatus {
				t.Errorf("Expected status %d, got %d", tc.expectedStatus, w.Code)
			}

			var errorResp ErrorResponse
			if err := json.NewDecoder(w.Body).Decode(&errorResp); err != nil {
				t.Fatalf("Failed to decode error response: %v", err)
			}

			if errorResp.Error != tc.expectedError {
				t.Errorf("Expected error type '%s', got '%s'", tc.expectedError, errorResp.Error)
			}
		})
	}
}

// TestSearchHandler_MethodNotAllowed tests HTTP method validation
func TestSearchHandler_MethodNotAllowed(t *testing.T) {
	server := createMockTMDbServer(t, map[string]interface{}{})
	defer server.Close()

	handler := createTestSearchHandler(server)

	// Test POST method (should be rejected)
	req := httptest.NewRequest("POST", "/api/v1/search", nil)
	w := httptest.NewRecorder()

	handler.Search(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d, got %d", http.StatusMethodNotAllowed, w.Code)
	}

	var errorResp ErrorResponse
	if err := json.NewDecoder(w.Body).Decode(&errorResp); err != nil {
		t.Fatalf("Failed to decode error response: %v", err)
	}

	if errorResp.Error != "method_not_allowed" {
		t.Errorf("Expected error type 'method_not_allowed', got '%s'", errorResp.Error)
	}
}

// TestSearchHandler_OptionsRequest tests CORS preflight requests
func TestSearchHandler_OptionsRequest(t *testing.T) {
	server := createMockTMDbServer(t, map[string]interface{}{})
	defer server.Close()

	handler := createTestSearchHandler(server)

	// Test OPTIONS request
	req := httptest.NewRequest("OPTIONS", "/api/v1/search", nil)
	w := httptest.NewRecorder()

	handler.Search(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Check CORS headers
	if origin := w.Header().Get("Access-Control-Allow-Origin"); origin != "*" {
		t.Errorf("Expected Access-Control-Allow-Origin '*', got '%s'", origin)
	}
	if methods := w.Header().Get("Access-Control-Allow-Methods"); methods != "GET, OPTIONS" {
		t.Errorf("Expected Access-Control-Allow-Methods 'GET, OPTIONS', got '%s'", methods)
	}
}

// TestSearchHandler_HealthCheck tests health check endpoint
func TestSearchHandler_HealthCheck(t *testing.T) {
	server := createMockTMDbServer(t, map[string]interface{}{})
	defer server.Close()

	handler := createTestSearchHandler(server)

	req := httptest.NewRequest("GET", "/api/v1/health", nil)
	w := httptest.NewRecorder()

	handler.HealthCheck(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if status, ok := response["status"].(string); !ok || status != "healthy" {
		t.Errorf("Expected status 'healthy', got %v", response["status"])
	}
	if service, ok := response["service"].(string); !ok || service != "movie-api" {
		t.Errorf("Expected service 'movie-api', got %v", response["service"])
	}
}

// TestSearchHandler_GetSearchSuggestions tests search suggestions endpoint
func TestSearchHandler_GetSearchSuggestions(t *testing.T) {
	server := createMockTMDbServer(t, map[string]interface{}{})
	defer server.Close()

	handler := createTestSearchHandler(server)

	req := httptest.NewRequest("GET", "/api/v1/search/suggestions", nil)
	w := httptest.NewRecorder()

	handler.GetSearchSuggestions(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if suggestions, ok := response["suggestions"].([]interface{}); !ok || len(suggestions) != 0 {
		t.Errorf("Expected empty suggestions array, got %v", response["suggestions"])
	}
}

// TestSearchHandler_TMDbError tests TMDb API error handling
func TestSearchHandler_TMDbError(t *testing.T) {
	// Create server that returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			StatusCode:    500,
			StatusMessage: "Internal server error",
			Success:       false,
		})
	}))
	defer server.Close()

	handler := createTestSearchHandler(server)

	req := httptest.NewRequest("GET", "/api/v1/search?query=test", nil)
	w := httptest.NewRecorder()

	handler.Search(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, w.Code)
	}

	var errorResp ErrorResponse
	if err := json.NewDecoder(w.Body).Decode(&errorResp); err != nil {
		t.Fatalf("Failed to decode error response: %v", err)
	}

	if errorResp.Error != "search_error" {
		t.Errorf("Expected error type 'search_error', got '%s'", errorResp.Error)
	}
}

// TestParseSearchRequest tests search request parsing
func TestParseSearchRequest(t *testing.T) {
	// Create a mock server for the handler
	server := createMockTMDbServer(t, map[string]interface{}{})
	defer server.Close()
	
	handler := createTestSearchHandler(server)

	testCases := []struct {
		name        string
		queryParams string
		expected    *models.SearchRequest
		expectError bool
	}{
		{
			name:        "Valid request with all parameters",
			queryParams: "query=test&type=movie&page=2&language=en-US&year=2020",
			expected: &models.SearchRequest{
				Query:    "test",
				Type:     "movie",
				Page:     2,
				Language: "en-US",
				Year:     2020,
			},
			expectError: false,
		},
		{
			name:        "Valid request with minimal parameters",
			queryParams: "query=test",
			expected: &models.SearchRequest{
				Query: "test",
			},
			expectError: false,
		},
		{
			name:        "Invalid page parameter",
			queryParams: "query=test&page=invalid",
			expectError: true,
		},
		{
			name:        "Invalid year parameter",
			queryParams: "query=test&year=1800",
			expectError: true,
		},
		{
			name:        "Zero page parameter",
			queryParams: "query=test&page=0",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create request with query parameters
			req := httptest.NewRequest("GET", "/api/v1/search?"+tc.queryParams, nil)

			result, err := handler.parseSearchRequest(req)

			if tc.expectError {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result.Query != tc.expected.Query {
				t.Errorf("Expected query '%s', got '%s'", tc.expected.Query, result.Query)
			}
			if result.Type != tc.expected.Type {
				t.Errorf("Expected type '%s', got '%s'", tc.expected.Type, result.Type)
			}
			if result.Page != tc.expected.Page {
				t.Errorf("Expected page %d, got %d", tc.expected.Page, result.Page)
			}
			if result.Language != tc.expected.Language {
				t.Errorf("Expected language '%s', got '%s'", tc.expected.Language, result.Language)
			}
			if result.Year != tc.expected.Year {
				t.Errorf("Expected year %d, got %d", tc.expected.Year, result.Year)
			}
		})
	}
}

// TestSearchRequestDefaults tests default value setting
func TestSearchRequestDefaults(t *testing.T) {
	// Create a mock server for the handler
	server := createMockTMDbServer(t, map[string]interface{}{})
	defer server.Close()
	
	handler := createTestSearchHandler(server)

	req := httptest.NewRequest("GET", "/api/v1/search?query=test", nil)
	searchReq, err := handler.parseSearchRequest(req)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	searchReq.SetDefaults()

	if searchReq.Type != "all" {
		t.Errorf("Expected default type 'all', got '%s'", searchReq.Type)
	}
	if searchReq.Page != 1 {
		t.Errorf("Expected default page 1, got %d", searchReq.Page)
	}
	if searchReq.Language != "ja-JP" {
		t.Errorf("Expected default language 'ja-JP', got '%s'", searchReq.Language)
	}
}

// TestWriteJSONResponse tests JSON response writing
func TestWriteJSONResponse(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]string{"test": "value"}

	writeJSONResponse(w, http.StatusOK, data)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got '%s'", contentType)
	}

	var response map[string]string
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["test"] != "value" {
		t.Errorf("Expected response value 'value', got '%s'", response["test"])
	}
}

// TestWriteErrorResponse tests error response writing
func TestWriteErrorResponse(t *testing.T) {
	w := httptest.NewRecorder()

	writeErrorResponse(w, http.StatusBadRequest, "test_error", "Test error message")

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	var errorResp ErrorResponse
	if err := json.NewDecoder(w.Body).Decode(&errorResp); err != nil {
		t.Fatalf("Failed to decode error response: %v", err)
	}

	if errorResp.Error != "test_error" {
		t.Errorf("Expected error type 'test_error', got '%s'", errorResp.Error)
	}
	if errorResp.Message != "Test error message" {
		t.Errorf("Expected message 'Test error message', got '%s'", errorResp.Message)
	}
	if errorResp.Code != http.StatusBadRequest {
		t.Errorf("Expected code %d, got %d", http.StatusBadRequest, errorResp.Code)
	}
}

// Helper functions for pointer types
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func float64Ptr(f float64) *float64 {
	return &f
}