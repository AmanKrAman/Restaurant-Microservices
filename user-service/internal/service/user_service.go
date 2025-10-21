package service

import (
	"errors"

	"github.com/amankraman/restaurant-microservices/user-service/internal/models"
	"github.com/amankraman/restaurant-microservices/user-service/internal/repository"
	"github.com/amankraman/restaurant-microservices/user-service/internal/utils"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(name, email, password string) error {
	hashed, _ := utils.HashPassword(password)
	user := &models.User{Name: name, Email: email, Password: hashed}
	return s.repo.Create(user)
}

func (s *UserService) Login(email, password string) (map[string]interface{}, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Name)
	if err != nil {
		return nil, err
	}

	err = s.repo.SaveUserToken(user.ID, token)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"user_id": user.ID,
		"name":    user.Name,
		"email":   user.Email,
		"token":   token,
	}, nil
}

func (s *UserService) GetUsers(input string) ([]models.User, error) {
	return s.repo.GetUsers(input)
}

func (s *UserService) Logout(token string) error {
	return s.repo.DeleteUserToken(token)
}
