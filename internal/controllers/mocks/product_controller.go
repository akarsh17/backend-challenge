package mocks

import (
	"backend-challenge/internal/models"

	"github.com/stretchr/testify/mock"
)

type ProductController struct {
	mock.Mock
}

func (m *ProductController) ListProducts() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *ProductController) GetProduct(id int64) (*models.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}
