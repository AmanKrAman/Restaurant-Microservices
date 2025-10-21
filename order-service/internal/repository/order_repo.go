package repository

import (
	"github.com/amankraman/restaurant-microservices/order-service/internal/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) UpdateOrder(order *models.Order) error {
	return r.db.Save(order).Error
}

func (r *OrderRepository) GetOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Items").Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) GetOrderByID(id uint) (models.Order, error) {
	var order models.Order
	err := r.db.Preload("Items").First(&order, id).Error
	return order, err
}
