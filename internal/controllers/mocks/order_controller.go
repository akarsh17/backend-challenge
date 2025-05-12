package mocks

import (
	"backend-challenge/internal/models"

	"github.com/stretchr/testify/mock"
)

type OrderController struct {
	mock.Mock
}

func (m *OrderController) PlaceOrder(req models.OrderRequest) (*models.Order, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Order), args.Error(1)
}
