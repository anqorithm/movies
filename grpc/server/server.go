package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/qahta0/movies/grpc/proto"
	"github.com/qahta0/movies/models"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type server struct {
	proto.UnimplementedMoviesServiceServer
	DB *gorm.DB
}

func (s *server) GetLatestMovies(ctx context.Context, req *proto.LatestMoviesRequest) (*proto.LatestMoviesResponse, error) {
	var movies []*proto.Movie
	result := s.DB.Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}
	return &proto.LatestMoviesResponse{
		Movies: movies,
	}, nil
}

func (s *server) SearchMovies(ctx context.Context, req *proto.SearchMoviesRequest) (*proto.SearchMoviesResponse, error) {
	var movies []*proto.Movie
	searchQuery := "%" + req.Query + "%"
	if err := s.DB.Where("title LIKE ? OR original_title LIKE ? OR overview LIKE ?", searchQuery, searchQuery, searchQuery).Find(&movies).Error; err != nil {
		return nil, err
	}
	return &proto.SearchMoviesResponse{Movies: movies}, nil
}

func (s *server) UpdateFavourites(ctx context.Context, req *proto.UpdateFavouritesRequest) (*proto.UpdateFavouritesResponse, error) {
	var favourite models.UserMovieFavourites
	var user models.User
	var movie models.Movie
	fmt.Println(req.UserId, req.MovieId)
	if err := s.DB.First(&user, "id = ?", req.UserId).Error; err != nil {
		return nil, fmt.Errorf("user with ID %d does not exist", req.UserId)
	}
	if err := s.DB.First(&movie, "id = ?", req.MovieId).Error; err != nil {
		return nil, fmt.Errorf("movie with ID %v does not exist", req.MovieId)
	}
	switch req.Action {
	case proto.FavouriteAction_ADD:
		if err := s.DB.First(&favourite, "user_id = ? AND movie_id = ?", req.UserId, req.MovieId).Error; err == nil {
			return &proto.UpdateFavouritesResponse{Message: "Movie is already in favourites!"}, nil
		}
		favourite = models.UserMovieFavourites{
			UserID:  int32(req.UserId),
			MovieID: int32(req.MovieId),
		}
		if err := s.DB.Create(&favourite).Error; err != nil {
			return nil, err
		}
		return &proto.UpdateFavouritesResponse{Message: "Movie added to favourites successfully!"}, nil
	case proto.FavouriteAction_REMOVE:
		if err := s.DB.Where("user_id = ? AND movie_id = ?", req.UserId, req.MovieId).Delete(&favourite).Error; err != nil {
			return nil, err
		}
		return &proto.UpdateFavouritesResponse{Message: "Movie removed from favourites successfully"}, nil
	default:
		return nil, fmt.Errorf("unknown action: %v", req.Action)
	}
}

func StartGRPCServer(dbConnection *gorm.DB) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterMoviesServiceServer(grpcServer, &server{DB: dbConnection})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
