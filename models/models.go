package models

type Movie struct {
	ID                  uint                `gorm:"primaryKey" json:"id"`
	Adult               bool                `json:"adult"`
	BackdropPath        string              `json:"backdrop_path"`
	BelongsToCollection string              `json:"belongs_to_collection"`
	Budget              int                 `json:"budget"`
	Homepage            string              `json:"homepage"`
	ImdbID              string              `gorm:"unique" json:"imdb_id"`
	OriginalLanguage    string              `json:"original_language"`
	OriginalTitle       string              `json:"original_title"`
	Overview            string              `json:"overview"`
	Popularity          float32             `json:"popularity"`
	PosterPath          string              `json:"poster_path"`
	ReleaseDate         string              `json:"release_date"`
	Revenue             int                 `json:"revenue"`
	Runtime             int                 `json:"runtime"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
	Title               string              `json:"title"`
	Video               bool                `json:"video"`
	VoteAverage         float32             `json:"vote_average"`
	VoteCount           int                 `json:"vote_count"`
	Genres              []Genre             `gorm:"many2many:movie_genres;" json:"genres"`
	ProductionCompanies []ProductionCompany `gorm:"many2many:movie_production_companies;" json:"production_companies"`
	ProductionCountries []ProductionCountry `gorm:"many2many:movie_production_countries;" json:"production_countries"`
	SpokenLanguages     []SpokenLanguage    `gorm:"many2many:movie_spoken_languages;" json:"spoken_languages"`
}

type Genre struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type ProductionCompany struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type ProductionCountry struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type SpokenLanguage struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
