package repository

import (
	"github.com/amankraman/restaurant-microservices/menu-service/internal/models"
	"gorm.io/gorm"
)

type MenuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

// -------------------- Restaurant CRUD --------------------

// CreateRestaurant inserts a new restaurant into the database
func (r *MenuRepository) CreateRestaurant(restaurant *models.Restaurant) error {
	return r.db.Create(restaurant).Error
}

// GetRestaurants fetches all restaurants along with their menus and dishes
func (r *MenuRepository) GetRestaurants() ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	err := r.db.Preload("Menus.Dishes").Find(&restaurants).Error
	return restaurants, err
}

// GetRestaurantByID fetches a single restaurant by ID with menus and dishes
func (r *MenuRepository) GetRestaurantByID(id uint) (models.Restaurant, error) {
	var restaurant models.Restaurant
	err := r.db.Preload("Menus.Dishes").First(&restaurant, id).Error
	return restaurant, err
}

// -------------------- Menu CRUD --------------------

// CreateMenu inserts a new menu into the database
func (r *MenuRepository) CreateMenu(menu *models.Menu) error {
	return r.db.Create(menu).Error
}

// GetMenuByID fetches a single menu by ID along with its dishes
func (r *MenuRepository) GetMenuByID(id uint) (models.Menu, error) {
	var menu models.Menu
	err := r.db.Preload("Dishes").First(&menu, id).Error
	return menu, err
}

// -------------------- Dish CRUD --------------------

// CreateDish inserts a new dish into the database
func (r *MenuRepository) CreateDish(dish *models.Dish) error {
	return r.db.Create(dish).Error
}
