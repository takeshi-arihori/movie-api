// Package services provides TMDb API client implementation for movie and TV show data retrieval.
package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/takeshi-arihori/movie-api/internal/config"
	"github.com/takeshi-arihori/movie-api/internal/models"
)

// TMDbClient represents a client for The Movie Database API
type TMDbClient struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// TMDbError represents an error response from TMDb API
type TMDbError struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}

// Error implements the error interface for TMDbError
func (e *TMDbError) Error() string {
	return fmt.Sprintf("TMDb API error %d: %s", e.StatusCode, e.StatusMessage)
}

// NewTMDbClient creates a new TMDb API client
func NewTMDbClient(cfg *config.Config) *TMDbClient {
	return &TMDbClient{
		apiKey:  cfg.TMDb.APIKey,
		baseURL: cfg.TMDb.BaseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
			},
		},
	}
}

// makeRequest performs an HTTP request to TMDb API with proper authentication and error handling
func (c *TMDbClient) makeRequest(ctx context.Context, endpoint string, params url.Values) (*http.Response, error) {
	// Build URL
	u, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid endpoint URL: %w", err)
	}

	// Add API key to parameters
	if params == nil {
		params = url.Values{}
	}
	params.Set("api_key", c.apiKey)
	u.RawQuery = params.Encode()

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Movie-API-Client/1.0")

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}

// handleResponse processes HTTP response and handles TMDb API errors
func (c *TMDbClient) handleResponse(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode >= 400 {
		var tmdbErr TMDbError
		if err := json.Unmarshal(body, &tmdbErr); err != nil {
			return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
		}
		return &tmdbErr
	}

	// Parse successful response
	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("failed to parse response JSON: %w", err)
	}

	return nil
}

// SearchMovies searches for movies by query string
func (c *TMDbClient) SearchMovies(ctx context.Context, query string, page int) (*models.MovieSearchResponse, error) {
	if query == "" {
		return nil, fmt.Errorf("search query cannot be empty")
	}

	params := url.Values{
		"query": {query},
	}

	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}

	resp, err := c.makeRequest(ctx, "/search/movie", params)
	if err != nil {
		return nil, fmt.Errorf("search movies request failed: %w", err)
	}

	var result models.MovieSearchResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("search movies response handling failed: %w", err)
	}

	return &result, nil
}

// SearchTVShows searches for TV shows by query string
func (c *TMDbClient) SearchTVShows(ctx context.Context, query string, page int) (*models.TVSearchResponse, error) {
	if query == "" {
		return nil, fmt.Errorf("search query cannot be empty")
	}

	params := url.Values{
		"query": {query},
	}

	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}

	resp, err := c.makeRequest(ctx, "/search/tv", params)
	if err != nil {
		return nil, fmt.Errorf("search TV shows request failed: %w", err)
	}

	var result models.TVSearchResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("search TV shows response handling failed: %w", err)
	}

	return &result, nil
}

// GetMovieDetails retrieves detailed information for a specific movie
func (c *TMDbClient) GetMovieDetails(ctx context.Context, movieID int) (*models.MovieDetails, error) {
	if movieID <= 0 {
		return nil, fmt.Errorf("invalid movie ID: %d", movieID)
	}

	endpoint := fmt.Sprintf("/movie/%d", movieID)
	resp, err := c.makeRequest(ctx, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("get movie details request failed: %w", err)
	}

	var result models.MovieDetails
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("get movie details response handling failed: %w", err)
	}

	return &result, nil
}

// GetTVShowDetails retrieves detailed information for a specific TV show
func (c *TMDbClient) GetTVShowDetails(ctx context.Context, tvID int) (*models.TVShowDetails, error) {
	if tvID <= 0 {
		return nil, fmt.Errorf("invalid TV show ID: %d", tvID)
	}

	endpoint := fmt.Sprintf("/tv/%d", tvID)
	resp, err := c.makeRequest(ctx, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("get TV show details request failed: %w", err)
	}

	var result models.TVShowDetails
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("get TV show details response handling failed: %w", err)
	}

	return &result, nil
}

// GetMovieCredits retrieves cast and crew information for a specific movie
func (c *TMDbClient) GetMovieCredits(ctx context.Context, movieID int) (*models.MovieCredits, error) {
	if movieID <= 0 {
		return nil, fmt.Errorf("invalid movie ID: %d", movieID)
	}

	endpoint := fmt.Sprintf("/movie/%d/credits", movieID)
	resp, err := c.makeRequest(ctx, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("get movie credits request failed: %w", err)
	}

	var result models.MovieCredits
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("get movie credits response handling failed: %w", err)
	}

	return &result, nil
}

// GetTVShowCredits retrieves cast and crew information for a specific TV show
func (c *TMDbClient) GetTVShowCredits(ctx context.Context, tvID int) (*models.TVCredits, error) {
	if tvID <= 0 {
		return nil, fmt.Errorf("invalid TV show ID: %d", tvID)
	}

	endpoint := fmt.Sprintf("/tv/%d/credits", tvID)
	resp, err := c.makeRequest(ctx, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("get TV show credits request failed: %w", err)
	}

	var result models.TVCredits
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("get TV show credits response handling failed: %w", err)
	}

	return &result, nil
}

// GetPersonDetails retrieves detailed information for a specific person
func (c *TMDbClient) GetPersonDetails(ctx context.Context, personID int) (*models.PersonDetails, error) {
	if personID <= 0 {
		return nil, fmt.Errorf("invalid person ID: %d", personID)
	}

	endpoint := fmt.Sprintf("/person/%d", personID)
	resp, err := c.makeRequest(ctx, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("get person details request failed: %w", err)
	}

	var result models.PersonDetails
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("get person details response handling failed: %w", err)
	}

	return &result, nil
}

// GetMovieReviews retrieves reviews for a specific movie
func (c *TMDbClient) GetMovieReviews(ctx context.Context, movieID int, page int) (*models.MovieReviews, error) {
	if movieID <= 0 {
		return nil, fmt.Errorf("invalid movie ID: %d", movieID)
	}

	params := url.Values{}
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}

	endpoint := fmt.Sprintf("/movie/%d/reviews", movieID)
	resp, err := c.makeRequest(ctx, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("get movie reviews request failed: %w", err)
	}

	var result models.MovieReviews
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("get movie reviews response handling failed: %w", err)
	}

	return &result, nil
}

// GetTVShowReviews retrieves reviews for a specific TV show
func (c *TMDbClient) GetTVShowReviews(ctx context.Context, tvID int, page int) (*models.TVReviews, error) {
	if tvID <= 0 {
		return nil, fmt.Errorf("invalid TV show ID: %d", tvID)
	}

	params := url.Values{}
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}

	endpoint := fmt.Sprintf("/tv/%d/reviews", tvID)
	resp, err := c.makeRequest(ctx, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("get TV show reviews request failed: %w", err)
	}

	var result models.TVReviews
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("get TV show reviews response handling failed: %w", err)
	}

	return &result, nil
}

// GetPersonMovieCredits retrieves movie credits for a specific person
func (c *TMDbClient) GetPersonMovieCredits(ctx context.Context, personID int) (*models.PersonMovieCredits, error) {
	if personID <= 0 {
		return nil, fmt.Errorf("invalid person ID: %d", personID)
	}

	endpoint := fmt.Sprintf("/person/%d/movie_credits", personID)
	resp, err := c.makeRequest(ctx, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("get person movie credits request failed: %w", err)
	}

	var result models.PersonMovieCredits
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("get person movie credits response handling failed: %w", err)
	}

	return &result, nil
}

// GetPersonTVCredits retrieves TV credits for a specific person
func (c *TMDbClient) GetPersonTVCredits(ctx context.Context, personID int) (*models.PersonTVCredits, error) {
	if personID <= 0 {
		return nil, fmt.Errorf("invalid person ID: %d", personID)
	}

	endpoint := fmt.Sprintf("/person/%d/tv_credits", personID)
	resp, err := c.makeRequest(ctx, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("get person TV credits request failed: %w", err)
	}

	var result models.PersonTVCredits
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("get person TV credits response handling failed: %w", err)
	}

	return &result, nil
}

// GetPersonCombinedCredits retrieves combined movie and TV credits for a specific person
func (c *TMDbClient) GetPersonCombinedCredits(ctx context.Context, personID int) (*models.PersonCombinedCredits, error) {
	if personID <= 0 {
		return nil, fmt.Errorf("invalid person ID: %d", personID)
	}

	endpoint := fmt.Sprintf("/person/%d/combined_credits", personID)
	resp, err := c.makeRequest(ctx, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("get person combined credits request failed: %w", err)
	}

	var result models.PersonCombinedCredits
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("get person combined credits response handling failed: %w", err)
	}

	return &result, nil
}

// GetPopularMovies retrieves popular movies
func (c *TMDbClient) GetPopularMovies(ctx context.Context, page int) (*models.PopularMovies, error) {
	params := url.Values{}
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}

	resp, err := c.makeRequest(ctx, "/movie/popular", params)
	if err != nil {
		return nil, fmt.Errorf("get popular movies request failed: %w", err)
	}

	var result models.PopularMovies
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("get popular movies response handling failed: %w", err)
	}

	return &result, nil
}

// GetTopRatedMovies retrieves top rated movies
func (c *TMDbClient) GetTopRatedMovies(ctx context.Context, page int) (*models.TopRatedMovies, error) {
	params := url.Values{}
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}

	resp, err := c.makeRequest(ctx, "/movie/top_rated", params)
	if err != nil {
		return nil, fmt.Errorf("get top rated movies request failed: %w", err)
	}

	var result models.TopRatedMovies
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("get top rated movies response handling failed: %w", err)
	}

	return &result, nil
}

// GetTrendingMovies retrieves trending movies
func (c *TMDbClient) GetTrendingMovies(ctx context.Context, timeWindow string, page int) (*models.SearchResponse[models.Movie], error) {
	if timeWindow != "day" && timeWindow != "week" {
		timeWindow = "week" // Default to week
	}

	params := url.Values{}
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}

	endpoint := fmt.Sprintf("/trending/movie/%s", timeWindow)
	resp, err := c.makeRequest(ctx, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("get trending movies request failed: %w", err)
	}

	var result models.SearchResponse[models.Movie]
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("get trending movies response handling failed: %w", err)
	}

	return &result, nil
}

// GetTrendingTVShows retrieves trending TV shows
func (c *TMDbClient) GetTrendingTVShows(ctx context.Context, timeWindow string, page int) (*models.SearchResponse[models.TVShow], error) {
	if timeWindow != "day" && timeWindow != "week" {
		timeWindow = "week" // Default to week
	}

	params := url.Values{}
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}

	endpoint := fmt.Sprintf("/trending/tv/%s", timeWindow)
	resp, err := c.makeRequest(ctx, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("get trending TV shows request failed: %w", err)
	}

	var result models.SearchResponse[models.TVShow]
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("get trending TV shows response handling failed: %w", err)
	}

	return &result, nil
}

// MultiSearch performs a multi-search across movies, TV shows, and people
func (c *TMDbClient) MultiSearch(ctx context.Context, query string, page int, language string) (*models.MultiSearchResponse, error) {
	if query == "" {
		return nil, fmt.Errorf("search query cannot be empty")
	}

	params := url.Values{
		"query": {query},
	}

	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}

	if language != "" {
		params.Set("language", language)
	}

	resp, err := c.makeRequest(ctx, "/search/multi", params)
	if err != nil {
		return nil, fmt.Errorf("multi search request failed: %w", err)
	}

	var result models.MultiSearchResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("multi search response handling failed: %w", err)
	}

	return &result, nil
}

// SearchByType performs a search filtered by media type (movie, tv, or person)
func (c *TMDbClient) SearchByType(ctx context.Context, searchType, query string, page int, language string) (*models.MultiSearchResponse, error) {
	if query == "" {
		return nil, fmt.Errorf("search query cannot be empty")
	}

	var endpoint string
	switch searchType {
	case "movie":
		endpoint = "/search/movie"
	case "tv":
		endpoint = "/search/tv"
	case "person":
		endpoint = "/search/person"
	default:
		return c.MultiSearch(ctx, query, page, language)
	}

	params := url.Values{
		"query": {query},
	}

	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}

	if language != "" {
		params.Set("language", language)
	}

	resp, err := c.makeRequest(ctx, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("search by type request failed: %w", err)
	}

	// For type-specific searches, we need to convert the response to MultiSearchResponse
	switch searchType {
	case "movie":
		var movieResult models.MovieSearchResponse
		if err := c.handleResponse(resp, &movieResult); err != nil {
			return nil, fmt.Errorf("movie search response handling failed: %w", err)
		}
		return c.convertMovieSearchToMultiSearch(&movieResult), nil

	case "tv":
		var tvResult models.TVSearchResponse
		if err := c.handleResponse(resp, &tvResult); err != nil {
			return nil, fmt.Errorf("TV search response handling failed: %w", err)
		}
		return c.convertTVSearchToMultiSearch(&tvResult), nil

	case "person":
		var personResult models.PersonSearchResponse
		if err := c.handleResponse(resp, &personResult); err != nil {
			return nil, fmt.Errorf("person search response handling failed: %w", err)
		}
		return c.convertPersonSearchToMultiSearch(&personResult), nil

	default:
		return nil, fmt.Errorf("unsupported search type: %s", searchType)
	}
}

// convertMovieSearchToMultiSearch converts MovieSearchResponse to MultiSearchResponse
func (c *TMDbClient) convertMovieSearchToMultiSearch(movieResult *models.MovieSearchResponse) *models.MultiSearchResponse {
	results := make([]models.MultiSearchResult, len(movieResult.Results))
	for i, movie := range movieResult.Results {
		results[i] = models.MultiSearchResult{
			ID:               movie.ID,
			MediaType:        models.SearchItemTypeMovie,
			Popularity:       movie.Popularity,
			Adult:            &movie.Adult,
			BackdropPath:     movie.BackdropPath,
			GenreIDs:         movie.GenreIDs,
			OriginalLanguage: &movie.OriginalLanguage,
			OriginalTitle:    &movie.OriginalTitle,
			Overview:         &movie.Overview,
			PosterPath:       movie.PosterPath,
			ReleaseDate:      movie.ReleaseDate,
			Title:            &movie.Title,
			Video:            &movie.Video,
			VoteAverage:      &movie.VoteAverage,
			VoteCount:        &movie.VoteCount,
		}
	}

	return &models.MultiSearchResponse{
		Page:         movieResult.Page,
		Results:      results,
		TotalPages:   movieResult.TotalPages,
		TotalResults: movieResult.TotalResults,
	}
}

// convertTVSearchToMultiSearch converts TVSearchResponse to MultiSearchResponse
func (c *TMDbClient) convertTVSearchToMultiSearch(tvResult *models.TVSearchResponse) *models.MultiSearchResponse {
	results := make([]models.MultiSearchResult, len(tvResult.Results))
	for i, tv := range tvResult.Results {
		results[i] = models.MultiSearchResult{
			ID:               tv.ID,
			MediaType:        models.SearchItemTypeTV,
			Popularity:       tv.Popularity,
			Adult:            &tv.Adult,
			BackdropPath:     tv.BackdropPath,
			GenreIDs:         tv.GenreIDs,
			OriginalLanguage: &tv.OriginalLanguage,
			OriginalName:     &tv.OriginalName,
			Overview:         &tv.Overview,
			PosterPath:       tv.PosterPath,
			FirstAirDate:     tv.FirstAirDate,
			Name:             &tv.Name,
			OriginCountry:    tv.OriginCountry,
			VoteAverage:      &tv.VoteAverage,
			VoteCount:        &tv.VoteCount,
		}
	}

	return &models.MultiSearchResponse{
		Page:         tvResult.Page,
		Results:      results,
		TotalPages:   tvResult.TotalPages,
		TotalResults: tvResult.TotalResults,
	}
}

// convertPersonSearchToMultiSearch converts PersonSearchResponse to MultiSearchResponse
func (c *TMDbClient) convertPersonSearchToMultiSearch(personResult *models.PersonSearchResponse) *models.MultiSearchResponse {
	results := make([]models.MultiSearchResult, len(personResult.Results))
	for i, person := range personResult.Results {
		results[i] = models.MultiSearchResult{
			ID:                 person.ID,
			MediaType:          models.SearchItemTypePerson,
			Popularity:         person.Popularity,
			Adult:              &person.Adult,
			Gender:             &person.Gender,
			KnownFor:           person.KnownFor,
			KnownForDepartment: &person.KnownForDepartment,
			Name:               &person.Name,
			OriginalName:       &person.OriginalName,
			ProfilePath:        person.ProfilePath,
		}
	}

	return &models.MultiSearchResponse{
		Page:         personResult.Page,
		Results:      results,
		TotalPages:   personResult.TotalPages,
		TotalResults: personResult.TotalResults,
	}
}