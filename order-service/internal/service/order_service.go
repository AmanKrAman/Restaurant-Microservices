package service

import (
	"github.com/amankraman/restaurant-microservices/order-service/internal/models"
	"github.com/amankraman/restaurant-microservices/order-service/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	// calculate total price
	var total float64
	for i := range order.Items {
		total += order.Items[i].Price * float64(order.Items[i].Qty)
	}
	order.TotalPrice = total

	order.Status = "pending"
	return s.repo.CreateOrder(order)
}

func (s *OrderService) UpdateOrderStatus(id uint, status string) error {
	order, err := s.repo.GetOrderByID(id)
	if err != nil {
		return err
	}
	order.Status = status
	return s.repo.UpdateOrder(&order)
}

func (s *OrderService) GetOrders() ([]models.Order, error) {
	return s.repo.GetOrders()
}

func (s *OrderService) GetOrderByID(id uint) (models.Order, error) {
	return s.repo.GetOrderByID(id)
}
