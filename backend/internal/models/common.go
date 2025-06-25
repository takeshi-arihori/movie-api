// Package models provides data structures for TMDb API responses and internal application entities.
// These models are designed to support JSON marshaling/unmarshaling and validation for the Movie API service.
package models

import "time"

// Genre represents a movie/TV show genre from TMDb API
type Genre struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

// ProductionCompany represents a production company from TMDb API
type ProductionCompany struct {
	ID            int     `json:"id" validate:"required"`
	LogoPath      *string `json:"logo_path"`
	Name          string  `json:"name" validate:"required"`
	OriginCountry string  `json:"origin_country"`
}

// ProductionCountry represents a production country from TMDb API
type ProductionCountry struct {
	ISO31661 string `json:"iso_3166_1" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

// SpokenLanguage represents a spoken language from TMDb API
type SpokenLanguage struct {
	EnglishName string `json:"english_name" validate:"required"`
	ISO6391     string `json:"iso_639_1" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

// Image represents an image from TMDb API (poster, backdrop, profile, etc.)
type Image struct {
	AspectRatio float64 `json:"aspect_ratio"`
	FilePath    string  `json:"file_path" validate:"required"`
	Height      int     `json:"height" validate:"required"`
	ISO6391     *string `json:"iso_639_1"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Width       int     `json:"width" validate:"required"`
}

// Video represents a video from TMDb API (trailer, teaser, etc.)
type Video struct {
	ID           string    `json:"id" validate:"required"`
	ISO6391      string    `json:"iso_639_1" validate:"required"`
	ISO31661     string    `json:"iso_3166_1" validate:"required"`
	Key          string    `json:"key" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	Official     bool      `json:"official"`
	PublishedAt  time.Time `json:"published_at"`
	Site         string    `json:"site" validate:"required"`
	Size         int       `json:"size"`
	Type         string    `json:"type" validate:"required"`
}

// SearchResponse represents a generic paginated search response from TMDb API
type SearchResponse[T any] struct {
	Page         int `json:"page" validate:"min=1"`
	Results      []T `json:"results"`
	TotalPages   int `json:"total_pages" validate:"min=0"`
	TotalResults int `json:"total_results" validate:"min=0"`
}

// ErrorResponse represents an error response from TMDb API
type ErrorResponse struct {
	StatusCode    int    `json:"status_code" validate:"required"`
	StatusMessage string `json:"status_message" validate:"required"`
	Success       bool   `json:"success"`
}

// Collection represents a movie collection from TMDb API
type Collection struct {
	ID           int     `json:"id" validate:"required"`
	Name         string  `json:"name" validate:"required"`
	PosterPath   *string `json:"poster_path"`
	BackdropPath *string `json:"backdrop_path"`
}

// ExternalIDs represents external IDs from TMDb API
type ExternalIDs struct {
	IMDbID      *string `json:"imdb_id"`
	FacebookID  *string `json:"facebook_id"`
	InstagramID *string `json:"instagram_id"`
	TwitterID   *string `json:"twitter_id"`
	WikidataID  *string `json:"wikidata_id"`
}

// Network represents a TV network from TMDb API
type Network struct {
	ID            int     `json:"id" validate:"required"`
	LogoPath      *string `json:"logo_path"`
	Name          string  `json:"name" validate:"required"`
	OriginCountry string  `json:"origin_country"`
}

// Keywords represents keywords from TMDb API
type Keywords struct {
	Keywords []Keyword `json:"keywords"`
	Results  []Keyword `json:"results"` // For TV shows, keywords are in "results"
}

// Keyword represents a single keyword from TMDb API
type Keyword struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

// ReleaseDates represents release dates information from TMDb API
type ReleaseDates struct {
	Results []ReleaseDateCountry `json:"results"`
}

// ReleaseDateCountry represents release dates for a specific country
type ReleaseDateCountry struct {
	ISO31661     string        `json:"iso_3166_1" validate:"required"`
	ReleaseDates []ReleaseDate `json:"release_dates"`
}

// ReleaseDate represents a single release date
type ReleaseDate struct {
	Certification string     `json:"certification"`
	ISO6391       string     `json:"iso_639_1"`
	Note          string     `json:"note"`
	ReleaseDate   *time.Time `json:"release_date"`
	Type          int        `json:"type" validate:"required"`
}

// ContentRatings represents content ratings for TV shows from TMDb API
type ContentRatings struct {
	Results []ContentRating `json:"results"`
}

// ContentRating represents a single content rating
type ContentRating struct {
	Descriptors []string `json:"descriptors"`
	ISO31661    string   `json:"iso_3166_1" validate:"required"`
	Rating      string   `json:"rating" validate:"required"`
}