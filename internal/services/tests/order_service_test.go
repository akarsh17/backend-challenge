package services_test

import (
	"backend-challenge/internal/models"
	"backend-challenge/internal/services"
	"backend-challenge/pkg/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestOrderService() services.IOrderService {
	return services.NewOrderService(
		services.ProductServiceImpl{}, // real product service
		dummyCouponService{},          // stubbed coupon service
	)
}

// Dummy coupon service that always validates successfully
type dummyCouponService struct{}

func (d dummyCouponService) ValidateCoupon(code string) (bool, error) {
	return true, nil
}

func TestOrderService_PlaceOrder(t *testing.T) {
	service := getTestOrderService()

	tests := []struct {
		name        string
		request     models.OrderRequest
		expectedErr error
	}{
		{
			name: "valid order",
			request: models.OrderRequest{
				Items: []models.OrderItem{
					{ProductID: 10, Quantity: 2},
				},
			},
		},
		{
			name: "empty items",
			request: models.OrderRequest{
				Items: []models.OrderItem{},
			},
			expectedErr: errors.ValidationError("at least one item is required"),
		},
		{
			name: "invalid product ID",
			request: models.OrderRequest{
				Items: []models.OrderItem{
					{ProductID: 999, Quantity: 1},
				},
			},
			expectedErr: errors.ValidationError("invalid product ID: 999"),
		},
		{
			name: "invalid quantity",
			request: models.OrderRequest{
				Items: []models.OrderItem{
					{ProductID: 10, Quantity: 0},
				},
			},
			expectedErr: errors.ValidationError("invalid quantity for product ID: 10"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, err := service.PlaceOrder(tt.request)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, order)
			} else {
				assert.NoError(t, err)
				assert.Regexp(t, `^[a-f0-9\-]{36}$`, order.ID)
				assert.Len(t, order.Items, 1)
				assert.Len(t, order.Products, 1)
			}
		})
	}
}
