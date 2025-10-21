package repository

import (
	"strings"
	"time"

	"github.com/amankraman/restaurant-microservices/user-service/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) GetUsers(input string) ([]models.User, error) {
	var users []models.User
	query := r.DB

	if input != "" {
		ilikeInput := "%" + strings.ToLower(input) + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(email) LIKE ?", ilikeInput, ilikeInput)
	}

	err := query.Order("created_at DESC").Find(&users).Error
	return users, err
}

func (r *UserRepository) SaveUserToken(userID uint, token string) error {
	userToken := models.UserToken{
		UserID:  userID,
		Token:   token,
		LoginAt: time.Now(),
		// ExpiresAt: time.Now().Add(24 * time.Hour), // optional
	}
	return r.DB.Create(&userToken).Error
}

func (r *UserRepository) DeleteUserToken(token string) error {
	return r.DB.Where("token = ?", token).Delete(&models.UserToken{}).Error
}
