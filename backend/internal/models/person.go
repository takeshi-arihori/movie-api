// Package models provides person-related data structures for TMDb API integration.
package models

// Person represents a person from TMDb API (used in search results and lists)
type Person struct {
	Adult              bool     `json:"adult"`
	Gender             int      `json:"gender"`
	ID                 int      `json:"id" validate:"required"`
	KnownFor           []KnownForItem `json:"known_for"`
	KnownForDepartment string   `json:"known_for_department"`
	Name               string   `json:"name" validate:"required"`
	OriginalName       string   `json:"original_name"`
	Popularity         float64  `json:"popularity"`
	ProfilePath        *string  `json:"profile_path"`
}

// PersonDetails represents detailed person information from TMDb API
type PersonDetails struct {
	Adult              bool     `json:"adult"`
	AlsoKnownAs        []string `json:"also_known_as"`
	Biography          string   `json:"biography"`
	Birthday           *string  `json:"birthday"` // Format: YYYY-MM-DD
	Deathday           *string  `json:"deathday"` // Format: YYYY-MM-DD
	Gender             int      `json:"gender"`
	Homepage           *string  `json:"homepage"`
	ID                 int      `json:"id" validate:"required"`
	IMDbID             *string  `json:"imdb_id"`
	KnownForDepartment string   `json:"known_for_department"`
	Name               string   `json:"name" validate:"required"`
	PlaceOfBirth       *string  `json:"place_of_birth"`
	Popularity         float64  `json:"popularity"`
	ProfilePath        *string  `json:"profile_path"`
}

// KnownForItem represents an item a person is known for (can be movie or TV show)
type KnownForItem struct {
	Adult            *bool    `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id" validate:"required"`
	MediaType        string   `json:"media_type" validate:"required"` // "movie" or "tv"
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    *string  `json:"original_title,omitempty"` // For movies
	OriginalName     *string  `json:"original_name,omitempty"`  // For TV shows
	Overview         string   `json:"overview"`
	PosterPath       *string  `json:"poster_path"`
	ReleaseDate      *string  `json:"release_date,omitempty"`   // For movies, Format: YYYY-MM-DD
	FirstAirDate     *string  `json:"first_air_date,omitempty"` // For TV shows, Format: YYYY-MM-DD
	Title            *string  `json:"title,omitempty"`          // For movies
	Name             *string  `json:"name,omitempty"`           // For TV shows
	Video            *bool    `json:"video,omitempty"`          // For movies
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	OriginCountry    []string `json:"origin_country,omitempty"` // For TV shows
}

// PersonSearchResponse represents a search response for people
type PersonSearchResponse = SearchResponse[Person]

// PersonMovieCredits represents movie credits for a person
type PersonMovieCredits struct {
	Cast []PersonMovieCast `json:"cast"`
	Crew []PersonMovieCrew `json:"crew"`
	ID   int               `json:"id" validate:"required"`
}

// PersonMovieCast represents a movie cast credit for a person
type PersonMovieCast struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id" validate:"required"`
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    string   `json:"original_title"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	ReleaseDate      *string  `json:"release_date"` // Format: YYYY-MM-DD
	Title            string   `json:"title" validate:"required"`
	Video            bool     `json:"video"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	Character        string   `json:"character"`
	CreditID         string   `json:"credit_id" validate:"required"`
	Order            int      `json:"order"`
}

// PersonMovieCrew represents a movie crew credit for a person
type PersonMovieCrew struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id" validate:"required"`
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    string   `json:"original_title"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	ReleaseDate      *string  `json:"release_date"` // Format: YYYY-MM-DD
	Title            string   `json:"title" validate:"required"`
	Video            bool     `json:"video"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	CreditID         string   `json:"credit_id" validate:"required"`
	Department       string   `json:"department" validate:"required"`
	Job              string   `json:"job" validate:"required"`
}

// PersonTVCredits represents TV credits for a person
type PersonTVCredits struct {
	Cast []PersonTVCast `json:"cast"`
	Crew []PersonTVCrew `json:"crew"`
	ID   int            `json:"id" validate:"required"`
}

// PersonTVCast represents a TV cast credit for a person
type PersonTVCast struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id" validate:"required"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	FirstAirDate     *string  `json:"first_air_date"` // Format: YYYY-MM-DD
	Name             string   `json:"name" validate:"required"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	Character        string   `json:"character"`
	CreditID         string   `json:"credit_id" validate:"required"`
	EpisodeCount     int      `json:"episode_count"`
}

// PersonTVCrew represents a TV crew credit for a person
type PersonTVCrew struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id" validate:"required"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	FirstAirDate     *string  `json:"first_air_date"` // Format: YYYY-MM-DD
	Name             string   `json:"name" validate:"required"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	CreditID         string   `json:"credit_id" validate:"required"`
	Department       string   `json:"department" validate:"required"`
	EpisodeCount     int      `json:"episode_count"`
	Job              string   `json:"job" validate:"required"`
}

// PersonCombinedCredits represents combined movie and TV credits for a person
type PersonCombinedCredits struct {
	Cast []PersonCombinedCast `json:"cast"`
	Crew []PersonCombinedCrew `json:"crew"`
	ID   int                  `json:"id" validate:"required"`
}

// PersonCombinedCast represents a combined cast credit (movie or TV)
type PersonCombinedCast struct {
	Adult            *bool    `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id" validate:"required"`
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    *string  `json:"original_title,omitempty"` // For movies
	OriginalName     *string  `json:"original_name,omitempty"`  // For TV shows
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	ReleaseDate      *string  `json:"release_date,omitempty"`   // For movies
	FirstAirDate     *string  `json:"first_air_date,omitempty"` // For TV shows
	Title            *string  `json:"title,omitempty"`          // For movies
	Name             *string  `json:"name,omitempty"`           // For TV shows
	Video            *bool    `json:"video,omitempty"`          // For movies
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	Character        string   `json:"character"`
	CreditID         string   `json:"credit_id" validate:"required"`
	Order            *int     `json:"order,omitempty"`          // For movies
	MediaType        string   `json:"media_type" validate:"required"`
	OriginCountry    []string `json:"origin_country,omitempty"` // For TV shows
	EpisodeCount     *int     `json:"episode_count,omitempty"`  // For TV shows
}

// PersonCombinedCrew represents a combined crew credit (movie or TV)
type PersonCombinedCrew struct {
	Adult            *bool    `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id" validate:"required"`
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    *string  `json:"original_title,omitempty"` // For movies
	OriginalName     *string  `json:"original_name,omitempty"`  // For TV shows
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	ReleaseDate      *string  `json:"release_date,omitempty"`   // For movies
	FirstAirDate     *string  `json:"first_air_date,omitempty"` // For TV shows
	Title            *string  `json:"title,omitempty"`          // For movies
	Name             *string  `json:"name,omitempty"`           // For TV shows
	Video            *bool    `json:"video,omitempty"`          // For movies
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	CreditID         string   `json:"credit_id" validate:"required"`
	Department       string   `json:"department" validate:"required"`
	Job              string   `json:"job" validate:"required"`
	MediaType        string   `json:"media_type" validate:"required"`
	OriginCountry    []string `json:"origin_country,omitempty"` // For TV shows
	EpisodeCount     *int     `json:"episode_count,omitempty"`  // For TV shows
}

// PersonImages represents images for a person
type PersonImages struct {
	ID       int     `json:"id" validate:"required"`
	Profiles []Image `json:"profiles"`
}

// PersonTranslations represents translations for a person
type PersonTranslations struct {
	ID           int                  `json:"id" validate:"required"`
	Translations []PersonTranslation  `json:"translations"`
}

// PersonTranslation represents a single person translation
type PersonTranslation struct {
	ISO31661    string                  `json:"iso_3166_1" validate:"required"`
	ISO6391     string                  `json:"iso_639_1" validate:"required"`
	Name        string                  `json:"name" validate:"required"`
	EnglishName string                  `json:"english_name" validate:"required"`
	Data        PersonTranslationData   `json:"data"`
}

// PersonTranslationData represents person translation data
type PersonTranslationData struct {
	Biography string `json:"biography"`
}

// PopularPeople represents popular people
type PopularPeople = SearchResponse[Person]

// PersonTaggedImages represents tagged images for a person
type PersonTaggedImages struct {
	ID      int                   `json:"id" validate:"required"`
	Results []PersonTaggedImage   `json:"results"`
	Page    int                   `json:"page"`
	TotalPages   int              `json:"total_pages"`
	TotalResults int              `json:"total_results"`
}

// PersonTaggedImage represents a tagged image of a person
type PersonTaggedImage struct {
	AspectRatio float64       `json:"aspect_ratio"`
	FilePath    string        `json:"file_path" validate:"required"`
	Height      int           `json:"height" validate:"required"`
	ID          string        `json:"id" validate:"required"`
	ISO6391     *string       `json:"iso_639_1"`
	VoteAverage float64       `json:"vote_average"`
	VoteCount   int           `json:"vote_count"`
	Width       int           `json:"width" validate:"required"`
	ImageType   string        `json:"image_type" validate:"required"`
	Media       TaggedMedia   `json:"media"`
	MediaType   string        `json:"media_type" validate:"required"`
}

// TaggedMedia represents the media (movie/TV) associated with a tagged image
type TaggedMedia struct {
	Adult            *bool    `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id" validate:"required"`
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    *string  `json:"original_title,omitempty"` // For movies
	OriginalName     *string  `json:"original_name,omitempty"`  // For TV shows
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	ReleaseDate      *string  `json:"release_date,omitempty"`   // For movies
	FirstAirDate     *string  `json:"first_air_date,omitempty"` // For TV shows
	Title            *string  `json:"title,omitempty"`          // For movies
	Name             *string  `json:"name,omitempty"`           // For TV shows
	Video            *bool    `json:"video,omitempty"`          // For movies
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	OriginCountry    []string `json:"origin_country,omitempty"` // For TV shows
}