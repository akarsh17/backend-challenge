package routes

import (
	"backend-challenge/api/handlers"
	"backend-challenge/api/middleware"
	"backend-challenge/internal/controllers"
	"backend-challenge/internal/services"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Initialize services
	productService := services.ProductServiceImpl{}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting the current working directory:", err)
	}

	coupon1 := filepath.Join(dir, "data/couponbase1.gz")
	coupon2 := filepath.Join(dir, "data/couponbase2.gz")
	coupon3 := filepath.Join(dir, "data/couponbase3.gz")
	// Provide actual file paths for coupon data
	couponFiles := []string{
		coupon1,
		coupon2,
		coupon3,
	}
	couponService := services.NewCouponService(couponFiles)

	orderService := services.NewOrderService(productService, couponService)

	// Initialize controllers
	productController := controllers.NewProductController(productService)
	orderController := controllers.NewOrderController(orderService)

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productController)
	orderHandler := handlers.NewOrderHandler(orderController)

	// Product routes
	productGroup := r.Group("/product")
	{
		productGroup.GET("", productHandler.ListProducts)
		productGroup.GET("/:productId", productHandler.GetProduct)
	}

	// Order routes
	orderGroup := r.Group("/order")
	orderGroup.Use(middleware.APIKeyAuthMiddleware())
	{
		orderGroup.POST("", orderHandler.PlaceOrder)
	}

	return r
}
