package main

import (
	"fmt"
	"log"

	"github.com/amankraman/restaurant-microservices/menu-service/internal/config"
	"github.com/amankraman/restaurant-microservices/menu-service/internal/handlers"
	"github.com/amankraman/restaurant-microservices/menu-service/internal/repository"
	"github.com/amankraman/restaurant-microservices/menu-service/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	db, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Run migrations
	config.RunMigrations(db)

	repo := repository.NewMenuRepository(db)
	menuService := service.NewMenuService(repo)
	menuHandler := handlers.NewMenuHandler(menuService)

	r := gin.Default()

	// Restaurant routes
	r.POST("/restaurants", menuHandler.CreateRestaurant)
	r.GET("/restaurants", menuHandler.GetRestaurants)
	r.GET("/restaurants/:id", menuHandler.GetRestaurantByID)

	// Menu routes
	r.POST("/menus", menuHandler.CreateMenu)

	// Dish routes
	r.POST("/dishes", menuHandler.CreateDish)

	log.Printf("menu-service running on port %s", cfg.Port)
	r.Run(fmt.Sprintf(":%s", cfg.Port))
}
