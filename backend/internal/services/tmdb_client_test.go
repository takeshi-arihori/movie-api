package services

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/takeshi-arihori/movie-api/internal/config"
	"github.com/takeshi-arihori/movie-api/internal/models"
)

// createTestClient creates a TMDb client configured for testing
func createTestClient(serverURL string) *TMDbClient {
	cfg := &config.Config{
		TMDb: config.TMDbConfig{
			APIKey:  "test-api-key",
			BaseURL: serverURL,
		},
	}
	return NewTMDbClient(cfg)
}

// createMockServer creates a test HTTP server with predefined responses
func createMockServer(t *testing.T, responses map[string]interface{}) *httptest.Server {
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

// TestNewTMDbClient tests client creation
func TestNewTMDbClient(t *testing.T) {
	cfg := &config.Config{
		TMDb: config.TMDbConfig{
			APIKey:  "test-key",
			BaseURL: "https://api.themoviedb.org/3",
		},
	}

	client := NewTMDbClient(cfg)

	if client.apiKey != "test-key" {
		t.Errorf("Expected API key 'test-key', got '%s'", client.apiKey)
	}
	if client.baseURL != "https://api.themoviedb.org/3" {
		t.Errorf("Expected base URL 'https://api.themoviedb.org/3', got '%s'", client.baseURL)
	}
	if client.httpClient == nil {
		t.Error("HTTP client should not be nil")
	}
}

// TestSearchMovies tests movie search functionality
func TestSearchMovies(t *testing.T) {
	// Mock response data
	mockResponse := models.MovieSearchResponse{
		Page:         1,
		Results:      []models.Movie{
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
		"/search/movie": mockResponse,
	}

	server := createMockServer(t, responses)
	defer server.Close()

	client := createTestClient(server.URL)
	ctx := context.Background()

	// Test successful search
	result, err := client.SearchMovies(ctx, "Fight Club", 1)
	if err != nil {
		t.Fatalf("SearchMovies failed: %v", err)
	}

	if len(result.Results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(result.Results))
	}

	movie := result.Results[0]
	if movie.ID != 550 {
		t.Errorf("Expected movie ID 550, got %d", movie.ID)
	}
	if movie.Title != "Fight Club" {
		t.Errorf("Expected title 'Fight Club', got '%s'", movie.Title)
	}

	// Test empty query
	_, err = client.SearchMovies(ctx, "", 1)
	if err == nil {
		t.Error("Expected error for empty query, got nil")
	}
}

// TestSearchTVShows tests TV show search functionality
func TestSearchTVShows(t *testing.T) {
	// Mock response data
	mockResponse := models.TVSearchResponse{
		Page:         1,
		Results:      []models.TVShow{
			{
				ID:               1396,
				Name:             "Breaking Bad",
				OriginalName:     "Breaking Bad",
				OriginalLanguage: "en",
				Overview:         "A high school chemistry teacher...",
				VoteAverage:      8.9,
				VoteCount:        12859,
			},
		},
		TotalPages:   1,
		TotalResults: 1,
	}

	responses := map[string]interface{}{
		"/search/tv": mockResponse,
	}

	server := createMockServer(t, responses)
	defer server.Close()

	client := createTestClient(server.URL)
	ctx := context.Background()

	// Test successful search
	result, err := client.SearchTVShows(ctx, "Breaking Bad", 1)
	if err != nil {
		t.Fatalf("SearchTVShows failed: %v", err)
	}

	if len(result.Results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(result.Results))
	}

	tvShow := result.Results[0]
	if tvShow.ID != 1396 {
		t.Errorf("Expected TV show ID 1396, got %d", tvShow.ID)
	}
	if tvShow.Name != "Breaking Bad" {
		t.Errorf("Expected name 'Breaking Bad', got '%s'", tvShow.Name)
	}

	// Test empty query
	_, err = client.SearchTVShows(ctx, "", 1)
	if err == nil {
		t.Error("Expected error for empty query, got nil")
	}
}

// TestGetMovieDetails tests movie details retrieval
func TestGetMovieDetails(t *testing.T) {
	// Mock response data
	mockResponse := models.MovieDetails{
		ID:               550,
		Title:            "Fight Club",
		OriginalTitle:    "Fight Club",
		OriginalLanguage: "en",
		Overview:         stringPtr("A ticking-time-bomb insomniac..."),
		VoteAverage:      8.4,
		VoteCount:        26280,
		Runtime:          intPtr(139),
		Status:           "Released",
		Genres: []models.Genre{
			{ID: 18, Name: "Drama"},
			{ID: 53, Name: "Thriller"},
		},
	}

	responses := map[string]interface{}{
		"/movie/550": mockResponse,
	}

	server := createMockServer(t, responses)
	defer server.Close()

	client := createTestClient(server.URL)
	ctx := context.Background()

	// Test successful retrieval
	result, err := client.GetMovieDetails(ctx, 550)
	if err != nil {
		t.Fatalf("GetMovieDetails failed: %v", err)
	}

	if result.ID != 550 {
		t.Errorf("Expected movie ID 550, got %d", result.ID)
	}
	if result.Title != "Fight Club" {
		t.Errorf("Expected title 'Fight Club', got '%s'", result.Title)
	}
	if len(result.Genres) != 2 {
		t.Errorf("Expected 2 genres, got %d", len(result.Genres))
	}

	// Test invalid movie ID
	_, err = client.GetMovieDetails(ctx, 0)
	if err == nil {
		t.Error("Expected error for invalid movie ID, got nil")
	}

	_, err = client.GetMovieDetails(ctx, -1)
	if err == nil {
		t.Error("Expected error for negative movie ID, got nil")
	}
}

// TestGetTVShowDetails tests TV show details retrieval
func TestGetTVShowDetails(t *testing.T) {
	// Mock response data
	mockResponse := models.TVShowDetails{
		ID:               1396,
		Name:             "Breaking Bad",
		OriginalName:     "Breaking Bad",
		OriginalLanguage: "en",
		Overview:         stringPtr("A high school chemistry teacher..."),
		VoteAverage:      8.9,
		VoteCount:        12859,
		NumberOfSeasons:  5,
		NumberOfEpisodes: 62,
		Status:           "Ended",
		Type:             "Scripted",
		Genres: []models.Genre{
			{ID: 18, Name: "Drama"},
			{ID: 80, Name: "Crime"},
		},
	}

	responses := map[string]interface{}{
		"/tv/1396": mockResponse,
	}

	server := createMockServer(t, responses)
	defer server.Close()

	client := createTestClient(server.URL)
	ctx := context.Background()

	// Test successful retrieval
	result, err := client.GetTVShowDetails(ctx, 1396)
	if err != nil {
		t.Fatalf("GetTVShowDetails failed: %v", err)
	}

	if result.ID != 1396 {
		t.Errorf("Expected TV show ID 1396, got %d", result.ID)
	}
	if result.Name != "Breaking Bad" {
		t.Errorf("Expected name 'Breaking Bad', got '%s'", result.Name)
	}
	if result.NumberOfSeasons != 5 {
		t.Errorf("Expected 5 seasons, got %d", result.NumberOfSeasons)
	}

	// Test invalid TV show ID
	_, err = client.GetTVShowDetails(ctx, 0)
	if err == nil {
		t.Error("Expected error for invalid TV show ID, got nil")
	}
}

// TestTMDbError tests TMDb API error handling
func TestTMDbError(t *testing.T) {
	// Mock error response
	errorResponses := map[string]interface{}{
		"/movie/999999": models.ErrorResponse{
			StatusCode:    404,
			StatusMessage: "The resource you requested could not be found.",
			Success:       false,
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		endpoint := r.URL.Path
		if response, exists := errorResponses[endpoint]; exists {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := createTestClient(server.URL)
	ctx := context.Background()

	// Test error handling
	_, err := client.GetMovieDetails(ctx, 999999)
	if err == nil {
		t.Error("Expected error for non-existent movie, got nil")
	}

	// Check if error contains the expected message
	if err.Error() == "" {
		t.Error("Expected non-empty error message")
	}
}

// TestUnauthorizedAccess tests unauthorized API access
func TestUnauthorizedAccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Always return unauthorized for this test
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			StatusCode:    401,
			StatusMessage: "Invalid API key",
			Success:       false,
		})
	}))
	defer server.Close()

	// Create client with invalid API key
	cfg := &config.Config{
		TMDb: config.TMDbConfig{
			APIKey:  "invalid-key",
			BaseURL: server.URL,
		},
	}
	client := NewTMDbClient(cfg)
	ctx := context.Background()

	// Test unauthorized access
	_, err := client.SearchMovies(ctx, "test", 1)
	if err == nil {
		t.Error("Expected error for unauthorized access, got nil")
	}

	// Check if error contains the expected message
	if err.Error() == "" {
		t.Error("Expected non-empty error message")
	}
}

// TestMakeRequestWithContext tests context cancellation
func TestMakeRequestWithContext(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate slow response
		select {
		case <-r.Context().Done():
			return
		}
	}))
	defer server.Close()

	client := createTestClient(server.URL)

	// Create context with immediate cancellation
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	// Test context cancellation
	_, err := client.SearchMovies(ctx, "test", 1)
	if err == nil {
		t.Error("Expected error for cancelled context, got nil")
	}
}

// TestGetPopularMovies tests popular movies retrieval
func TestGetPopularMovies(t *testing.T) {
	mockResponse := models.PopularMovies{
		Page: 1,
		Results: []models.Movie{
			{
				ID:               123,
				Title:            "Popular Movie",
				OriginalTitle:    "Popular Movie",
				OriginalLanguage: "en",
				VoteAverage:      7.5,
			},
		},
		TotalPages:   1,
		TotalResults: 1,
	}

	responses := map[string]interface{}{
		"/movie/popular": mockResponse,
	}

	server := createMockServer(t, responses)
	defer server.Close()

	client := createTestClient(server.URL)
	ctx := context.Background()

	result, err := client.GetPopularMovies(ctx, 1)
	if err != nil {
		t.Fatalf("GetPopularMovies failed: %v", err)
	}

	if len(result.Results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(result.Results))
	}
}

// TestGetTrendingMovies tests trending movies retrieval
func TestGetTrendingMovies(t *testing.T) {
	mockResponse := models.SearchResponse[models.Movie]{
		Page: 1,
		Results: []models.Movie{
			{
				ID:               456,
				Title:            "Trending Movie",
				OriginalTitle:    "Trending Movie",
				OriginalLanguage: "en",
				VoteAverage:      8.0,
			},
		},
		TotalPages:   1,
		TotalResults: 1,
	}

	responses := map[string]interface{}{
		"/trending/movie/week": mockResponse,
		"/trending/movie/day":  mockResponse,
	}

	server := createMockServer(t, responses)
	defer server.Close()

	client := createTestClient(server.URL)
	ctx := context.Background()

	// Test week trending
	result, err := client.GetTrendingMovies(ctx, "week", 1)
	if err != nil {
		t.Fatalf("GetTrendingMovies failed: %v", err)
	}

	if len(result.Results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(result.Results))
	}

	// Test invalid time window (should default to week)
	result, err = client.GetTrendingMovies(ctx, "invalid", 1)
	if err != nil {
		t.Fatalf("GetTrendingMovies with invalid time window failed: %v", err)
	}
}

// Helper functions for pointer types
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}