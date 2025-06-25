package models

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
)

// Test data constants
const (
	testMovieJSON = `{
		"adult": false,
		"backdrop_path": "/path/to/backdrop.jpg",
		"genre_ids": [28, 12, 878],
		"id": 550,
		"original_language": "en",
		"original_title": "Fight Club",
		"overview": "A ticking-time-bomb insomniac and a slippery soap salesman channel primal male aggression into a shocking new form of therapy.",
		"popularity": 61.416,
		"poster_path": "/path/to/poster.jpg",
		"release_date": "1999-10-15",
		"title": "Fight Club",
		"video": false,
		"vote_average": 8.4,
		"vote_count": 26280
	}`

	testTVShowJSON = `{
		"adult": false,
		"backdrop_path": "/path/to/backdrop.jpg",
		"genre_ids": [18, 80],
		"id": 1396,
		"origin_country": ["US"],
		"original_language": "en",
		"original_name": "Breaking Bad",
		"overview": "A high school chemistry teacher diagnosed with inoperable lung cancer turns to manufacturing and selling methamphetamine in order to secure his family's future.",
		"popularity": 449.316,
		"poster_path": "/path/to/poster.jpg",
		"first_air_date": "2008-01-20",
		"name": "Breaking Bad",
		"vote_average": 8.9,
		"vote_count": 12859
	}`

	testPersonJSON = `{
		"adult": false,
		"gender": 2,
		"id": 819,
		"known_for": [],
		"known_for_department": "Acting",
		"name": "Edward Norton",
		"original_name": "Edward Norton",
		"popularity": 7.861,
		"profile_path": "/path/to/profile.jpg"
	}`

	testReviewJSON = `{
		"author": "John Doe",
		"author_details": {
			"name": "John Doe",
			"username": "johndoe",
			"avatar_path": "/path/to/avatar.jpg",
			"rating": 8.5
		},
		"content": "This is a great movie!",
		"created_at": "2023-01-01T00:00:00.000Z",
		"id": "63f1e8b2c44d3e007e0c5b1a",
		"updated_at": "2023-01-01T00:00:00.000Z",
		"url": "https://www.themoviedb.org/review/63f1e8b2c44d3e007e0c5b1a"
	}`
)

// TestMovieJSONMarshaling tests Movie struct JSON marshaling/unmarshaling
func TestMovieJSONMarshaling(t *testing.T) {
	// Test unmarshaling
	var movie Movie
	if err := json.Unmarshal([]byte(testMovieJSON), &movie); err != nil {
		t.Fatalf("Failed to unmarshal movie JSON: %v", err)
	}

	// Verify required fields
	if movie.ID != 550 {
		t.Errorf("Expected ID 550, got %d", movie.ID)
	}
	if movie.Title != "Fight Club" {
		t.Errorf("Expected title 'Fight Club', got '%s'", movie.Title)
	}
	if movie.OriginalTitle != "Fight Club" {
		t.Errorf("Expected original title 'Fight Club', got '%s'", movie.OriginalTitle)
	}
	if movie.OriginalLanguage != "en" {
		t.Errorf("Expected original language 'en', got '%s'", movie.OriginalLanguage)
	}

	// Test marshaling
	data, err := json.Marshal(movie)
	if err != nil {
		t.Fatalf("Failed to marshal movie: %v", err)
	}

	// Test unmarshaling the marshaled data
	var movie2 Movie
	if err := json.Unmarshal(data, &movie2); err != nil {
		t.Fatalf("Failed to unmarshal marshaled movie: %v", err)
	}

	// Compare critical fields
	if movie.ID != movie2.ID {
		t.Errorf("ID mismatch after marshal/unmarshal: %d != %d", movie.ID, movie2.ID)
	}
	if movie.Title != movie2.Title {
		t.Errorf("Title mismatch after marshal/unmarshal: %s != %s", movie.Title, movie2.Title)
	}
}

// TestTVShowJSONMarshaling tests TVShow struct JSON marshaling/unmarshaling
func TestTVShowJSONMarshaling(t *testing.T) {
	// Test unmarshaling
	var tvShow TVShow
	if err := json.Unmarshal([]byte(testTVShowJSON), &tvShow); err != nil {
		t.Fatalf("Failed to unmarshal TV show JSON: %v", err)
	}

	// Verify required fields
	if tvShow.ID != 1396 {
		t.Errorf("Expected ID 1396, got %d", tvShow.ID)
	}
	if tvShow.Name != "Breaking Bad" {
		t.Errorf("Expected name 'Breaking Bad', got '%s'", tvShow.Name)
	}
	if tvShow.OriginalName != "Breaking Bad" {
		t.Errorf("Expected original name 'Breaking Bad', got '%s'", tvShow.OriginalName)
	}
	if tvShow.OriginalLanguage != "en" {
		t.Errorf("Expected original language 'en', got '%s'", tvShow.OriginalLanguage)
	}

	// Test marshaling
	data, err := json.Marshal(tvShow)
	if err != nil {
		t.Fatalf("Failed to marshal TV show: %v", err)
	}

	// Test unmarshaling the marshaled data
	var tvShow2 TVShow
	if err := json.Unmarshal(data, &tvShow2); err != nil {
		t.Fatalf("Failed to unmarshal marshaled TV show: %v", err)
	}

	// Compare critical fields
	if tvShow.ID != tvShow2.ID {
		t.Errorf("ID mismatch after marshal/unmarshal: %d != %d", tvShow.ID, tvShow2.ID)
	}
	if tvShow.Name != tvShow2.Name {
		t.Errorf("Name mismatch after marshal/unmarshal: %s != %s", tvShow.Name, tvShow2.Name)
	}
}

// TestPersonJSONMarshaling tests Person struct JSON marshaling/unmarshaling
func TestPersonJSONMarshaling(t *testing.T) {
	// Test unmarshaling
	var person Person
	if err := json.Unmarshal([]byte(testPersonJSON), &person); err != nil {
		t.Fatalf("Failed to unmarshal person JSON: %v", err)
	}

	// Verify required fields
	if person.ID != 819 {
		t.Errorf("Expected ID 819, got %d", person.ID)
	}
	if person.Name != "Edward Norton" {
		t.Errorf("Expected name 'Edward Norton', got '%s'", person.Name)
	}

	// Test marshaling
	data, err := json.Marshal(person)
	if err != nil {
		t.Fatalf("Failed to marshal person: %v", err)
	}

	// Test unmarshaling the marshaled data
	var person2 Person
	if err := json.Unmarshal(data, &person2); err != nil {
		t.Fatalf("Failed to unmarshal marshaled person: %v", err)
	}

	// Compare critical fields
	if person.ID != person2.ID {
		t.Errorf("ID mismatch after marshal/unmarshal: %d != %d", person.ID, person2.ID)
	}
	if person.Name != person2.Name {
		t.Errorf("Name mismatch after marshal/unmarshal: %s != %s", person.Name, person2.Name)
	}
}

// TestReviewJSONMarshaling tests Review struct JSON marshaling/unmarshaling
func TestReviewJSONMarshaling(t *testing.T) {
	// Test unmarshaling
	var review Review
	if err := json.Unmarshal([]byte(testReviewJSON), &review); err != nil {
		t.Fatalf("Failed to unmarshal review JSON: %v", err)
	}

	// Verify required fields
	if review.ID != "63f1e8b2c44d3e007e0c5b1a" {
		t.Errorf("Expected ID '63f1e8b2c44d3e007e0c5b1a', got '%s'", review.ID)
	}
	if review.Author != "John Doe" {
		t.Errorf("Expected author 'John Doe', got '%s'", review.Author)
	}
	if review.Content != "This is a great movie!" {
		t.Errorf("Expected content 'This is a great movie!', got '%s'", review.Content)
	}

	// Test marshaling
	data, err := json.Marshal(review)
	if err != nil {
		t.Fatalf("Failed to marshal review: %v", err)
	}

	// Test unmarshaling the marshaled data
	var review2 Review
	if err := json.Unmarshal(data, &review2); err != nil {
		t.Fatalf("Failed to unmarshal marshaled review: %v", err)
	}

	// Compare critical fields
	if review.ID != review2.ID {
		t.Errorf("ID mismatch after marshal/unmarshal: %s != %s", review.ID, review2.ID)
	}
	if review.Author != review2.Author {
		t.Errorf("Author mismatch after marshal/unmarshal: %s != %s", review.Author, review2.Author)
	}
}

// TestGenreStruct tests Genre struct
func TestGenreStruct(t *testing.T) {
	genre := Genre{
		ID:   28,
		Name: "Action",
	}

	data, err := json.Marshal(genre)
	if err != nil {
		t.Fatalf("Failed to marshal genre: %v", err)
	}

	var genre2 Genre
	if err := json.Unmarshal(data, &genre2); err != nil {
		t.Fatalf("Failed to unmarshal genre: %v", err)
	}

	if genre.ID != genre2.ID {
		t.Errorf("ID mismatch: %d != %d", genre.ID, genre2.ID)
	}
	if genre.Name != genre2.Name {
		t.Errorf("Name mismatch: %s != %s", genre.Name, genre2.Name)
	}
}

// TestSearchResponse tests generic SearchResponse
func TestSearchResponse(t *testing.T) {
	movies := []Movie{
		{ID: 1, Title: "Movie 1", OriginalTitle: "Movie 1", OriginalLanguage: "en"},
		{ID: 2, Title: "Movie 2", OriginalTitle: "Movie 2", OriginalLanguage: "en"},
	}

	response := SearchResponse[Movie]{
		Page:         1,
		Results:      movies,
		TotalPages:   1,
		TotalResults: 2,
	}

	data, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Failed to marshal search response: %v", err)
	}

	var response2 SearchResponse[Movie]
	if err := json.Unmarshal(data, &response2); err != nil {
		t.Fatalf("Failed to unmarshal search response: %v", err)
	}

	if response.Page != response2.Page {
		t.Errorf("Page mismatch: %d != %d", response.Page, response2.Page)
	}
	if len(response.Results) != len(response2.Results) {
		t.Errorf("Results length mismatch: %d != %d", len(response.Results), len(response2.Results))
	}
	if response.TotalResults != response2.TotalResults {
		t.Errorf("TotalResults mismatch: %d != %d", response.TotalResults, response2.TotalResults)
	}
}

// TestMovieDetails tests MovieDetails struct
func TestMovieDetails(t *testing.T) {
	movieDetails := MovieDetails{
		ID:               550,
		Title:            "Fight Club",
		OriginalTitle:    "Fight Club",
		OriginalLanguage: "en",
		Genres: []Genre{
			{ID: 18, Name: "Drama"},
			{ID: 53, Name: "Thriller"},
		},
		Runtime: intPtr(139),
		Status:  "Released",
	}

	data, err := json.Marshal(movieDetails)
	if err != nil {
		t.Fatalf("Failed to marshal movie details: %v", err)
	}

	var movieDetails2 MovieDetails
	if err := json.Unmarshal(data, &movieDetails2); err != nil {
		t.Fatalf("Failed to unmarshal movie details: %v", err)
	}

	if movieDetails.ID != movieDetails2.ID {
		t.Errorf("ID mismatch: %d != %d", movieDetails.ID, movieDetails2.ID)
	}
	if len(movieDetails.Genres) != len(movieDetails2.Genres) {
		t.Errorf("Genres length mismatch: %d != %d", len(movieDetails.Genres), len(movieDetails2.Genres))
	}
	if movieDetails.Runtime != nil && movieDetails2.Runtime != nil {
		if *movieDetails.Runtime != *movieDetails2.Runtime {
			t.Errorf("Runtime mismatch: %d != %d", *movieDetails.Runtime, *movieDetails2.Runtime)
		}
	}
}

// TestTVShowDetails tests TVShowDetails struct
func TestTVShowDetails(t *testing.T) {
	tvDetails := TVShowDetails{
		ID:               1396,
		Name:             "Breaking Bad",
		OriginalName:     "Breaking Bad",
		OriginalLanguage: "en",
		NumberOfSeasons:  5,
		NumberOfEpisodes: 62,
		Status:           "Ended",
		Type:             "Scripted",
		Seasons: []Season{
			{
				ID:           3572,
				Name:         "Season 1",
				SeasonNumber: 1,
				EpisodeCount: 7,
			},
		},
	}

	data, err := json.Marshal(tvDetails)
	if err != nil {
		t.Fatalf("Failed to marshal TV details: %v", err)
	}

	var tvDetails2 TVShowDetails
	if err := json.Unmarshal(data, &tvDetails2); err != nil {
		t.Fatalf("Failed to unmarshal TV details: %v", err)
	}

	if tvDetails.ID != tvDetails2.ID {
		t.Errorf("ID mismatch: %d != %d", tvDetails.ID, tvDetails2.ID)
	}
	if tvDetails.NumberOfSeasons != tvDetails2.NumberOfSeasons {
		t.Errorf("NumberOfSeasons mismatch: %d != %d", tvDetails.NumberOfSeasons, tvDetails2.NumberOfSeasons)
	}
	if len(tvDetails.Seasons) != len(tvDetails2.Seasons) {
		t.Errorf("Seasons length mismatch: %d != %d", len(tvDetails.Seasons), len(tvDetails2.Seasons))
	}
}

// TestPersonDetails tests PersonDetails struct
func TestPersonDetails(t *testing.T) {
	personDetails := PersonDetails{
		ID:                 819,
		Name:               "Edward Norton",
		KnownForDepartment: "Acting",
		Gender:             2,
		Biography:          "Edward Harrison Norton is an American actor and filmmaker.",
		Birthday:           stringPtr("1969-08-18"),
		PlaceOfBirth:       stringPtr("Boston, Massachusetts, USA"),
		AlsoKnownAs:        []string{"Ed Norton"},
	}

	data, err := json.Marshal(personDetails)
	if err != nil {
		t.Fatalf("Failed to marshal person details: %v", err)
	}

	var personDetails2 PersonDetails
	if err := json.Unmarshal(data, &personDetails2); err != nil {
		t.Fatalf("Failed to unmarshal person details: %v", err)
	}

	if personDetails.ID != personDetails2.ID {
		t.Errorf("ID mismatch: %d != %d", personDetails.ID, personDetails2.ID)
	}
	if personDetails.Name != personDetails2.Name {
		t.Errorf("Name mismatch: %s != %s", personDetails.Name, personDetails2.Name)
	}
	if len(personDetails.AlsoKnownAs) != len(personDetails2.AlsoKnownAs) {
		t.Errorf("AlsoKnownAs length mismatch: %d != %d", len(personDetails.AlsoKnownAs), len(personDetails2.AlsoKnownAs))
	}
}

// TestUserReview tests UserReview struct
func TestUserReview(t *testing.T) {
	now := time.Now()
	userReview := UserReview{
		ID:        "user-review-1",
		UserID:    "user-123",
		MediaID:   550,
		MediaType: "movie",
		Title:     "Amazing Movie!",
		Content:   "This movie is absolutely fantastic. Great acting and storyline.",
		Rating:    8.5,
		Spoiler:   false,
		CreatedAt: now,
		UpdatedAt: now,
		Helpful:   10,
		NotHelpful: 2,
		Status:    "approved",
	}

	data, err := json.Marshal(userReview)
	if err != nil {
		t.Fatalf("Failed to marshal user review: %v", err)
	}

	var userReview2 UserReview
	if err := json.Unmarshal(data, &userReview2); err != nil {
		t.Fatalf("Failed to unmarshal user review: %v", err)
	}

	if userReview.ID != userReview2.ID {
		t.Errorf("ID mismatch: %s != %s", userReview.ID, userReview2.ID)
	}
	if userReview.MediaID != userReview2.MediaID {
		t.Errorf("MediaID mismatch: %d != %d", userReview.MediaID, userReview2.MediaID)
	}
	if userReview.Rating != userReview2.Rating {
		t.Errorf("Rating mismatch: %f != %f", userReview.Rating, userReview2.Rating)
	}
}

// TestCastMember tests CastMember struct
func TestCastMember(t *testing.T) {
	castMember := CastMember{
		ID:                 819,
		Name:               "Edward Norton",
		Character:          "The Narrator",
		CreditID:           "52fe4250c3a36847f80149f3",
		Order:              0,
		KnownForDepartment: "Acting",
		Gender:             intPtr(2),
		ProfilePath:        stringPtr("/path/to/profile.jpg"),
	}

	data, err := json.Marshal(castMember)
	if err != nil {
		t.Fatalf("Failed to marshal cast member: %v", err)
	}

	var castMember2 CastMember
	if err := json.Unmarshal(data, &castMember2); err != nil {
		t.Fatalf("Failed to unmarshal cast member: %v", err)
	}

	if castMember.ID != castMember2.ID {
		t.Errorf("ID mismatch: %d != %d", castMember.ID, castMember2.ID)
	}
	if castMember.Character != castMember2.Character {
		t.Errorf("Character mismatch: %s != %s", castMember.Character, castMember2.Character)
	}
	if castMember.Order != castMember2.Order {
		t.Errorf("Order mismatch: %d != %d", castMember.Order, castMember2.Order)
	}
}

// Helper functions for pointer types
func intPtr(i int) *int {
	return &i
}

func stringPtr(s string) *string {
	return &s
}

func float64Ptr(f float64) *float64 {
	return &f
}

func boolPtr(b bool) *bool {
	return &b
}

// TestPointerFields tests that pointer fields work correctly
func TestPointerFields(t *testing.T) {
	// Test with nil pointer fields
	movie := Movie{
		ID:               1,
		Title:            "Test Movie",
		OriginalTitle:    "Test Movie",
		OriginalLanguage: "en",
		BackdropPath:     nil,
		PosterPath:       nil,
		ReleaseDate:      nil,
	}

	data, err := json.Marshal(movie)
	if err != nil {
		t.Fatalf("Failed to marshal movie with nil pointers: %v", err)
	}

	var movie2 Movie
	if err := json.Unmarshal(data, &movie2); err != nil {
		t.Fatalf("Failed to unmarshal movie with nil pointers: %v", err)
	}

	if movie2.BackdropPath != nil {
		t.Errorf("Expected nil BackdropPath, got %v", movie2.BackdropPath)
	}
	if movie2.PosterPath != nil {
		t.Errorf("Expected nil PosterPath, got %v", movie2.PosterPath)
	}

	// Test with non-nil pointer fields
	backdropPath := "/path/to/backdrop.jpg"
	posterPath := "/path/to/poster.jpg"
	releaseDate := "2023-01-01"

	movie3 := Movie{
		ID:               2,
		Title:            "Test Movie 2",
		OriginalTitle:    "Test Movie 2",
		OriginalLanguage: "en",
		BackdropPath:     &backdropPath,
		PosterPath:       &posterPath,
		ReleaseDate:      &releaseDate,
	}

	data2, err := json.Marshal(movie3)
	if err != nil {
		t.Fatalf("Failed to marshal movie with non-nil pointers: %v", err)
	}

	var movie4 Movie
	if err := json.Unmarshal(data2, &movie4); err != nil {
		t.Fatalf("Failed to unmarshal movie with non-nil pointers: %v", err)
	}

	if movie4.BackdropPath == nil || *movie4.BackdropPath != backdropPath {
		t.Errorf("BackdropPath mismatch: expected %s, got %v", backdropPath, movie4.BackdropPath)
	}
	if movie4.PosterPath == nil || *movie4.PosterPath != posterPath {
		t.Errorf("PosterPath mismatch: expected %s, got %v", posterPath, movie4.PosterPath)
	}
	if movie4.ReleaseDate == nil || *movie4.ReleaseDate != releaseDate {
		t.Errorf("ReleaseDate mismatch: expected %s, got %v", releaseDate, movie4.ReleaseDate)
	}
}

// Validation tests
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// TestMovieValidation tests Movie struct validation
func TestMovieValidation(t *testing.T) {
	// Valid movie
	validMovie := Movie{
		ID:               550,
		Title:            "Fight Club",
		OriginalTitle:    "Fight Club",
		OriginalLanguage: "en",
	}

	if err := validate.Struct(validMovie); err != nil {
		t.Errorf("Valid movie failed validation: %v", err)
	}

	// Invalid movie - missing required fields
	invalidMovie := Movie{
		Title: "Fight Club", // Missing ID, OriginalTitle, OriginalLanguage
	}

	if err := validate.Struct(invalidMovie); err == nil {
		t.Error("Invalid movie passed validation when it should have failed")
	}

	// Test individual field validation
	emptyTitleMovie := Movie{
		ID:               550,
		Title:            "", // Empty title should fail
		OriginalTitle:    "Fight Club",
		OriginalLanguage: "en",
	}

	if err := validate.Struct(emptyTitleMovie); err == nil {
		t.Error("Movie with empty title passed validation when it should have failed")
	}
}

// TestTVShowValidation tests TVShow struct validation
func TestTVShowValidation(t *testing.T) {
	// Valid TV show
	validTVShow := TVShow{
		ID:               1396,
		Name:             "Breaking Bad",
		OriginalName:     "Breaking Bad",
		OriginalLanguage: "en",
	}

	if err := validate.Struct(validTVShow); err != nil {
		t.Errorf("Valid TV show failed validation: %v", err)
	}

	// Invalid TV show - missing required fields
	invalidTVShow := TVShow{
		Name: "Breaking Bad", // Missing ID, OriginalName, OriginalLanguage
	}

	if err := validate.Struct(invalidTVShow); err == nil {
		t.Error("Invalid TV show passed validation when it should have failed")
	}
}

// TestPersonValidation tests Person struct validation
func TestPersonValidation(t *testing.T) {
	// Valid person
	validPerson := Person{
		ID:   819,
		Name: "Edward Norton",
	}

	if err := validate.Struct(validPerson); err != nil {
		t.Errorf("Valid person failed validation: %v", err)
	}

	// Invalid person - missing required fields
	invalidPerson := Person{
		Name: "Edward Norton", // Missing ID
	}

	if err := validate.Struct(invalidPerson); err == nil {
		t.Error("Invalid person passed validation when it should have failed")
	}

	// Person with empty name
	emptyNamePerson := Person{
		ID:   819,
		Name: "", // Empty name should fail
	}

	if err := validate.Struct(emptyNamePerson); err == nil {
		t.Error("Person with empty name passed validation when it should have failed")
	}
}

// TestReviewValidation tests Review struct validation
func TestReviewValidation(t *testing.T) {
	now := time.Now()

	// Valid review
	validReview := Review{
		ID:        "review-123",
		Author:    "John Doe",
		Content:   "Great movie!",
		CreatedAt: now,
		UpdatedAt: now,
		URL:       "https://example.com/review/123",
	}

	if err := validate.Struct(validReview); err != nil {
		t.Errorf("Valid review failed validation: %v", err)
	}

	// Invalid review - missing required fields
	invalidReview := Review{
		Author: "John Doe", // Missing ID, Content, CreatedAt, UpdatedAt, URL
	}

	if err := validate.Struct(invalidReview); err == nil {
		t.Error("Invalid review passed validation when it should have failed")
	}
}

// TestUserReviewValidation tests UserReview struct validation
func TestUserReviewValidation(t *testing.T) {
	now := time.Now()

	// Valid user review
	validUserReview := UserReview{
		ID:        "user-review-123",
		UserID:    "user-456",
		MediaID:   550,
		MediaType: "movie",
		Title:     "Amazing Film",
		Content:   "This movie is absolutely fantastic.",
		Rating:    8.5,
		CreatedAt: now,
		UpdatedAt: now,
		Status:    "approved",
	}

	if err := validate.Struct(validUserReview); err != nil {
		t.Errorf("Valid user review failed validation: %v", err)
	}

	// Invalid media type
	invalidMediaTypeReview := UserReview{
		ID:        "user-review-123",
		UserID:    "user-456",
		MediaID:   550,
		MediaType: "invalid", // Should be "movie" or "tv"
		Title:     "Amazing Film",
		Content:   "This movie is absolutely fantastic.",
		Rating:    8.5,
		CreatedAt: now,
		UpdatedAt: now,
		Status:    "approved",
	}

	if err := validate.Struct(invalidMediaTypeReview); err == nil {
		t.Error("User review with invalid media type passed validation when it should have failed")
	}

	// Invalid rating - too high
	invalidRatingReview := UserReview{
		ID:        "user-review-123",
		UserID:    "user-456",
		MediaID:   550,
		MediaType: "movie",
		Title:     "Amazing Film",
		Content:   "This movie is absolutely fantastic.",
		Rating:    15.0, // Should be between 1 and 10
		CreatedAt: now,
		UpdatedAt: now,
		Status:    "approved",
	}

	if err := validate.Struct(invalidRatingReview); err == nil {
		t.Error("User review with invalid rating passed validation when it should have failed")
	}

	// Invalid rating - too low
	invalidLowRatingReview := UserReview{
		ID:        "user-review-123",
		UserID:    "user-456",
		MediaID:   550,
		MediaType: "movie",
		Title:     "Amazing Film",
		Content:   "This movie is absolutely fantastic.",
		Rating:    0.5, // Should be between 1 and 10
		CreatedAt: now,
		UpdatedAt: now,
		Status:    "approved",
	}

	if err := validate.Struct(invalidLowRatingReview); err == nil {
		t.Error("User review with invalid low rating passed validation when it should have failed")
	}

	// Invalid status
	invalidStatusReview := UserReview{
		ID:        "user-review-123",
		UserID:    "user-456",
		MediaID:   550,
		MediaType: "movie",
		Title:     "Amazing Film",
		Content:   "This movie is absolutely fantastic.",
		Rating:    8.5,
		CreatedAt: now,
		UpdatedAt: now,
		Status:    "invalid", // Should be "pending", "approved", or "rejected"
	}

	if err := validate.Struct(invalidStatusReview); err == nil {
		t.Error("User review with invalid status passed validation when it should have failed")
	}
}

// TestGenreValidation tests Genre struct validation
func TestGenreValidation(t *testing.T) {
	// Valid genre
	validGenre := Genre{
		ID:   28,
		Name: "Action",
	}

	if err := validate.Struct(validGenre); err != nil {
		t.Errorf("Valid genre failed validation: %v", err)
	}

	// Invalid genre - missing required fields
	invalidGenre := Genre{
		Name: "Action", // Missing ID
	}

	if err := validate.Struct(invalidGenre); err == nil {
		t.Error("Invalid genre passed validation when it should have failed")
	}
}

// TestSearchResponseValidation tests SearchResponse struct validation
func TestSearchResponseValidation(t *testing.T) {
	// Valid search response
	validResponse := SearchResponse[Movie]{
		Page:         1,
		Results:      []Movie{},
		TotalPages:   1,
		TotalResults: 0,
	}

	if err := validate.Struct(validResponse); err != nil {
		t.Errorf("Valid search response failed validation: %v", err)
	}

	// Invalid search response - page 0
	invalidPageResponse := SearchResponse[Movie]{
		Page:         0, // Should be min=1
		Results:      []Movie{},
		TotalPages:   1,
		TotalResults: 0,
	}

	if err := validate.Struct(invalidPageResponse); err == nil {
		t.Error("Search response with page 0 passed validation when it should have failed")
	}

	// Invalid search response - negative total pages
	invalidTotalPagesResponse := SearchResponse[Movie]{
		Page:         1,
		Results:      []Movie{},
		TotalPages:   -1, // Should be min=0
		TotalResults: 0,
	}

	if err := validate.Struct(invalidTotalPagesResponse); err == nil {
		t.Error("Search response with negative total pages passed validation when it should have failed")
	}
}

// TestCastMemberValidation tests CastMember struct validation
func TestCastMemberValidation(t *testing.T) {
	// Valid cast member
	validCast := CastMember{
		ID:       819,
		Name:     "Edward Norton",
		CreditID: "52fe4250c3a36847f80149f3",
	}

	if err := validate.Struct(validCast); err != nil {
		t.Errorf("Valid cast member failed validation: %v", err)
	}

	// Invalid cast member - missing required fields
	invalidCast := CastMember{
		Name: "Edward Norton", // Missing ID and CreditID
	}

	if err := validate.Struct(invalidCast); err == nil {
		t.Error("Invalid cast member passed validation when it should have failed")
	}
}