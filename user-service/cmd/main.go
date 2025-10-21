package main

import (
	"fmt"
	"log"
	"restaurant-microservices/common/utils"

	"github.com/amankraman/restaurant-microservices/user-service/internal/config"
	"github.com/amankraman/restaurant-microservices/user-service/internal/handlers"
	"github.com/amankraman/restaurant-microservices/user-service/internal/repository"
	"github.com/amankraman/restaurant-microservices/user-service/internal/service"
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

	config.RunMigrations(db)

	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		rb := &utils.ResponseBuilder{}
		response, err := rb.GenerateResponse(true, "Service is up and running", nil, nil)
		if err != nil {
			c.JSON(500, gin.H{"status": "error", "message": "Internal Server Error"})
			return
		}

		c.Data(200, "application/json", []byte(response))
	})

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	r.GET("/users", userHandler.GetUsers)
	r.POST("/logout", userHandler.Logout)

	log.Printf("user-service running on port %s", cfg.Port)
	r.Run(fmt.Sprintf(":%s", cfg.Port))
}
