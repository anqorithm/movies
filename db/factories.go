package db

import (
	"log"

	"github.com/go-faker/faker/v4"
	"github.com/qahta0/movies/models"
	"gorm.io/gorm"
)

func RunFactories(db *gorm.DB, isAllowedToRunFactories bool, n int) {
	if isAllowedToRunFactories {
		generateUsers(db, n)
	}
}

func generateUsers(db *gorm.DB, n int) {
	for i := 0; i < n; i++ {
		user := &models.User{
			Password: "qahta0",
		}
		err := faker.FakeData(user)
		if err != nil {
			log.Fatalf("err: %v", err)
		}
		if err := db.Create(user).Error; err != nil {
			log.Fatalf("err: %v", err)
		}
	}
}
