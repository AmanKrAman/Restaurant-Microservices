package service

import (
	"github.com/amankraman/restaurant-microservices/menu-service/internal/models"
	"github.com/amankraman/restaurant-microservices/menu-service/internal/repository"
)

type MenuService struct {
	repo *repository.MenuRepository
}

func NewMenuService(repo *repository.MenuRepository) *MenuService {
	return &MenuService{repo: repo}
}

// -------------------- Restaurant CRUD --------------------

// CreateRestaurant creates a new restaurant
func (s *MenuService) CreateRestaurant(restaurant *models.Restaurant) error {
	return s.repo.CreateRestaurant(restaurant)
}

// GetRestaurants returns all restaurants with menus and dishes
func (s *MenuService) GetRestaurants() ([]models.Restaurant, error) {
	return s.repo.GetRestaurants()
}

// GetRestaurantByID returns a single restaurant by ID
func (s *MenuService) GetRestaurantByID(id uint) (models.Restaurant, error) {
	return s.repo.GetRestaurantByID(id)
}

// -------------------- Menu CRUD --------------------

// CreateMenu creates a new menu
func (s *MenuService) CreateMenu(menu *models.Menu) error {
	return s.repo.CreateMenu(menu)
}

// GetMenuByID returns a single menu by ID with dishes
func (s *MenuService) GetMenuByID(id uint) (models.Menu, error) {
	return s.repo.GetMenuByID(id)
}

// -------------------- Dish CRUD --------------------

// CreateDish creates a new dish
func (s *MenuService) CreateDish(dish *models.Dish) error {
	return s.repo.CreateDish(dish)
}
