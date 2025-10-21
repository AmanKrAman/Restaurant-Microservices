package config

import (
	"log"

	"github.com/amankraman/restaurant-microservices/order-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB connects to PostgreSQL using cfg.DBUrl
func ConnectDB() *gorm.DB {
	cfg := LoadConfig()
	db, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	log.Println("Connected to DB")
	return db
}

// RunMigrations runs GORM auto migrations
func RunMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Order{}, &models.OrderItem{}); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	log.Println("Database migrated successfully")
}
