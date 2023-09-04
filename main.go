package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/qahta0/movies/db"
	"github.com/qahta0/movies/grpc/server"
	"github.com/qahta0/movies/helpers"
	"github.com/qahta0/movies/integrations"
	"github.com/robfig/cron/v3"
)

func main() {
	helpers.LoadEnv()
	dbConnection := db.Connect()
	accessTokenAuth := os.Getenv("ACCESS_TOKEN_AUTH")
	isRunSeeders, err := strconv.ParseBool(os.Getenv("RUN_SEEDERS"))
	if err != nil {
		log.Fatal(err)
	}
	isRunFactories, err := strconv.ParseBool(os.Getenv("RUN_FACTORIES"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isRunFactories, isRunSeeders)
	db.FlushAndMigrate(dbConnection, false)
	db.RunSeeders(dbConnection, isRunSeeders)
	db.RunFactories(dbConnection, isRunFactories, 10)
	c := cron.New(cron.WithSeconds())
	c.AddFunc("@every 3s", func() { integrations.FetchAndStoreLatestMovie(dbConnection, &accessTokenAuth) })
	c.Start()
	server.StartGRPCServer(dbConnection)
	select {}
}
