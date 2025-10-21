package handlers

import (
	"net/http"
	"strconv"

	"github.com/amankraman/restaurant-microservices/menu-service/internal/models"
	"github.com/amankraman/restaurant-microservices/menu-service/internal/service"
	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	service *service.MenuService
}

func NewMenuHandler(service *service.MenuService) *MenuHandler {
	return &MenuHandler{service: service}
}

// -------------------- Restaurants --------------------

// CreateRestaurant creates a new restaurant
func (h *MenuHandler) CreateRestaurant(c *gin.Context) {
	var r models.Restaurant
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateRestaurant(&r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}

// GetRestaurants returns all restaurants with their menus and dishes
func (h *MenuHandler) GetRestaurants(c *gin.Context) {
	restaurants, err := h.service.GetRestaurants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, restaurants)
}

// GetRestaurantByID returns a single restaurant by ID with menus and dishes
func (h *MenuHandler) GetRestaurantByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid restaurant ID"})
		return
	}

	restaurant, err := h.service.GetRestaurantByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	c.JSON(http.StatusOK, restaurant)
}

// -------------------- Menus --------------------

// CreateMenu creates a new menu for an existing restaurant
func (h *MenuHandler) CreateMenu(c *gin.Context) {
	var m models.Menu
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate that the restaurant exists
	if _, err := h.service.GetRestaurantByID(m.RestaurantID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid restaurant_id"})
		return
	}

	if err := h.service.CreateMenu(&m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, m)
}

// -------------------- Dishes --------------------

// CreateDish creates a new dish under an existing menu
func (h *MenuHandler) CreateDish(c *gin.Context) {
	var d models.Dish
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Optional: Validate that the menu exists
	if _, err := h.service.GetMenuByID(d.MenuID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu_id"})
		return
	}

	if err := h.service.CreateDish(&d); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, d)
}
