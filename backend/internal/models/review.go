// Package models provides review and rating-related data structures for TMDb API and user-generated content.
package models

import "time"

// Review represents a review from TMDb API
type Review struct {
	Author        string      `json:"author" validate:"required"`
	AuthorDetails AuthorDetails `json:"author_details"`
	Content       string      `json:"content" validate:"required"`
	CreatedAt     time.Time   `json:"created_at" validate:"required"`
	ID            string      `json:"id" validate:"required"`
	UpdatedAt     time.Time   `json:"updated_at" validate:"required"`
	URL           string      `json:"url" validate:"required"`
}

// AuthorDetails represents author details in a review
type AuthorDetails struct {
	Name       string   `json:"name"`
	Username   string   `json:"username"`
	AvatarPath *string  `json:"avatar_path"`
	Rating     *float64 `json:"rating"` // User rating out of 10
}

// MovieReviews represents reviews for a movie
type MovieReviews struct {
	ID           int      `json:"id" validate:"required"`
	Page         int      `json:"page" validate:"required,min=1"`
	Results      []Review `json:"results" validate:"required"`
	TotalPages   int      `json:"total_pages" validate:"required,min=0"`
	TotalResults int      `json:"total_results" validate:"required,min=0"`
}

// TVReviews represents reviews for a TV show
type TVReviews struct {
	ID           int      `json:"id" validate:"required"`
	Page         int      `json:"page" validate:"required,min=1"`
	Results      []Review `json:"results" validate:"required"`
	TotalPages   int      `json:"total_pages" validate:"required,min=0"`
	TotalResults int      `json:"total_results" validate:"required,min=0"`
}

// UserReview represents a user-generated review (for future implementation)
type UserReview struct {
	ID        string    `json:"id" validate:"required"`
	UserID    string    `json:"user_id" validate:"required"`
	MediaID   int       `json:"media_id" validate:"required"`
	MediaType string    `json:"media_type" validate:"required,oneof=movie tv"`
	Title     string    `json:"title" validate:"required,max=200"`
	Content   string    `json:"content" validate:"required,max=5000"`
	Rating    float64   `json:"rating" validate:"required,min=1,max=10"`
	Spoiler   bool      `json:"spoiler"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
	// Additional user review fields
	Helpful   int    `json:"helpful"`
	NotHelpful int   `json:"not_helpful"`
	Status    string `json:"status" validate:"required,oneof=pending approved rejected"`
}

// ReviewStats represents aggregated review statistics
type ReviewStats struct {
	MediaID      int     `json:"media_id" validate:"required"`
	MediaType    string  `json:"media_type" validate:"required,oneof=movie tv"`
	TotalReviews int     `json:"total_reviews" validate:"required,min=0"`
	AverageRating float64 `json:"average_rating" validate:"required,min=0,max=10"`
	RatingDistribution RatingDistribution `json:"rating_distribution"`
	// TMDb stats
	TMDbReviews    int     `json:"tmdb_reviews" validate:"required,min=0"`
	TMDbRating     float64 `json:"tmdb_rating" validate:"required,min=0,max=10"`
	// User stats
	UserReviews    int     `json:"user_reviews" validate:"required,min=0"`
	UserRating     float64 `json:"user_rating" validate:"required,min=0,max=10"`
}

// RatingDistribution represents the distribution of ratings
type RatingDistribution struct {
	Rating1  int `json:"rating_1" validate:"required,min=0"`
	Rating2  int `json:"rating_2" validate:"required,min=0"`
	Rating3  int `json:"rating_3" validate:"required,min=0"`
	Rating4  int `json:"rating_4" validate:"required,min=0"`
	Rating5  int `json:"rating_5" validate:"required,min=0"`
	Rating6  int `json:"rating_6" validate:"required,min=0"`
	Rating7  int `json:"rating_7" validate:"required,min=0"`
	Rating8  int `json:"rating_8" validate:"required,min=0"`
	Rating9  int `json:"rating_9" validate:"required,min=0"`
	Rating10 int `json:"rating_10" validate:"required,min=0"`
}

// ReviewRequest represents a request to create a new user review
type ReviewRequest struct {
	MediaID   int     `json:"media_id" validate:"required"`
	MediaType string  `json:"media_type" validate:"required,oneof=movie tv"`
	Title     string  `json:"title" validate:"required,max=200"`
	Content   string  `json:"content" validate:"required,max=5000"`
	Rating    float64 `json:"rating" validate:"required,min=1,max=10"`
	Spoiler   bool    `json:"spoiler"`
}

// ReviewUpdateRequest represents a request to update an existing user review
type ReviewUpdateRequest struct {
	Title   *string  `json:"title,omitempty" validate:"omitempty,max=200"`
	Content *string  `json:"content,omitempty" validate:"omitempty,max=5000"`
	Rating  *float64 `json:"rating,omitempty" validate:"omitempty,min=1,max=10"`
	Spoiler *bool    `json:"spoiler,omitempty"`
}

// ReviewResponse represents the response when creating/updating a review
type ReviewResponse struct {
	Review  UserReview `json:"review"`
	Message string     `json:"message"`
	Success bool       `json:"success"`
}

// ReviewListResponse represents a paginated list of user reviews
type ReviewListResponse struct {
	Reviews      []UserReview `json:"reviews"`
	Page         int          `json:"page" validate:"required,min=1"`
	TotalPages   int          `json:"total_pages" validate:"required,min=0"`
	TotalResults int          `json:"total_results" validate:"required,min=0"`
	HasNext      bool         `json:"has_next"`
	HasPrevious  bool         `json:"has_previous"`
}

// ReviewFilter represents filters for querying reviews
type ReviewFilter struct {
	MediaID     *int     `json:"media_id,omitempty"`
	MediaType   *string  `json:"media_type,omitempty" validate:"omitempty,oneof=movie tv"`
	UserID      *string  `json:"user_id,omitempty"`
	MinRating   *float64 `json:"min_rating,omitempty" validate:"omitempty,min=1,max=10"`
	MaxRating   *float64 `json:"max_rating,omitempty" validate:"omitempty,min=1,max=10"`
	HasSpoiler  *bool    `json:"has_spoiler,omitempty"`
	Status      *string  `json:"status,omitempty" validate:"omitempty,oneof=pending approved rejected"`
	SortBy      string   `json:"sort_by" validate:"omitempty,oneof=created_at updated_at rating helpful"`
	SortOrder   string   `json:"sort_order" validate:"omitempty,oneof=asc desc"`
	Page        int      `json:"page" validate:"required,min=1"`
	Limit       int      `json:"limit" validate:"required,min=1,max=100"`
}

// ReviewModerationRequest represents a request to moderate a review
type ReviewModerationRequest struct {
	ReviewID string `json:"review_id" validate:"required"`
	Action   string `json:"action" validate:"required,oneof=approve reject"`
	Reason   string `json:"reason,omitempty" validate:"omitempty,max=500"`
}

// ReviewHelpfulnessRequest represents a request to mark a review as helpful/not helpful
type ReviewHelpfulnessRequest struct {
	ReviewID string `json:"review_id" validate:"required"`
	UserID   string `json:"user_id" validate:"required"`
	Helpful  bool   `json:"helpful"`
}

// ReviewHelpfulnessResponse represents the response to a helpfulness action
type ReviewHelpfulnessResponse struct {
	ReviewID    string `json:"review_id"`
	Helpful     int    `json:"helpful"`
	NotHelpful  int    `json:"not_helpful"`
	UserAction  string `json:"user_action"` // "helpful", "not_helpful", or "removed"
	Message     string `json:"message"`
	Success     bool   `json:"success"`
}