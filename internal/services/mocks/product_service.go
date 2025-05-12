package mocks

import (
	"backend-challenge/internal/models"

	"github.com/stretchr/testify/mock"
)

type ProductService struct {
	mock.Mock
}

func (m *ProductService) ListProducts() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *ProductService) GetProduct(id int64) (*models.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}
