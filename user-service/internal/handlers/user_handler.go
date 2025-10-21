package handlers

import (
	"net/http"
	"restaurant-microservices/common/utils"
	"strings"

	"github.com/amankraman/restaurant-microservices/user-service/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.service.Register(req.Name, req.Email, req.Password)
	rb := &utils.ResponseBuilder{}
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	response, err := rb.GenerateResponse(true, "user registered", map[string]interface{}{"name": req.Name, "email": req.Email}, nil)

	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Internal Server Error"})
		return
	}
	c.Data(200, "application/json", []byte(response))
	// c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	rb := &utils.ResponseBuilder{}
	user, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	response, err := rb.GenerateResponse(true, "user registered", map[string]interface{}{"user": user}, nil)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Internal Server Error"})
		return
	}
	c.Data(200, "application/json", []byte(response))
	// c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	input := c.Query("input")

	users, err := h.service.GetUsers(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rb := &utils.ResponseBuilder{}
	response, err := rb.GenerateResponse(true, "users fetched successfully", users, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal Server Error"})
		return
	}

	c.Data(200, "application/json", []byte(response))
}

func (h *UserHandler) Logout(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "No token provided"})
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	err := h.service.Logout(tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Logout failed"})
		return
	}

	rb := &utils.ResponseBuilder{}
	response, err := rb.GenerateResponse(true, "Logged out successfully", nil, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal Server Error"})
		return
	}

	c.Data(200, "application/json", []byte(response))

	// c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Logged out successfully"})
}
