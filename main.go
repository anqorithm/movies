package main

import (
	"os"

	"github.com/qahta0/movies/db"
	"github.com/qahta0/movies/helpers"
	"github.com/qahta0/movies/integrations"
	"github.com/robfig/cron/v3"
)

func main() {
	helpers.LoadEnv()
	dbConnection := db.Connect()
	db.FlushAndMigrate(dbConnection, false)
	accessTokenAuth := os.Getenv("ACCESS_TOKEN_AUTH")
	c := cron.New(cron.WithSeconds())
	c.AddFunc("@every 3s", func() { integrations.FetchAndStoreLatestMovie(dbConnection, accessTokenAuth) })
	c.Start()
	select {}
}
