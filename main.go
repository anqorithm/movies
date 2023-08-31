package main

import (
	"os"

	"github.com/qahta0/movies/db"
	"github.com/qahta0/movies/helpers"
	"github.com/qahta0/movies/integrations"
)

func main() {
	helpers.LoadEnv()
	dbConnection := db.Connect()
	db.FlushAndMigrate(dbConnection, false)
	accessTokenAuth := os.Getenv("ACCESS_TOKEN_AUTH")
	latestMovie, err := integrations.FetchLatestMovie(accessTokenAuth)
	if err != nil {
		return
	}
	integrations.StoreLatestMovie(dbConnection, latestMovie)
	select {}
}
