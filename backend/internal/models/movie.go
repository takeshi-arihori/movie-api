// Package models provides movie-related data structures for TMDb API integration.
package models

// Movie represents a movie from TMDb API (used in search results and lists)
type Movie struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id" validate:"required"`
	OriginalLanguage string   `json:"original_language" validate:"required"`
	OriginalTitle    string   `json:"original_title" validate:"required"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	ReleaseDate      *string  `json:"release_date"` // Format: YYYY-MM-DD
	Title            string   `json:"title" validate:"required"`
	Video            bool     `json:"video"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

// MovieDetails represents detailed movie information from TMDb API
type MovieDetails struct {
	Adult               bool                 `json:"adult"`
	BackdropPath        *string              `json:"backdrop_path"`
	BelongsToCollection *Collection          `json:"belongs_to_collection"`
	Budget              int64                `json:"budget"`
	Genres              []Genre              `json:"genres"`
	Homepage            *string              `json:"homepage"`
	ID                  int                  `json:"id" validate:"required"`
	IMDbID              *string              `json:"imdb_id"`
	OriginalLanguage    string               `json:"original_language" validate:"required"`
	OriginalTitle       string               `json:"original_title" validate:"required"`
	Overview            *string              `json:"overview"`
	Popularity          float64              `json:"popularity"`
	PosterPath          *string              `json:"poster_path"`
	ProductionCompanies []ProductionCompany  `json:"production_companies"`
	ProductionCountries []ProductionCountry  `json:"production_countries"`
	ReleaseDate         *string              `json:"release_date"` // Format: YYYY-MM-DD
	Revenue             int64                `json:"revenue"`
	Runtime             *int                 `json:"runtime"` // In minutes
	SpokenLanguages     []SpokenLanguage     `json:"spoken_languages"`
	Status              string               `json:"status"`
	Tagline             *string              `json:"tagline"`
	Title               string               `json:"title" validate:"required"`
	Video               bool                 `json:"video"`
	VoteAverage         float64              `json:"vote_average"`
	VoteCount           int                  `json:"vote_count"`
}

// MovieSearchResponse represents a search response for movies
type MovieSearchResponse = SearchResponse[Movie]

// MovieCredits represents cast and crew for a movie
type MovieCredits struct {
	ID   int                `json:"id" validate:"required"`
	Cast []CastMember       `json:"cast"`
	Crew []CrewMember       `json:"crew"`
}

// CastMember represents a cast member in movie credits
type CastMember struct {
	Adult              bool     `json:"adult"`
	Gender             *int     `json:"gender"`
	ID                 int      `json:"id" validate:"required"`
	KnownForDepartment string   `json:"known_for_department"`
	Name               string   `json:"name" validate:"required"`
	OriginalName       string   `json:"original_name"`
	Popularity         float64  `json:"popularity"`
	ProfilePath        *string  `json:"profile_path"`
	CastID             int      `json:"cast_id"`
	Character          string   `json:"character"`
	CreditID           string   `json:"credit_id" validate:"required"`
	Order              int      `json:"order"`
}

// CrewMember represents a crew member in movie credits
type CrewMember struct {
	Adult              bool     `json:"adult"`
	Gender             *int     `json:"gender"`
	ID                 int      `json:"id" validate:"required"`
	KnownForDepartment string   `json:"known_for_department"`
	Name               string   `json:"name" validate:"required"`
	OriginalName       string   `json:"original_name"`
	Popularity         float64  `json:"popularity"`
	ProfilePath        *string  `json:"profile_path"`
	CreditID           string   `json:"credit_id" validate:"required"`
	Department         string   `json:"department" validate:"required"`
	Job                string   `json:"job" validate:"required"`
}

// MovieImages represents images for a movie
type MovieImages struct {
	ID        int     `json:"id" validate:"required"`
	Backdrops []Image `json:"backdrops"`
	Logos     []Image `json:"logos"`
	Posters   []Image `json:"posters"`
}

// MovieVideos represents videos for a movie
type MovieVideos struct {
	ID      int     `json:"id" validate:"required"`
	Results []Video `json:"results"`
}

// MovieRecommendations represents recommended movies
type MovieRecommendations = SearchResponse[Movie]

// MovieSimilar represents similar movies
type MovieSimilar = SearchResponse[Movie]

// MovieWatchProviders represents watch providers for a movie
type MovieWatchProviders struct {
	ID      int                                `json:"id" validate:"required"`
	Results map[string]CountryWatchProviders  `json:"results"`
}

// CountryWatchProviders represents watch providers for a specific country
type CountryWatchProviders struct {
	Link     string          `json:"link"`
	Flatrate []WatchProvider `json:"flatrate"`
	Rent     []WatchProvider `json:"rent"`
	Buy      []WatchProvider `json:"buy"`
}

// WatchProvider represents a single watch provider
type WatchProvider struct {
	DisplayPriority int    `json:"display_priority"`
	LogoPath        string `json:"logo_path" validate:"required"`
	ProviderID      int    `json:"provider_id" validate:"required"`
	ProviderName    string `json:"provider_name" validate:"required"`
}

// MovieTranslations represents translations for a movie
type MovieTranslations struct {
	ID           int                 `json:"id" validate:"required"`
	Translations []MovieTranslation  `json:"translations"`
}

// MovieTranslation represents a single translation
type MovieTranslation struct {
	ISO31661    string                  `json:"iso_3166_1" validate:"required"`
	ISO6391     string                  `json:"iso_639_1" validate:"required"`
	Name        string                  `json:"name" validate:"required"`
	EnglishName string                  `json:"english_name" validate:"required"`
	Data        MovieTranslationData    `json:"data"`
}

// MovieTranslationData represents translation data
type MovieTranslationData struct {
	Homepage *string `json:"homepage"`
	Overview *string `json:"overview"`
	Runtime  *int    `json:"runtime"`
	Tagline  *string `json:"tagline"`
	Title    *string `json:"title"`
}

// MovieAlternativeTitles represents alternative titles for a movie
type MovieAlternativeTitles struct {
	ID     int                     `json:"id" validate:"required"`
	Titles []MovieAlternativeTitle `json:"titles"`
}

// MovieAlternativeTitle represents a single alternative title
type MovieAlternativeTitle struct {
	ISO31661 string `json:"iso_3166_1" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Type     string `json:"type"`
}

// NowPlayingMovies represents currently playing movies with dates
type NowPlayingMovies struct {
	SearchResponse[Movie]
	Dates MovieDates `json:"dates"`
}

// MovieDates represents date range for now playing movies
type MovieDates struct {
	Maximum string `json:"maximum" validate:"required"` // Format: YYYY-MM-DD
	Minimum string `json:"minimum" validate:"required"` // Format: YYYY-MM-DD
}

// PopularMovies represents popular movies
type PopularMovies = SearchResponse[Movie]

// TopRatedMovies represents top rated movies
type TopRatedMovies = SearchResponse[Movie]

// UpcomingMovies represents upcoming movies with dates
type UpcomingMovies struct {
	SearchResponse[Movie]
	Dates MovieDates `json:"dates"`
}

// DiscoverMovies represents discover movies response
type DiscoverMovies = SearchResponse[Movie]