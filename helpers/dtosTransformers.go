package helpers

import (
	"github.com/qahta0/movies/dtos"
	"github.com/qahta0/movies/grpc/proto"
)

func MovieProtoToDTO(movie *proto.Movie) *dtos.MovieDTO {
	return &dtos.MovieDTO{
		ID:                  movie.GetId(),
		Adult:               movie.GetAdult(),
		BackdropPath:        movie.GetBackdropPath(),
		BelongsToCollection: movie.GetBelongsToCollection(),
		Budget:              movie.GetBudget(),
		Homepage:            movie.GetHomepage(),
		ImdbID:              movie.GetImdbId(),
		OriginalLanguage:    movie.GetOriginalLanguage(),
		OriginalTitle:       movie.GetOriginalTitle(),
		Overview:            movie.GetOverview(),
		Popularity:          movie.GetPopularity(),
		PosterPath:          movie.GetPosterPath(),
		ReleaseDate:         movie.GetReleaseDate(),
		Revenue:             movie.GetRevenue(),
		Runtime:             movie.GetRuntime(),
		Status:              movie.GetStatus(),
		Tagline:             movie.GetTagline(),
		Title:               movie.GetTitle(),
		Video:               movie.GetVideo(),
		VoteAverage:         movie.GetVoteAverage(),
		VoteCount:           movie.GetVoteCount(),
	}
}

func MoveDetailsProtoToDTO(movieDetails *proto.MovieDetails) *dtos.MoveDetailsDTO {
	return &dtos.MoveDetailsDTO{
		ID:         movieDetails.GetId(),
		Title:      movieDetails.GetTitle(),
		Overview:   movieDetails.GetOverview(),
		PosterPath: movieDetails.GetPosterPath(),
	}
}
