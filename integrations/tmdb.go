package integrations

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/qahta0/movies/models"
	"gorm.io/gorm"
)

const (
	baseMovieEndpoint   = "api.themoviedb.org/3/movie"
	latestMovieEndpoint = "latest"
)

func FetchAndStoreLatestMovie(dbConnection *gorm.DB, accessTokenAuth string) error {
	url := fmt.Sprintf("https://%s/%s", baseMovieEndpoint, latestMovieEndpoint)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessTokenAuth))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	movie := &models.Movie{}
	if err := json.NewDecoder(res.Body).Decode(movie); err != nil {
		return nil
	}
	storeLatestMovie(dbConnection, movie)
	return nil
}

func storeLatestMovie(dbConnection *gorm.DB, movie *models.Movie) error {
	if err := dbConnection.Where("id = ?", movie.ID).First(movie).Error; err == nil {
		log.Printf("Movie already exists in DB (id: %d title: %s)", movie.ID, movie.Title)
		return nil
	}
	if err := dbConnection.Create(&movie).Error; err != nil {
		return err
	}
	return nil
}
