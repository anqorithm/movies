package db

import (
	"log"

	"github.com/qahta0/movies/models"
	"gorm.io/gorm"
)

func FlushAndMigrate(db *gorm.DB, flushDatabase bool) {
	if flushDatabase {
		if err := db.Migrator().DropTable(&models.Movie{}, &models.Genre{}, &models.ProductionCompany{}, &models.ProductionCountry{}, &models.SpokenLanguage{}); err != nil {
			log.Fatalf("Could not drop tables: %v", err)
		}
	}
	if err := db.AutoMigrate(&models.Movie{}, &models.Genre{}, &models.ProductionCompany{}, &models.ProductionCountry{}, &models.SpokenLanguage{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}
