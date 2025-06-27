package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/takeshi-arihori/movie-api/internal/models"
	"github.com/takeshi-arihori/movie-api/internal/services"
)

// MockReviewClient is a mock implementation of ReviewClient for testing
type MockReviewClient struct {
	movieReviews *models.MovieReviews
	tvReviews    *models.TVReviews
	err          error
}

func (m *MockReviewClient) GetMovieReviews(ctx context.Context, movieID int, page int) (*models.MovieReviews, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.movieReviews, nil
}

func (m *MockReviewClient) GetTVShowReviews(ctx context.Context, tvID int, page int) (*models.TVReviews, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.tvReviews, nil
}

func TestReviewHandler_GetMovieReviews(t *testing.T) {
	tests := []struct {
		name           string
		movieID        string
		queryParams    string
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
				TotalResults: 2,
				Results: []models.Review{
					{
						ID:      "review1",
						Author:  "John Doe",
						Content: "Great movie!",
						CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
						URL:     "http://example.com/review1",
						AuthorDetails: models.AuthorDetails{
							Name:     "John Doe",
							Username: "johndoe",
							Rating:   floatPtr(8.5),
						},
					},
					{
						ID:      "review2",
						Author:  "Jane Smith",
						Content: "Amazing cinematography!",
						CreatedAt: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
						URL:     "http://example.com/review2",
						AuthorDetails: models.AuthorDetails{
							Name:     "Jane Smith",
							Username: "janesmith",
							Rating:   floatPtr(9.0),
						},
					},
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:        "movie reviews with page parameter",
			movieID:     "123",
			queryParams: "?page=2",
			mockResponse: &models.MovieReviews{
				ID:           123,
				Page:         2,
				TotalPages:   3,
				TotalResults: 25,
				Results: []models.Review{
					{
						ID:      "review3",
						Author:  "Bob Wilson",
						Content: "Not bad, but could be better.",
						CreatedAt: time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC),
						URL:     "http://example.com/review3",
						AuthorDetails: models.AuthorDetails{
							Name:     "Bob Wilson",
							Username: "bobwilson",
							Rating:   floatPtr(6.0),
						},
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
			mockClient := &MockReviewClient{
				movieReviews: tt.mockResponse,
				err:          tt.mockError,
			}
			handler := NewReviewHandler(mockClient)

			// Create request
			url := "/api/v1/movies/" + tt.movieID + "/reviews" + tt.queryParams
			req := httptest.NewRequest(http.MethodGet, url, nil)
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
				if movieReviews.Page != tt.mockResponse.Page {
					t.Errorf("expected page %d, got %d", tt.mockResponse.Page, movieReviews.Page)
				}
				if len(movieReviews.Results) != len(tt.mockResponse.Results) {
					t.Errorf("expected %d reviews, got %d", len(tt.mockResponse.Results), len(movieReviews.Results))
				}
			}
		})
	}
}

func TestReviewHandler_GetTVReviews(t *testing.T) {
	tests := []struct {
		name           string
		tvID           string
		queryParams    string
		mockResponse   *models.TVReviews
		mockError      error
		expectedStatus int
		expectedError  string
	}{
		{
			name: "successful TV reviews retrieval",
			tvID: "456",
			mockResponse: &models.TVReviews{
				ID:           456,
				Page:         1,
				TotalPages:   1,
				TotalResults: 1,
				Results: []models.Review{
					{
						ID:      "tv_review1",
						Author:  "TV Critic",
						Content: "Excellent series!",
						CreatedAt: time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
						URL:     "http://example.com/tv_review1",
						AuthorDetails: models.AuthorDetails{
							Name:     "TV Critic",
							Username: "tvcritic",
							Rating:   floatPtr(9.5),
						},
					},
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:        "TV reviews with page parameter",
			tvID:        "456",
			queryParams: "?page=3",
			mockResponse: &models.TVReviews{
				ID:           456,
				Page:         3,
				TotalPages:   5,
				TotalResults: 48,
				Results:      []models.Review{},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid TV ID - zero",
			tvID:           "0",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid_parameter",
		},
		{
			name:           "invalid TV ID - non-numeric",
			tvID:           "xyz",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid_parameter",
		},
		{
			name:           "TV show not found",
			tvID:           "999",
			mockError:      &services.TMDbError{StatusCode: 404, StatusMessage: "Not found"},
			expectedStatus: http.StatusNotFound,
			expectedError:  "tv_not_found",
		},
		{
			name:           "TMDb API error",
			tvID:           "456",
			mockError:      &services.TMDbError{StatusCode: 500, StatusMessage: "Internal server error"},
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "api_error",
		},
		{
			name:           "generic error",
			tvID:           "456",
			mockError:      fmt.Errorf("network timeout"),
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "api_error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock client
			mockClient := &MockReviewClient{
				tvReviews: tt.mockResponse,
				err:       tt.mockError,
			}
			handler := NewReviewHandler(mockClient)

			// Create request
			url := "/api/v1/tv/" + tt.tvID + "/reviews" + tt.queryParams
			req := httptest.NewRequest(http.MethodGet, url, nil)
			w := httptest.NewRecorder()

			// Setup router to extract path parameters
			router := mux.NewRouter()
			router.HandleFunc("/api/v1/tv/{id}/reviews", handler.GetTVReviews).Methods("GET")
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
				var tvReviews models.TVReviews
				if err := json.NewDecoder(w.Body).Decode(&tvReviews); err != nil {
					t.Fatalf("failed to decode TV reviews response: %v", err)
				}
				if tvReviews.ID != tt.mockResponse.ID {
					t.Errorf("expected TV ID %d, got %d", tt.mockResponse.ID, tvReviews.ID)
				}
				if tvReviews.Page != tt.mockResponse.Page {
					t.Errorf("expected page %d, got %d", tt.mockResponse.Page, tvReviews.Page)
				}
				if len(tvReviews.Results) != len(tt.mockResponse.Results) {
					t.Errorf("expected %d reviews, got %d", len(tt.mockResponse.Results), len(tvReviews.Results))
				}
			}
		})
	}
}

func TestReviewHandler_MethodNotAllowed(t *testing.T) {
	mockClient := &MockReviewClient{}
	handler := NewReviewHandler(mockClient)

	// Test POST method (not allowed) for movie reviews
	req := httptest.NewRequest(http.MethodPost, "/api/v1/movies/123/reviews", nil)
	w := httptest.NewRecorder()

	handler.GetMovieReviews(w, req)

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

func TestReviewHandler_CORS(t *testing.T) {
	mockClient := &MockReviewClient{}
	handler := NewReviewHandler(mockClient)

	// Test OPTIONS request for movie reviews
	req := httptest.NewRequest(http.MethodOptions, "/api/v1/movies/123/reviews", nil)
	w := httptest.NewRecorder()

	handler.GetMovieReviews(w, req)

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

	// Test OPTIONS request for TV reviews
	req = httptest.NewRequest(http.MethodOptions, "/api/v1/tv/456/reviews", nil)
	w = httptest.NewRecorder()

	handler.GetTVReviews(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	for header, expectedValue := range expectedHeaders {
		actualValue := w.Header().Get(header)
		if actualValue != expectedValue {
			t.Errorf("expected header %s: %q, got %q", header, expectedValue, actualValue)
		}
	}
}

// Helper function
func floatPtr(f float64) *float64 {
	return &f
}