package models

import (
	"time"
	"gorm.io/gorm"
)

type Restaurant struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	OwnerID   uint           `json:"owner_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Menus     []Menu         `json:"menus,omitempty"`
}
