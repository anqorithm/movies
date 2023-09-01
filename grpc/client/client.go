package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/qahta0/movies/grpc/proto"
	"github.com/qahta0/movies/helpers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getLatestMovies(client proto.MoviesServiceClient) {
	res, err := client.GetLatestMovies(context.Background(), &proto.LatestMoviesRequest{})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	for _, movie := range res.GetMovies() {
		movieDTO := helpers.MovieProtoToDTO(movie)
		movieJSON, err := json.Marshal(movieDTO)
		if err != nil {
			log.Fatalf("err: %v", err)
		}
		var out bytes.Buffer
		json.Indent(&out, movieJSON, "", " ")
		out.WriteTo(os.Stdout)
		println(",")
	}
}

func searchMovies(client proto.MoviesServiceClient, query string) {
	res, err := client.SearchMovies(context.Background(), &proto.SearchMoviesRequest{Query: query})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	for _, movie := range res.Movies {
		movieDTO := helpers.MovieProtoToDTO(movie)
		movieJSON, err := json.Marshal(movieDTO)
		if err != nil {
			log.Fatalf("err: %v", err)
		}
		var out bytes.Buffer
		json.Indent(&out, movieJSON, "", " ")
		out.WriteTo(os.Stdout)
		println(",")
	}
}

func updateFavourites(client proto.MoviesServiceClient, action proto.FavouriteAction, user_id uint32, movie_id uint32) {
	res, err := client.UpdateFavourites(context.Background(), &proto.UpdateFavouritesRequest{Action: action, UserId: user_id, MovieId: movie_id})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Printf("%v", res.Message)
}

func getMovieDetials(client proto.MoviesServiceClient, movie_id uint32) {
	res, err := client.GetMovieDetials(context.Background(), &proto.MovieDetialsRequest{MovieId: movie_id})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	movieDTO := helpers.MoveDetailsProtoToDTO(res.MovieDetails)
	movieJSON, err := json.Marshal(movieDTO)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	var out bytes.Buffer
	json.Indent(&out, movieJSON, "", " ")
	out.WriteTo(os.Stdout)
	println(",")
}

func main() {
	connection, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer connection.Close()
	client := proto.NewMoviesServiceClient(connection)
	var choice string
	fmt.Println("1: Get latest movies")
	fmt.Println("2: Search movies")
	fmt.Println("3: Update Favorites")
	fmt.Println("4: Get movie details")
	fmt.Print("Enter a number (1-4): ")
	fmt.Scanln(&choice)
	switch strings.TrimSpace(choice) {
	case "1":
		getLatestMovies(client)
	case "2":
		var searchQuery string
		fmt.Print("Enter search query: ")
		fmt.Scanln(&searchQuery)
		searchMovies(client, searchQuery)
	case "3":
		var userId uint32
		var movieId uint32
		var action proto.FavouriteAction
		fmt.Print("Enter UserID: ")
		fmt.Scanln(&userId)
		fmt.Print("Enter MovieID: ")
		fmt.Scanln(&movieId)
		fmt.Print("Enter Action: ")
		fmt.Scanln(&action)
		updateFavourites(client, action, userId, movieId)
	case "4":
		var movieId uint32
		fmt.Print("Enter MovieID: ")
		fmt.Scanln(&movieId)
		getMovieDetials(client, movieId)
	default:
		log.Fatalf("Invalid choice: %s", choice)
	}
}
