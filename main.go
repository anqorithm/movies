package main

import (
	"github.com/qahta0/movies/db"
	"github.com/qahta0/movies/helpers"
)

func main() {
	helpers.LoadEnv()
	dbConnection := db.Connect()
	db.FlushAndMigrate(dbConnection, false)
	select {}
}
