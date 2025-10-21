package main

import (
	"fmt"
	"log"

	"github.com/amankraman/restaurant-microservices/payment-service/internal/config"
	"github.com/amankraman/restaurant-microservices/payment-service/internal/handlers"
	"github.com/amankraman/restaurant-microservices/payment-service/internal/repository"
	"github.com/amankraman/restaurant-microservices/payment-service/internal/service"

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

	// Initialize repo, service, handler
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepo)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	r := gin.Default()

	// Routes
	r.POST("/payments", paymentHandler.CreatePayment)
	r.PATCH("/payments/:id/status", paymentHandler.UpdatePaymentStatus)
	r.GET("/payments/:id", paymentHandler.GetPaymentByID)
	r.GET("/payments/order/:order_id", paymentHandler.GetPaymentByOrderID)

	log.Printf("payment-service running on port %s", cfg.Port)
	r.Run(fmt.Sprintf(":%s", cfg.Port))
}
