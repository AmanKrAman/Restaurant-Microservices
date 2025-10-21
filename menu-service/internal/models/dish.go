package models

import (
	"time"

	"gorm.io/gorm"
)

type Dish struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Price     float64        `json:"price"`
	MenuID    uint           `json:"menu_id"` 
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}