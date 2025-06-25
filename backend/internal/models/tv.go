// Package models provides TV show-related data structures for TMDb API integration.
package models

// TVShow represents a TV show from TMDb API (used in search results and lists)
type TVShow struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id" validate:"required"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language" validate:"required"`
	OriginalName     string   `json:"original_name" validate:"required"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	FirstAirDate     *string  `json:"first_air_date"` // Format: YYYY-MM-DD
	Name             string   `json:"name" validate:"required"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

// TVShowDetails represents detailed TV show information from TMDb API
type TVShowDetails struct {
	Adult               bool                 `json:"adult"`
	BackdropPath        *string              `json:"backdrop_path"`
	CreatedBy           []TVCreator          `json:"created_by"`
	EpisodeRunTime      []int                `json:"episode_run_time"`
	FirstAirDate        *string              `json:"first_air_date"` // Format: YYYY-MM-DD
	Genres              []Genre              `json:"genres"`
	Homepage            *string              `json:"homepage"`
	ID                  int                  `json:"id" validate:"required"`
	InProduction        bool                 `json:"in_production"`
	Languages           []string             `json:"languages"`
	LastAirDate         *string              `json:"last_air_date"` // Format: YYYY-MM-DD
	LastEpisodeToAir    *Episode             `json:"last_episode_to_air"`
	Name                string               `json:"name" validate:"required"`
	NextEpisodeToAir    *Episode             `json:"next_episode_to_air"`
	Networks            []Network            `json:"networks"`
	NumberOfEpisodes    int                  `json:"number_of_episodes"`
	NumberOfSeasons     int                  `json:"number_of_seasons"`
	OriginCountry       []string             `json:"origin_country"`
	OriginalLanguage    string               `json:"original_language" validate:"required"`
	OriginalName        string               `json:"original_name" validate:"required"`
	Overview            *string              `json:"overview"`
	Popularity          float64              `json:"popularity"`
	PosterPath          *string              `json:"poster_path"`
	ProductionCompanies []ProductionCompany  `json:"production_companies"`
	ProductionCountries []ProductionCountry  `json:"production_countries"`
	Seasons             []Season             `json:"seasons"`
	SpokenLanguages     []SpokenLanguage     `json:"spoken_languages"`
	Status              string               `json:"status"`
	Tagline             *string              `json:"tagline"`
	Type                string               `json:"type"`
	VoteAverage         float64              `json:"vote_average"`
	VoteCount           int                  `json:"vote_count"`
}

// TVCreator represents a TV show creator
type TVCreator struct {
	ID          int     `json:"id" validate:"required"`
	CreditID    string  `json:"credit_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Gender      *int    `json:"gender"`
	ProfilePath *string `json:"profile_path"`
}

// Season represents a TV season
type Season struct {
	AirDate      *string `json:"air_date"` // Format: YYYY-MM-DD
	EpisodeCount int     `json:"episode_count"`
	ID           int     `json:"id" validate:"required"`
	Name         string  `json:"name" validate:"required"`
	Overview     string  `json:"overview"`
	PosterPath   *string `json:"poster_path"`
	SeasonNumber int     `json:"season_number" validate:"required"`
	VoteAverage  float64 `json:"vote_average"`
}

// Episode represents a TV episode
type Episode struct {
	ID             int      `json:"id" validate:"required"`
	Name           string   `json:"name" validate:"required"`
	Overview       string   `json:"overview"`
	VoteAverage    float64  `json:"vote_average"`
	VoteCount      int      `json:"vote_count"`
	AirDate        *string  `json:"air_date"` // Format: YYYY-MM-DD
	EpisodeNumber  int      `json:"episode_number" validate:"required"`
	ProductionCode *string  `json:"production_code"`
	Runtime        *int     `json:"runtime"`
	SeasonNumber   int      `json:"season_number" validate:"required"`
	ShowID         int      `json:"show_id"`
	StillPath      *string  `json:"still_path"`
}

// TVSearchResponse represents a search response for TV shows
type TVSearchResponse = SearchResponse[TVShow]

// TVCredits represents cast and crew for a TV show
type TVCredits struct {
	ID   int                `json:"id" validate:"required"`
	Cast []TVCastMember     `json:"cast"`
	Crew []TVCrewMember     `json:"crew"`
}

// TVCastMember represents a cast member in TV show credits
type TVCastMember struct {
	Adult              bool     `json:"adult"`
	Gender             *int     `json:"gender"`
	ID                 int      `json:"id" validate:"required"`
	KnownForDepartment string   `json:"known_for_department"`
	Name               string   `json:"name" validate:"required"`
	OriginalName       string   `json:"original_name"`
	Popularity         float64  `json:"popularity"`
	ProfilePath        *string  `json:"profile_path"`
	Character          string   `json:"character"`
	CreditID           string   `json:"credit_id" validate:"required"`
	Order              int      `json:"order"`
}

// TVCrewMember represents a crew member in TV show credits
type TVCrewMember struct {
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

// SeasonDetails represents detailed season information
type SeasonDetails struct {
	AirDate      *string   `json:"air_date"` // Format: YYYY-MM-DD
	Episodes     []Episode `json:"episodes"`
	Name         string    `json:"name" validate:"required"`
	Overview     string    `json:"overview"`
	ID           int       `json:"id" validate:"required"`
	PosterPath   *string   `json:"poster_path"`
	SeasonNumber int       `json:"season_number" validate:"required"`
	VoteAverage  float64   `json:"vote_average"`
}

// EpisodeDetails represents detailed episode information
type EpisodeDetails struct {
	AirDate        *string       `json:"air_date"` // Format: YYYY-MM-DD
	Crew           []TVCrewMember `json:"crew"`
	EpisodeNumber  int           `json:"episode_number" validate:"required"`
	GuestStars     []TVCastMember `json:"guest_stars"`
	Name           string        `json:"name" validate:"required"`
	Overview       string        `json:"overview"`
	ID             int           `json:"id" validate:"required"`
	ProductionCode *string       `json:"production_code"`
	Runtime        *int          `json:"runtime"`
	SeasonNumber   int           `json:"season_number" validate:"required"`
	StillPath      *string       `json:"still_path"`
	VoteAverage    float64       `json:"vote_average"`
	VoteCount      int           `json:"vote_count"`
}

// TVImages represents images for a TV show
type TVImages struct {
	ID        int     `json:"id" validate:"required"`
	Backdrops []Image `json:"backdrops"`
	Logos     []Image `json:"logos"`
	Posters   []Image `json:"posters"`
}

// TVVideos represents videos for a TV show
type TVVideos struct {
	ID      int     `json:"id" validate:"required"`
	Results []Video `json:"results"`
}

// TVRecommendations represents recommended TV shows
type TVRecommendations = SearchResponse[TVShow]

// TVSimilar represents similar TV shows
type TVSimilar = SearchResponse[TVShow]

// TVWatchProviders represents watch providers for a TV show
type TVWatchProviders struct {
	ID      int                               `json:"id" validate:"required"`
	Results map[string]CountryWatchProviders `json:"results"`
}

// TVTranslations represents translations for a TV show
type TVTranslations struct {
	ID           int               `json:"id" validate:"required"`
	Translations []TVTranslation   `json:"translations"`
}

// TVTranslation represents a single TV translation
type TVTranslation struct {
	ISO31661    string              `json:"iso_3166_1" validate:"required"`
	ISO6391     string              `json:"iso_639_1" validate:"required"`
	Name        string              `json:"name" validate:"required"`
	EnglishName string              `json:"english_name" validate:"required"`
	Data        TVTranslationData   `json:"data"`
}

// TVTranslationData represents TV translation data
type TVTranslationData struct {
	Homepage *string `json:"homepage"`
	Overview *string `json:"overview"`
	Tagline  *string `json:"tagline"`
	Name     *string `json:"name"`
}

// TVAlternativeTitles represents alternative titles for a TV show
type TVAlternativeTitles struct {
	ID      int                    `json:"id" validate:"required"`
	Results []TVAlternativeTitle   `json:"results"`
}

// TVAlternativeTitle represents a single alternative title for TV
type TVAlternativeTitle struct {
	ISO31661 string `json:"iso_3166_1" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Type     string `json:"type"`
}

// AiringTodayTVShows represents TV shows airing today
type AiringTodayTVShows = SearchResponse[TVShow]

// OnTheAirTVShows represents TV shows currently on the air
type OnTheAirTVShows = SearchResponse[TVShow]

// PopularTVShows represents popular TV shows
type PopularTVShows = SearchResponse[TVShow]

// TopRatedTVShows represents top rated TV shows
type TopRatedTVShows = SearchResponse[TVShow]

// DiscoverTVShows represents discover TV shows response
type DiscoverTVShows = SearchResponse[TVShow]

// TVSeasonCredits represents credits for a specific season
type TVSeasonCredits struct {
	Cast []TVCastMember `json:"cast"`
	Crew []TVCrewMember `json:"crew"`
	ID   int            `json:"id" validate:"required"`
}

// TVEpisodeCredits represents credits for a specific episode
type TVEpisodeCredits struct {
	Cast       []TVCastMember `json:"cast"`
	Crew       []TVCrewMember `json:"crew"`
	GuestStars []TVCastMember `json:"guest_stars"`
	ID         int           `json:"id" validate:"required"`
}