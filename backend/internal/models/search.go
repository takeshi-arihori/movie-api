// Package models provides search-related data structures for multi-search functionality.
package models

// SearchItemType represents the type of search result item
type SearchItemType string

const (
	SearchItemTypeMovie  SearchItemType = "movie"
	SearchItemTypeTV     SearchItemType = "tv"
	SearchItemTypePerson SearchItemType = "person"
)

// MultiSearchResult represents a unified search result that can be a movie, TV show, or person
type MultiSearchResult struct {
	// Common fields for all types
	ID           int             `json:"id" validate:"required"`
	MediaType    SearchItemType  `json:"media_type" validate:"required"`
	Popularity   float64         `json:"popularity"`
	
	// Movie/TV show specific fields
	Adult            *bool    `json:"adult,omitempty"`
	BackdropPath     *string  `json:"backdrop_path,omitempty"`
	GenreIDs         []int    `json:"genre_ids,omitempty"`
	OriginalLanguage *string  `json:"original_language,omitempty"`
	Overview         *string  `json:"overview,omitempty"`
	PosterPath       *string  `json:"poster_path,omitempty"`
	VoteAverage      *float64 `json:"vote_average,omitempty"`
	VoteCount        *int     `json:"vote_count,omitempty"`
	
	// Movie specific fields
	OriginalTitle *string `json:"original_title,omitempty"`
	ReleaseDate   *string `json:"release_date,omitempty"`
	Title         *string `json:"title,omitempty"`
	Video         *bool   `json:"video,omitempty"`
	
	// TV show specific fields
	FirstAirDate     *string  `json:"first_air_date,omitempty"`
	Name             *string  `json:"name,omitempty"`
	OriginalName     *string  `json:"original_name,omitempty"`
	OriginCountry    []string `json:"origin_country,omitempty"`
	
	// Person specific fields
	Gender             *int     `json:"gender,omitempty"`
	KnownFor           []KnownForItem `json:"known_for,omitempty"`
	KnownForDepartment *string  `json:"known_for_department,omitempty"`
	ProfilePath        *string  `json:"profile_path,omitempty"`
}

// MultiSearchResponse represents a multi-search response from TMDb API
type MultiSearchResponse struct {
	Page         int                  `json:"page" validate:"min=1"`
	Results      []MultiSearchResult  `json:"results"`
	TotalPages   int                  `json:"total_pages" validate:"min=0"`
	TotalResults int                  `json:"total_results" validate:"min=0"`
}

// SearchRequest represents the parameters for a search request
type SearchRequest struct {
	Query    string `json:"query" validate:"required,min=1"`
	Type     string `json:"type,omitempty"`     // movie, tv, person, all
	Page     int    `json:"page,omitempty"`     // Default: 1
	Language string `json:"language,omitempty"` // Default: ja-JP
	Year     int    `json:"year,omitempty"`     // For movies only
}

// APISearchResponse represents the API response for search requests
type APISearchResponse struct {
	Query        string               `json:"query"`
	Type         string               `json:"type"`
	Page         int                  `json:"page"`
	TotalPages   int                  `json:"total_pages"`
	TotalResults int                  `json:"total_results"`
	Results      []MultiSearchResult  `json:"results"`
	Language     string               `json:"language"`
}

// Validate validates the search request parameters
func (sr *SearchRequest) Validate() error {
	if sr.Query == "" {
		return &ValidationError{Field: "query", Message: "Query parameter is required"}
	}
	
	if sr.Type != "" && sr.Type != "movie" && sr.Type != "tv" && sr.Type != "person" && sr.Type != "all" {
		return &ValidationError{Field: "type", Message: "Type must be one of: movie, tv, person, all"}
	}
	
	if sr.Page < 0 {
		return &ValidationError{Field: "page", Message: "Page must be a positive number"}
	}
	
	if sr.Year < 0 {
		return &ValidationError{Field: "year", Message: "Year must be a positive number"}
	}
	
	return nil
}

// SetDefaults sets default values for optional fields
func (sr *SearchRequest) SetDefaults() {
	if sr.Type == "" {
		sr.Type = "all"
	}
	if sr.Page <= 0 {
		sr.Page = 1
	}
	if sr.Language == "" {
		sr.Language = "ja-JP"
	}
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Error implements the error interface
func (ve *ValidationError) Error() string {
	return ve.Message
}

// ToMovie converts a MultiSearchResult to a Movie (if applicable)
func (msr *MultiSearchResult) ToMovie() *Movie {
	if msr.MediaType != SearchItemTypeMovie {
		return nil
	}
	
	movie := &Movie{
		ID:               msr.ID,
		Popularity:       msr.Popularity,
		GenreIDs:         msr.GenreIDs,
		BackdropPath:     msr.BackdropPath,
		PosterPath:       msr.PosterPath,
		OriginalLanguage: getStringValue(msr.OriginalLanguage),
		Overview:         getStringValue(msr.Overview),
	}
	
	if msr.Adult != nil {
		movie.Adult = *msr.Adult
	}
	if msr.OriginalTitle != nil {
		movie.OriginalTitle = *msr.OriginalTitle
	}
	if msr.ReleaseDate != nil {
		movie.ReleaseDate = msr.ReleaseDate
	}
	if msr.Title != nil {
		movie.Title = *msr.Title
	}
	if msr.Video != nil {
		movie.Video = *msr.Video
	}
	if msr.VoteAverage != nil {
		movie.VoteAverage = *msr.VoteAverage
	}
	if msr.VoteCount != nil {
		movie.VoteCount = *msr.VoteCount
	}
	
	return movie
}

// ToTVShow converts a MultiSearchResult to a TVShow (if applicable)
func (msr *MultiSearchResult) ToTVShow() *TVShow {
	if msr.MediaType != SearchItemTypeTV {
		return nil
	}
	
	tvShow := &TVShow{
		ID:               msr.ID,
		Popularity:       msr.Popularity,
		GenreIDs:         msr.GenreIDs,
		BackdropPath:     msr.BackdropPath,
		PosterPath:       msr.PosterPath,
		OriginalLanguage: getStringValue(msr.OriginalLanguage),
		Overview:         getStringValue(msr.Overview),
		OriginCountry:    msr.OriginCountry,
	}
	
	if msr.Adult != nil {
		tvShow.Adult = *msr.Adult
	}
	if msr.FirstAirDate != nil {
		tvShow.FirstAirDate = msr.FirstAirDate
	}
	if msr.Name != nil {
		tvShow.Name = *msr.Name
	}
	if msr.OriginalName != nil {
		tvShow.OriginalName = *msr.OriginalName
	}
	if msr.VoteAverage != nil {
		tvShow.VoteAverage = *msr.VoteAverage
	}
	if msr.VoteCount != nil {
		tvShow.VoteCount = *msr.VoteCount
	}
	
	return tvShow
}

// ToPerson converts a MultiSearchResult to a Person (if applicable)
func (msr *MultiSearchResult) ToPerson() *Person {
	if msr.MediaType != SearchItemTypePerson {
		return nil
	}
	
	person := &Person{
		ID:         msr.ID,
		Popularity: msr.Popularity,
		KnownFor:   msr.KnownFor,
	}
	
	if msr.Adult != nil {
		person.Adult = *msr.Adult
	}
	if msr.Gender != nil {
		person.Gender = *msr.Gender
	}
	if msr.KnownForDepartment != nil {
		person.KnownForDepartment = *msr.KnownForDepartment
	}
	if msr.Name != nil {
		person.Name = *msr.Name
	}
	if msr.OriginalName != nil {
		person.OriginalName = *msr.OriginalName
	}
	if msr.ProfilePath != nil {
		person.ProfilePath = msr.ProfilePath
	}
	
	return person
}

// Helper function to safely get string value from pointer
func getStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// SearchFilter represents filtering options for search
type SearchFilter struct {
	IncludeAdult bool   `json:"include_adult,omitempty"`
	Year         int    `json:"year,omitempty"`
	PrimaryReleaseYear int `json:"primary_release_year,omitempty"`
	Region       string `json:"region,omitempty"`
}

// SearchOptions represents all search configuration options
type SearchOptions struct {
	Query    string       `json:"query" validate:"required"`
	Type     string       `json:"type,omitempty"`
	Page     int          `json:"page,omitempty"`
	Language string       `json:"language,omitempty"`
	Filter   SearchFilter `json:"filter,omitempty"`
}