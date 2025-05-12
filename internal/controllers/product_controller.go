package controllers

import (
	"backend-challenge/internal/models"
	"backend-challenge/internal/services"
)

type IProductController interface {
	ListProducts() ([]models.Product, error)
	GetProduct(id int64) (*models.Product, error)
}

type ProductController struct {
	service services.IProductService
}

func NewProductController(service services.IProductService) IProductController {
	return &ProductController{service: service}
}

func (pc *ProductController) ListProducts() ([]models.Product, error) {
	return pc.service.ListProducts()
}

func (pc *ProductController) GetProduct(id int64) (*models.Product, error) {
	return pc.service.GetProduct(id)
}
