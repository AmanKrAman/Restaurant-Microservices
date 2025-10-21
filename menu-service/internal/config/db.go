package config

import (
	"log"

	"github.com/amankraman/restaurant-microservices/menu-service/internal/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(
		&models.Restaurant{},
		&models.Menu{},
		&models.Dish{},
	); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	log.Println("Database migrated successfully")
}
