package server

import (
	"context"
	"log"
	"net"

	"github.com/qahta0/movies/grpc/proto"
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
