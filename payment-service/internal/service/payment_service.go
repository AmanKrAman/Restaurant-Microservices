package service

import (
	"errors"

	"github.com/amankraman/restaurant-microservices/payment-service/internal/models"
	"github.com/amankraman/restaurant-microservices/payment-service/internal/repository"
	"github.com/google/uuid"
)

type PaymentService struct {
	repo *repository.PaymentRepository
}

func NewPaymentService(repo *repository.PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

// Create a new payment for an order
func (s *PaymentService) CreatePayment(orderID uint, userID uint, amount float64, method string) (models.Payment, error) {
	if method == "" {
		return models.Payment{}, errors.New("payment method required")
	}

	// Generate a unique transaction ID
	transactionID := uuid.New().String()

	payment := models.Payment{
		OrderID:       orderID,
		UserID:        userID,
		Amount:        amount,
		PaymentMethod: method,
		Status:        "PENDING",
		TransactionID: transactionID,
	}

	if err := s.repo.CreatePayment(&payment); err != nil {
		return models.Payment{}, err
	}
	return payment, nil
}

// Update payment status
func (s *PaymentService) UpdatePaymentStatus(id uint, status string) error {
	if status != "PENDING" && status != "SUCCESS" && status != "FAILED" {
		return errors.New("invalid payment status")
	}
	return s.repo.UpdatePaymentStatus(id, status)
}

// Get payment by ID
func (s *PaymentService) GetPaymentByID(id uint) (models.Payment, error) {
	return s.repo.GetPaymentByID(id)
}

// Get payment by order ID
func (s *PaymentService) GetPaymentByOrderID(orderID uint) (models.Payment, error) {
	return s.repo.GetPaymentByOrderID(orderID)
}
