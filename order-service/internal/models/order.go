package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	UserID     uint        `json:"user_id"`
	Status     string      `json:"status"` // e.g., "pending", "completed"
	TotalPrice float64     `json:"total_price"`
	Items      []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}

type OrderItem struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	OrderID uint    `json:"order_id"`
	MenuID  uint    `json:"menu_id"`
	DishID  uint    `json:"dish_id"`
	Price   float64 `json:"price"`
	Qty     uint    `json:"qty"`
}
