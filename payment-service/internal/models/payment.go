package models

import (
	"time"
)

type Payment struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	OrderID       uint      `json:"order_id" gorm:"not null"`
	UserID        uint      `json:"user_id" gorm:"not null"`
	Amount        float64   `json:"amount" gorm:"not null"`
	PaymentMethod string    `json:"payment_method" gorm:"not null;column:method"`
	Status        string    `json:"status" gorm:"not null"`
	TransactionID string    `json:"transaction_id" gorm:"not null;unique"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
