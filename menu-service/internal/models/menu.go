package models

import (
	"time"
	"gorm.io/gorm"
)

type Menu struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Name         string         `json:"name"`
	RestaurantID uint           `json:"restaurant_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Dishes       []Dish         `json:"dishes,omitempty"`
}
