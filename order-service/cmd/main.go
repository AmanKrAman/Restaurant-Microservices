package main

import (
	"fmt"
	"log"

	"github.com/amankraman/restaurant-microservices/order-service/internal/config"
	"github.com/amankraman/restaurant-microservices/order-service/internal/handlers"
	"github.com/amankraman/restaurant-microservices/order-service/internal/repository"
	"github.com/amankraman/restaurant-microservices/order-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db := config.ConnectDB()
	config.RunMigrations(db)

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	r := gin.Default()

	r.POST("/orders", orderHandler.CreateOrder)
	r.PATCH("/orders/:id/status", orderHandler.UpdateOrderStatus)
	r.GET("/orders", orderHandler.GetOrders)

	log.Printf("order-service running on port %s", cfg.Port)
	r.Run(fmt.Sprintf(":%s", cfg.Port))
}
