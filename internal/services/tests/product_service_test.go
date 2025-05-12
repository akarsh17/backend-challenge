package services_test

import (
	"backend-challenge/internal/models"
	"backend-challenge/internal/services"
	"backend-challenge/pkg/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestProductService() services.IProductService {
	return services.ProductServiceImpl{}
}

func TestProductService_ListProducts(t *testing.T) {
	service := getTestProductService()

	t.Run("should return list of products", func(t *testing.T) {
		products, err := service.ListProducts()

		assert.NoError(t, err)
		assert.Len(t, products, 2)
		assert.Equal(t, int64(10), products[0].ID)
		assert.Equal(t, "Chicken Waffle", products[0].Name)
	})
}

func TestProductService_GetProduct(t *testing.T) {
	service := getTestProductService()

	tests := []struct {
		name        string
		productID   int64
		expected    *models.Product
		expectedErr error
	}{
		{
			name:      "existing product",
			productID: 10,
			expected: &models.Product{
				ID:       10,
				Name:     "Chicken Waffle",
				Price:    5.99,
				Category: "Waffle",
			},
		},
		{
			name:        "non-existent product",
			productID:   999,
			expectedErr: errors.NotFoundError("product not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := service.GetProduct(tt.productID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, product)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, product)
			}
		})
	}
}
