package controllers_test

import (
	"backend-challenge/internal/controllers"
	"backend-challenge/internal/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mock for ProductService ---
type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) ListProducts() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductService) GetProduct(id int64) (*models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Product), args.Error(1)
}

func TestListProducts_Success(t *testing.T) {
	mockService := new(MockProductService)
	controller := controllers.NewProductController(mockService)

	products := []models.Product{{ID: 1, Name: "Pizza"}, {ID: 2, Name: "Burger"}}
	mockService.On("ListProducts").Return(products, nil)

	result, err := controller.ListProducts()

	assert.NoError(t, err)
	assert.Equal(t, products, result)
	mockService.AssertExpectations(t)
}

func TestGetProduct_Success(t *testing.T) {
	mockService := new(MockProductService)
	controller := controllers.NewProductController(mockService)

	product := &models.Product{ID: 1, Name: "Pizza"}
	mockService.On("GetProduct", int64(1)).Return(product, nil)

	result, err := controller.GetProduct(1)

	assert.NoError(t, err)
	assert.Equal(t, product, result)
	mockService.AssertExpectations(t)
}

func TestGetProduct_NotFound(t *testing.T) {
	mockService := new(MockProductService)
	controller := controllers.NewProductController(mockService)

	mockService.On("GetProduct", int64(999)).Return((*models.Product)(nil), errors.New("product not found"))

	result, err := controller.GetProduct(999)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "product not found", err.Error())
}
