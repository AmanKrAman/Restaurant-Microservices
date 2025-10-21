package handlers

import (
	"net/http"
	"strconv"

	// "github.com/amankraman/restaurant-microservices/payment-service/internal/models"
	"github.com/amankraman/restaurant-microservices/payment-service/internal/service"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(service *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

// Create a payment
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req struct {
		OrderID uint    `json:"order_id"`
		UserID  uint    `json:"user_id"`         // add user_id
		Amount  float64 `json:"amount"`          // backend can calculate from order total
		Method  string  `json:"payment_method"`  // rename to match service argument
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment, err := h.service.CreatePayment(req.OrderID, req.UserID, req.Amount, req.Method)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payment)
}
// Update payment status
func (h *PaymentHandler) UpdatePaymentStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdatePaymentStatus(uint(id), req.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "status updated"})
}

// Get payment by ID
func (h *PaymentHandler) GetPaymentByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	payment, err := h.service.GetPaymentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}
	c.JSON(http.StatusOK, payment)
}

// Get payment by Order ID
func (h *PaymentHandler) GetPaymentByOrderID(c *gin.Context) {
	orderID, _ := strconv.Atoi(c.Param("order_id"))
	payment, err := h.service.GetPaymentByOrderID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}
	c.JSON(http.StatusOK, payment)
}
