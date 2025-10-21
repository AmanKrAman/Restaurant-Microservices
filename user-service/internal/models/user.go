package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type UserToken struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Token     string    `gorm:"unique;not null" json:"token"`
	LoginAt   time.Time `gorm:"autoCreateTime" json:"login_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
