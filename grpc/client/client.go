package main

import (
	"context"
	"fmt"
	"log"

	"github.com/qahta0/movies/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getLatestMovies(client proto.MoviesServiceClient) {
	res, err := client.GetLatestMovies(context.Background(), &proto.LatestMoviesRequest{})
	if err != nil {
		log.Fatalf("Could not fetch latest movies: %v", err)
	}
	for _, movie := range res.GetMovies() {
		fmt.Println("ID:", movie.GetId())
		fmt.Println("Adult:", movie.GetAdult())
		fmt.Println("Backdrop Path:", movie.GetBackdropPath())
		fmt.Println("Belongs to Collection:", movie.GetBelongsToCollection())
		fmt.Println("Budget:", movie.GetBudget())
		fmt.Println("Homepage:", movie.GetHomepage())
		fmt.Println("IMDB ID:", movie.GetImdbId())
		fmt.Println("Original Language:", movie.GetOriginalLanguage())
		fmt.Println("Original Title:", movie.GetOriginalTitle())
		fmt.Println("Overview:", movie.GetOverview())
		fmt.Println("Popularity:", movie.GetPopularity())
		fmt.Println("Poster Path:", movie.GetPosterPath())
		fmt.Println("Release Date:", movie.GetReleaseDate())
		fmt.Println("Revenue:", movie.GetRevenue())
		fmt.Println("Runtime:", movie.GetRuntime())
		fmt.Println("Status:", movie.GetStatus())
		fmt.Println("Tagline:", movie.GetTagline())
		fmt.Println("Title:", movie.GetTitle())
		fmt.Println("Video:", movie.GetVideo())
		fmt.Println("Vote Average:", movie.GetVoteAverage())
		fmt.Println("Vote Count:", movie.GetVoteCount())
		fmt.Println("########################################################")
	}
}

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer connection.Close()
	client := proto.NewMoviesServiceClient(connection)
	getLatestMovies(client)
}
