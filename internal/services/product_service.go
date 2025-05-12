package services

import (
	"backend-challenge/internal/models"
	"backend-challenge/pkg/errors"
)

// Define interface
type IProductService interface {
	ListProducts() ([]models.Product, error)
	GetProduct(id int64) (*models.Product, error)
}

// Rename implementation
type ProductServiceImpl struct{}

func (ps ProductServiceImpl) ListProducts() ([]models.Product, error) {
	return []models.Product{
		{ID: 10, Name: "Chicken Waffle", Price: 5.99, Category: "Waffle"},
		{ID: 11, Name: "Veggie Burger", Price: 7.49, Category: "Burger"},
	}, nil
}

func (ps ProductServiceImpl) GetProduct(id int64) (*models.Product, error) {
	products, err := ps.ListProducts()
	if err != nil {
		return nil, err
	}

	for _, p := range products {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, errors.NotFoundError("product not found")
}
