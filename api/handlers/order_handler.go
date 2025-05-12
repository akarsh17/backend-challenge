package handlers

import (
	"backend-challenge/internal/controllers"
	"backend-challenge/internal/models"
	"backend-challenge/pkg/errors"
	response "backend-challenge/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	controller controllers.IOrderController
}

func NewOrderHandler(controller controllers.IOrderController) *OrderHandler {
	return &OrderHandler{controller: controller}
}

func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	var req models.OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errors.BadRequestError("Invalid input"))
		return
	}

	order, err := h.controller.PlaceOrder(req)
	if err != nil {
		// Handle different error types
		switch err.(type) {
		case errors.APIError:
			// Pass through the API error directly
			response.Error(c, err.(errors.APIError))
		default:
			// Convert unknown errors to InvalidInputError
			response.Error(c, errors.InvalidInputError("Unable to process order"))
		}
		return
	}
	response.Success(c, http.StatusOK, order)
}
