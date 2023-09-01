package dtos

type MovieDTO struct {
	ID                  int32   `json:"id"`
	Adult               bool    `json:"adult"`
	BackdropPath        string  `json:"backdrop_path"`
	BelongsToCollection string  `json:"belongs_to_collection"`
	Budget              int32   `json:"budget"`
	Homepage            string  `json:"homepage"`
	ImdbID              string  `json:"imdb_id"`
	OriginalLanguage    string  `json:"original_language"`
	OriginalTitle       string  `json:"original_title"`
	Overview            string  `json:"overview"`
	Popularity          float32 `json:"popularity"`
	PosterPath          string  `json:"poster_path"`
	ReleaseDate         string  `json:"release_date"`
	Revenue             int32   `json:"revenue"`
	Runtime             int32   `json:"runtime"`
	Status              string  `json:"status"`
	Tagline             string  `json:"tagline"`
	Title               string  `json:"title"`
	Video               bool    `json:"video"`
	VoteAverage         float32 `json:"vote_average"`
	VoteCount           int32   `json:"vote_count"`
}

type MoveDetailsDTO struct {
	ID         int32  `json:"id"`
	Title      string `json:"name"`
	PosterPath string `json:"poster"`
	Overview   string `json:"description"`
}
