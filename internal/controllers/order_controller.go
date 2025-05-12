package controllers

import (
	"backend-challenge/internal/models"
	"backend-challenge/internal/services"
	"backend-challenge/pkg/errors"
)

type IOrderController interface {
	PlaceOrder(req models.OrderRequest) (*models.Order, error)
}

type OrderController struct {
	service services.IOrderService
}

func NewOrderController(service services.IOrderService) IOrderController {
	return &OrderController{service: service}
}

func (oc *OrderController) PlaceOrder(req models.OrderRequest) (*models.Order, error) {
	if len(req.Items) == 0 {
		return nil, errors.ValidationError("At least one item is required")
	}
	return oc.service.PlaceOrder(req)
}
