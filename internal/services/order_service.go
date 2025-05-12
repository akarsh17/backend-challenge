package services

import (
	"github.com/google/uuid"

	"backend-challenge/internal/models"
	"backend-challenge/pkg/errors"
	"fmt"
)

// Define interface
type IOrderService interface {
	PlaceOrder(req models.OrderRequest) (*models.Order, error)
}

// Rename implementation
type OrderServiceImpl struct {
	couponService  CouponService   // assume already an interface
	productService IProductService // now use interface
}

func NewOrderService(
	productService IProductService,
	couponService CouponService,
) IOrderService {
	return &OrderServiceImpl{
		productService: productService,
		couponService:  couponService,
	}
}

func (os OrderServiceImpl) PlaceOrder(req models.OrderRequest) (*models.Order, error) {
	if len(req.Items) == 0 {
		return nil, errors.ValidationError("at least one item is required")
	}

	if req.CouponCode != "" {
		valid, err := os.couponService.ValidateCoupon(req.CouponCode)
		if err != nil {
			return nil, errors.BadRequestError("coupon validation failed")
		}
		if !valid {
			return nil, errors.ValidationError("invalid coupon code")
		}
	}

	var products []models.Product
	var orderItems []models.OrderItem

	for _, item := range req.Items {
		product, err := os.productService.GetProduct(item.ProductID)
		if err != nil {
			return nil, errors.ValidationError(fmt.Sprintf("invalid product ID: %d", item.ProductID))
		}

		if item.Quantity <= 0 {
			return nil, errors.ValidationError(fmt.Sprintf("invalid quantity for product ID: %d", item.ProductID))
		}

		products = append(products, *product)
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return &models.Order{
		ID:       uuid.NewString(),
		Items:    orderItems,
		Products: products,
	}, nil
}
