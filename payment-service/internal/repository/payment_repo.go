package repository

import (
	"github.com/amankraman/restaurant-microservices/payment-service/internal/models"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) CreatePayment(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *PaymentRepository) GetPaymentByID(id uint) (models.Payment, error) {
	var payment models.Payment
	err := r.db.First(&payment, id).Error
	return payment, err
}

func (r *PaymentRepository) UpdatePaymentStatus(id uint, status string) error {
	return r.db.Model(&models.Payment{}).Where("id = ?", id).Update("status", status).Error
}

func (r *PaymentRepository) GetPaymentByOrderID(orderID uint) (models.Payment, error) {
	var payment models.Payment
	err := r.db.Where("order_id = ?", orderID).First(&payment).Error
	return payment, err
}
