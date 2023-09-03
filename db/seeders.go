package db

import (
	"log"

	"github.com/qahta0/movies/models"
	"gorm.io/gorm"
)

func RunSeeders(db *gorm.DB, isAllowedToRunSeeders bool) {
	if isAllowedToRunSeeders {
		seedUsers(db)
	}
}

func seedUsers(db *gorm.DB) {
	user := models.User{
		ID:       1,
		Email:    "abdullah.n.alqahtani@hotmail.com",
		Username: "qahta0",
		Password: "qahta0",
		Name:     "Abdullah Alqahtani",
	}
	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("err: %v", err)
	}
}
