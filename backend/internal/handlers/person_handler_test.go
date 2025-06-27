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

// MockPersonClient is a mock implementation of PersonClient for testing
type MockPersonClient struct {
	personDetails      *models.PersonDetails
	personMovieCredits *models.PersonMovieCredits
	personTVCredits    *models.PersonTVCredits
	combinedCredits    *models.PersonCombinedCredits
	err                error
}

func (m *MockPersonClient) GetPersonDetails(ctx context.Context, personID int) (*models.PersonDetails, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.personDetails, nil
}

func (m *MockPersonClient) GetPersonMovieCredits(ctx context.Context, personID int) (*models.PersonMovieCredits, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.personMovieCredits, nil
}

func (m *MockPersonClient) GetPersonTVCredits(ctx context.Context, personID int) (*models.PersonTVCredits, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.personTVCredits, nil
}

func (m *MockPersonClient) GetPersonCombinedCredits(ctx context.Context, personID int) (*models.PersonCombinedCredits, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.combinedCredits, nil
}

func TestPersonHandler_GetPersonDetails(t *testing.T) {
	tests := []struct {
		name           string
		personID       string
		mockResponse   *models.PersonDetails
		mockError      error
		expectedStatus int
		expectedError  string
	}{
		{
			name:     "successful person details retrieval",
			personID: "123",
			mockResponse: &models.PersonDetails{
				ID:                 123,
				Name:               "Test Actor",
				Biography:          "Test biography",
				Birthday:           stringPtr("1990-01-01"),
				Deathday:           nil,
				Gender:             2,
				KnownForDepartment: "Acting",
				Popularity:         85.5,
				PlaceOfBirth:       stringPtr("Test City"),
				ProfilePath:        stringPtr("/test_profile.jpg"),
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid person ID - zero",
			personID:       "0",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid_parameter",
		},
		{
			name:           "invalid person ID - non-numeric",
			personID:       "abc",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid_parameter",
		},
		{
			name:           "person not found",
			personID:       "999",
			mockError:      &services.TMDbError{StatusCode: 404, StatusMessage: "Not found"},
			expectedStatus: http.StatusNotFound,
			expectedError:  "person_not_found",
		},
		{
			name:           "TMDb API error",
			personID:       "123",
			mockError:      &services.TMDbError{StatusCode: 500, StatusMessage: "Internal server error"},
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "api_error",
		},
		{
			name:           "generic error",
			personID:       "123",
			mockError:      fmt.Errorf("network error"),
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "api_error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock client
			mockClient := &MockPersonClient{
				personDetails: tt.mockResponse,
				err:           tt.mockError,
			}
			handler := NewPersonHandler(mockClient)

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/api/v1/people/"+tt.personID, nil)
			w := httptest.NewRecorder()

			// Setup router to extract path parameters
			router := mux.NewRouter()
			router.HandleFunc("/api/v1/people/{id}", handler.GetPersonDetails).Methods("GET")
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
				var personDetails models.PersonDetails
				if err := json.NewDecoder(w.Body).Decode(&personDetails); err != nil {
					t.Fatalf("failed to decode person details response: %v", err)
				}
				if personDetails.ID != tt.mockResponse.ID {
					t.Errorf("expected person ID %d, got %d", tt.mockResponse.ID, personDetails.ID)
				}
				if personDetails.Name != tt.mockResponse.Name {
					t.Errorf("expected name %q, got %q", tt.mockResponse.Name, personDetails.Name)
				}
			}
		})
	}
}

func TestPersonHandler_GetPersonMovieCredits(t *testing.T) {
	tests := []struct {
		name           string
		personID       string
		mockResponse   *models.PersonMovieCredits
		mockError      error
		expectedStatus int
		expectedError  string
	}{
		{
			name:     "successful person movie credits retrieval",
			personID: "123",
			mockResponse: &models.PersonMovieCredits{
				ID: 123,
				Cast: []models.PersonMovieCast{
					{
						ID:        1,
						Title:     "Test Movie",
						Character: "Test Character",
						CreditID:  "credit1",
						Order:     1,
					},
				},
				Crew: []models.PersonMovieCrew{
					{
						ID:         2,
						Title:      "Test Movie 2",
						Job:        "Director",
						Department: "Directing",
						CreditID:   "credit2",
					},
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid person ID - zero",
			personID:       "0",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid_parameter",
		},
		{
			name:           "person not found",
			personID:       "999",
			mockError:      &services.TMDbError{StatusCode: 404, StatusMessage: "Not found"},
			expectedStatus: http.StatusNotFound,
			expectedError:  "person_not_found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock client
			mockClient := &MockPersonClient{
				personMovieCredits: tt.mockResponse,
				err:                tt.mockError,
			}
			handler := NewPersonHandler(mockClient)

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/api/v1/people/"+tt.personID+"/movie_credits", nil)
			w := httptest.NewRecorder()

			// Setup router to extract path parameters
			router := mux.NewRouter()
			router.HandleFunc("/api/v1/people/{id}/movie_credits", handler.GetPersonMovieCredits).Methods("GET")
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
				var movieCredits models.PersonMovieCredits
				if err := json.NewDecoder(w.Body).Decode(&movieCredits); err != nil {
					t.Fatalf("failed to decode movie credits response: %v", err)
				}
				if movieCredits.ID != tt.mockResponse.ID {
					t.Errorf("expected person ID %d, got %d", tt.mockResponse.ID, movieCredits.ID)
				}
				if len(movieCredits.Cast) != len(tt.mockResponse.Cast) {
					t.Errorf("expected %d cast members, got %d", len(tt.mockResponse.Cast), len(movieCredits.Cast))
				}
			}
		})
	}
}

func TestPersonHandler_GetPersonTVCredits(t *testing.T) {
	tests := []struct {
		name           string
		personID       string
		mockResponse   *models.PersonTVCredits
		mockError      error
		expectedStatus int
		expectedError  string
	}{
		{
			name:     "successful person TV credits retrieval",
			personID: "123",
			mockResponse: &models.PersonTVCredits{
				ID: 123,
				Cast: []models.PersonTVCast{
					{
						ID:           1,
						Name:         "Test TV Show",
						Character:    "Test Character",
						CreditID:     "credit1",
						EpisodeCount: 10,
					},
				},
				Crew: []models.PersonTVCrew{
					{
						ID:           2,
						Name:         "Test TV Show 2",
						Job:          "Producer",
						Department:   "Production",
						CreditID:     "credit2",
						EpisodeCount: 5,
					},
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid person ID - zero",
			personID:       "0",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid_parameter",
		},
		{
			name:           "person not found",
			personID:       "999",
			mockError:      &services.TMDbError{StatusCode: 404, StatusMessage: "Not found"},
			expectedStatus: http.StatusNotFound,
			expectedError:  "person_not_found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock client
			mockClient := &MockPersonClient{
				personTVCredits: tt.mockResponse,
				err:             tt.mockError,
			}
			handler := NewPersonHandler(mockClient)

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/api/v1/people/"+tt.personID+"/tv_credits", nil)
			w := httptest.NewRecorder()

			// Setup router to extract path parameters
			router := mux.NewRouter()
			router.HandleFunc("/api/v1/people/{id}/tv_credits", handler.GetPersonTVCredits).Methods("GET")
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
				var tvCredits models.PersonTVCredits
				if err := json.NewDecoder(w.Body).Decode(&tvCredits); err != nil {
					t.Fatalf("failed to decode TV credits response: %v", err)
				}
				if tvCredits.ID != tt.mockResponse.ID {
					t.Errorf("expected person ID %d, got %d", tt.mockResponse.ID, tvCredits.ID)
				}
				if len(tvCredits.Cast) != len(tt.mockResponse.Cast) {
					t.Errorf("expected %d cast members, got %d", len(tt.mockResponse.Cast), len(tvCredits.Cast))
				}
			}
		})
	}
}

func TestPersonHandler_GetPersonCombinedCredits(t *testing.T) {
	tests := []struct {
		name           string
		personID       string
		mockResponse   *models.PersonCombinedCredits
		mockError      error
		expectedStatus int
		expectedError  string
	}{
		{
			name:     "successful person combined credits retrieval",
			personID: "123",
			mockResponse: &models.PersonCombinedCredits{
				ID: 123,
				Cast: []models.PersonCombinedCast{
					{
						ID:        1,
						Title:     stringPtr("Test Movie"),
						Character: "Test Character",
						CreditID:  "credit1",
						MediaType: "movie",
						Order:     intPtr(1),
					},
				},
				Crew: []models.PersonCombinedCrew{
					{
						ID:         2,
						Title:      stringPtr("Test Movie 2"),
						Job:        "Director",
						Department: "Directing",
						CreditID:   "credit2",
						MediaType:  "movie",
					},
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid person ID - zero",
			personID:       "0",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid_parameter",
		},
		{
			name:           "person not found",
			personID:       "999",
			mockError:      &services.TMDbError{StatusCode: 404, StatusMessage: "Not found"},
			expectedStatus: http.StatusNotFound,
			expectedError:  "person_not_found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock client
			mockClient := &MockPersonClient{
				combinedCredits: tt.mockResponse,
				err:             tt.mockError,
			}
			handler := NewPersonHandler(mockClient)

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/api/v1/people/"+tt.personID+"/combined_credits", nil)
			w := httptest.NewRecorder()

			// Setup router to extract path parameters
			router := mux.NewRouter()
			router.HandleFunc("/api/v1/people/{id}/combined_credits", handler.GetPersonCombinedCredits).Methods("GET")
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
				var combinedCredits models.PersonCombinedCredits
				if err := json.NewDecoder(w.Body).Decode(&combinedCredits); err != nil {
					t.Fatalf("failed to decode combined credits response: %v", err)
				}
				if combinedCredits.ID != tt.mockResponse.ID {
					t.Errorf("expected person ID %d, got %d", tt.mockResponse.ID, combinedCredits.ID)
				}
				if len(combinedCredits.Cast) != len(tt.mockResponse.Cast) {
					t.Errorf("expected %d cast members, got %d", len(tt.mockResponse.Cast), len(combinedCredits.Cast))
				}
			}
		})
	}
}

func TestPersonHandler_MethodNotAllowed(t *testing.T) {
	mockClient := &MockPersonClient{}
	handler := NewPersonHandler(mockClient)

	// Test POST method (not allowed)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/people/123", nil)
	w := httptest.NewRecorder()

	handler.GetPersonDetails(w, req)

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

func TestPersonHandler_CORS(t *testing.T) {
	mockClient := &MockPersonClient{}
	handler := NewPersonHandler(mockClient)

	// Test OPTIONS request
	req := httptest.NewRequest(http.MethodOptions, "/api/v1/people/123", nil)
	w := httptest.NewRecorder()

	handler.GetPersonDetails(w, req)

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

// Helper functions for testing (moved to avoid duplication)
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}