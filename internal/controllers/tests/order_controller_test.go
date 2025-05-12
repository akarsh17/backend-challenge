package controllers_test

import (
	"backend-challenge/internal/controllers"
	"backend-challenge/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mock for OrderService ---
type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) PlaceOrder(req models.OrderRequest) (*models.Order, error) {
	args := m.Called(req)
	return args.Get(0).(*models.Order), args.Error(1)
}

func TestPlaceOrder_Success(t *testing.T) {
	mockService := new(MockOrderService)
	controller := controllers.NewOrderController(mockService)

	req := models.OrderRequest{
		Items: []models.OrderItem{{ProductID: 1, Quantity: 2}},
	}
	expectedOrder := &models.Order{ID: "1"}

	mockService.On("PlaceOrder", req).Return(expectedOrder, nil)

	result, err := controller.PlaceOrder(req)

	assert.NoError(t, err)
	assert.Equal(t, expectedOrder, result)
	mockService.AssertExpectations(t)
}

func TestPlaceOrder_EmptyItems(t *testing.T) {
	mockService := new(MockOrderService)
	controller := controllers.NewOrderController(mockService)

	req := models.OrderRequest{
		Items: []models.OrderItem{},
	}

	result, err := controller.PlaceOrder(req)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "At least one item is required")
}
