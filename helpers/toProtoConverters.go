package helpers

import (
	"github.com/qahta0/movies/grpc/proto"
	"github.com/qahta0/movies/models"
)

func ConvertToProtoGenres(genres *[]models.Genre) []*proto.Genre {
	protoGenres := make([]*proto.Genre, len(*genres))
	for i, genre := range *genres {
		protoGenres[i] = &proto.Genre{
			Id:   int32(genre.ID),
			Name: genre.Name,
		}
	}
	return protoGenres
}
