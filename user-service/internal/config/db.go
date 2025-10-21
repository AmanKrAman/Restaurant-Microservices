package config

import (
	"log"

	"github.com/amankraman/restaurant-microservices/user-service/internal/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(
		&models.User{},
		&models.UserToken{},
	); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	log.Println("Database migrated successfully")
}
