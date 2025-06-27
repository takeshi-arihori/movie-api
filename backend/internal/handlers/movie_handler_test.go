package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/takeshi-arihori/movie-api/internal/models"
	"github.com/takeshi-arihori/movie-api/internal/services"
)

// MockTMDbClient is a mock implementation of TMDbClient for testing
type MockTMDbClient struct {
	movieDetails *models.MovieDetails
	movieCredits *models.MovieCredits
	movieReviews *models.MovieReviews
	err          error
}

func (m *MockTMDbClient) GetMovieDetails(ctx context.Context, movieID int) (*models.MovieDetails, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.movieDetails, nil
}

func (m *MockTMDbClient) GetMovieCredits(ctx context.Context, movieID int) (*models.MovieCredits, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.movieCredits, nil
}

func (m *MockTMDbClient) GetMovieReviews(ctx context.Context, movieID int, page int) (*models.MovieReviews, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.movieReviews, nil
}

func TestMovieHandler_GetMovieDetails(t *testing.T) {
	tests := []struct {
		name           string
		movieID        string
		mockResponse   *models.MovieDetails
		mockError      error
		expectedStatus int
		expectedError  string
	}{
		{
			name:    "successful movie details retrieval",
			movieID: "123",
			mockResponse: &models.MovieDetails{
				ID:               123,
				Title:            "Test Movie",
				OriginalTitle:    "Test Movie Original",
				OriginalLanguage: "en",
				Overview:         stringPtr("Test movie overview"),
				Popularity:       7.5,
				VoteAverage:      8.2,
				VoteCount:        1000,
				Runtime:          intPtr(120),
				Budget:           1000000,
				Revenue:          5000000,
				Status:           "Released",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid movie ID - zero",
			movieID:        "0",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid_parameter",
		},
		{
			name:           "invalid movie ID - non-numeric",
			movieID:        "abc",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid_parameter",
		},
		{
			name:           "movie not found",
			movieID:        "999",
			mockError:      &services.TMDbError{StatusCode: 404, StatusMessage: "Not found"},
			expectedStatus: http.StatusNotFound,
			expectedError:  "movie_not_found",
		},
		{
			name:           "TMDb API error",
			movieID:        "123",
			mockError:      &services.TMDbError{StatusCode: 500, StatusMessage: "Internal server error"},
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "api_error",
		},
		{
			name:           "generic error",
			movieID:        "123",
			mockError:      fmt.Errorf("network error"),
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "api_error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock client
			mockClient := &MockTMDbClient{
				movieDetails: tt.mockResponse,
				err:          tt.mockError,
			}
			handler := NewMovieHandler(mockClient)

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/api/v1/movies/"+tt.movieID, nil)
			w := httptest.NewRecorder()

			// Setup router to extract path parameters
			router := mux.NewRouter()
			router.HandleFunc("/api/v1/movies/{id}", handler.GetMovieDetails).Methods("GET")
			router.ServeHTTP(w, req)

			// Check status code
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			// Check response body
			if tt.expectedError != "" {
				var errorResp ErrorResponse
				if err := json.NewDecoder(w.Body).Decode(&errorResp); err != nil {
					t.Fatalf("failed to decode error response: %v", err)
				}
				if errorResp.Error != tt.expectedError {
					t.Errorf("expected error %q, got %q", tt.expectedError, errorResp.Error)
				}
			} else if tt.mockResponse != nil {
				var movieDetails models.MovieDetails
				if err := json.NewDecoder(w.Body).Decode(&movieDetails); err != nil {
					t.Fatalf("failed to decode movie details response: %v", err)
				}
				if movieDetails.ID != tt.mockResponse.ID {
					t.Errorf("expected movie ID %d, got %d", tt.mockResponse.ID, movieDetails.ID)
				}
				if movieDetails.Title != tt.mockResponse.Title {
					t.Errorf("expected title %q, got %q", tt.mockResponse.Title, movieDetails.Title)
				}
			}
		})
	}
}

func TestMovieHandler_GetMovieCredits(t *testing.T) {
	tests := []struct {
		name           string
		movieID        string
		mockResponse   *models.MovieCredits
		mockError      error
		expectedStatus int
		expectedError  string
	}{
		{
			name:    "successful movie credits retrieval",
			movieID: "123",
			mockResponse: &models.MovieCredits{
				ID: 123,
				Cast: []models.CastMember{
					{
						ID:        1,
						Name:      "Actor 1",
						Character: "Character 1",
						CreditID:  "credit1",
					},
				},
				Crew: []models.CrewMember{
					{
						ID:         2,
						Name:       "Director 1",
						Job:        "Director",
						Department: "Directing",
						CreditID:   "credit2",
					},
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid movie ID - zero",
			movieID:        "0",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid_parameter",
		},
		{
			name:           "movie not found",
			movieID:        "999",
			mockError:      &services.TMDbError{StatusCode: 404, StatusMessage: "Not found"},
			expectedStatus: http.StatusNotFound,
			expectedError:  "movie_not_found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock client
			mockClient := &MockTMDbClient{
				movieCredits: tt.mockResponse,
				err:          tt.mockError,
			}
			handler := NewMovieHandler(mockClient)

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/api/v1/movies/"+tt.movieID+"/credits", nil)
			w := httptest.NewRecorder()

			// Setup router to extract path parameters
			router := mux.NewRouter()
			router.HandleFunc("/api/v1/movies/{id}/credits", handler.GetMovieCredits).Methods("GET")
			router.ServeHTTP(w, req)

			// Check status code
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			// Check response body
			if tt.expectedError != "" {
				var errorResp ErrorResponse
				if err := json.NewDecoder(w.Body).Decode(&errorResp); err != nil {
					t.Fatalf("failed to decode error response: %v", err)
				}
				if errorResp.Error != tt.expectedError {
					t.Errorf("expected error %q, got %q", tt.expectedError, errorResp.Error)
				}
			} else if tt.mockResponse != nil {
				var movieCredits models.MovieCredits
				if err := json.NewDecoder(w.Body).Decode(&movieCredits); err != nil {
					t.Fatalf("failed to decode movie credits response: %v", err)
				}
				if movieCredits.ID != tt.mockResponse.ID {
					t.Errorf("expected movie ID %d, got %d", tt.mockResponse.ID, movieCredits.ID)
				}
				if len(movieCredits.Cast) != len(tt.mockResponse.Cast) {
					t.Errorf("expected %d cast members, got %d", len(tt.mockResponse.Cast), len(movieCredits.Cast))
				}
			}
		})
	}
}

func TestMovieHandler_GetMovieReviews(t *testing.T) {
	tests := []struct {
		name           string
		movieID        string
		mockResponse   *models.MovieReviews
		mockError      error
		expectedStatus int
		expectedError  string
	}{
		{
			name:    "successful movie reviews retrieval",
			movieID: "123",
			mockResponse: &models.MovieReviews{
				ID:           123,
				Page:         1,
				TotalPages:   1,
				TotalResults: 1,
				Results: []models.Review{
					{
						ID:      "review1",
						Author:  "Reviewer 1",
						Content: "Great movie!",
						URL:     "http://example.com/review1",
					},
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid movie ID - zero",
			movieID:        "0",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid_parameter",
		},
		{
			name:           "movie not found",
			movieID:        "999",
			mockError:      &services.TMDbError{StatusCode: 404, StatusMessage: "Not found"},
			expectedStatus: http.StatusNotFound,
			expectedError:  "movie_not_found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock client
			mockClient := &MockTMDbClient{
				movieReviews: tt.mockResponse,
				err:          tt.mockError,
			}
			handler := NewMovieHandler(mockClient)

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/api/v1/movies/"+tt.movieID+"/reviews", nil)
			w := httptest.NewRecorder()

			// Setup router to extract path parameters
			router := mux.NewRouter()
			router.HandleFunc("/api/v1/movies/{id}/reviews", handler.GetMovieReviews).Methods("GET")
			router.ServeHTTP(w, req)

			// Check status code
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			// Check response body
			if tt.expectedError != "" {
				var errorResp ErrorResponse
				if err := json.NewDecoder(w.Body).Decode(&errorResp); err != nil {
					t.Fatalf("failed to decode error response: %v", err)
				}
				if errorResp.Error != tt.expectedError {
					t.Errorf("expected error %q, got %q", tt.expectedError, errorResp.Error)
				}
			} else if tt.mockResponse != nil {
				var movieReviews models.MovieReviews
				if err := json.NewDecoder(w.Body).Decode(&movieReviews); err != nil {
					t.Fatalf("failed to decode movie reviews response: %v", err)
				}
				if movieReviews.ID != tt.mockResponse.ID {
					t.Errorf("expected movie ID %d, got %d", tt.mockResponse.ID, movieReviews.ID)
				}
				if len(movieReviews.Results) != len(tt.mockResponse.Results) {
					t.Errorf("expected %d reviews, got %d", len(tt.mockResponse.Results), len(movieReviews.Results))
				}
			}
		})
	}
}

func TestMovieHandler_MethodNotAllowed(t *testing.T) {
	mockClient := &MockTMDbClient{}
	handler := NewMovieHandler(mockClient)

	// Test POST method (not allowed)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/movies/123", nil)
	w := httptest.NewRecorder()

	handler.GetMovieDetails(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status %d, got %d", http.StatusMethodNotAllowed, w.Code)
	}

	var errorResp ErrorResponse
	if err := json.NewDecoder(w.Body).Decode(&errorResp); err != nil {
		t.Fatalf("failed to decode error response: %v", err)
	}
	if errorResp.Error != "method_not_allowed" {
		t.Errorf("expected error 'method_not_allowed', got %q", errorResp.Error)
	}
}

func TestMovieHandler_CORS(t *testing.T) {
	mockClient := &MockTMDbClient{}
	handler := NewMovieHandler(mockClient)

	// Test OPTIONS request
	req := httptest.NewRequest(http.MethodOptions, "/api/v1/movies/123", nil)
	w := httptest.NewRecorder()

	handler.GetMovieDetails(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Check CORS headers
	expectedHeaders := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET, OPTIONS",
		"Access-Control-Allow-Headers": "Content-Type",
	}

	for header, expectedValue := range expectedHeaders {
		actualValue := w.Header().Get(header)
		if actualValue != expectedValue {
			t.Errorf("expected header %s: %q, got %q", header, expectedValue, actualValue)
		}
	}
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}