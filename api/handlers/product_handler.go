package handlers

import (
	"backend-challenge/internal/controllers"
	"backend-challenge/pkg/errors"
	response "backend-challenge/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	controller controllers.IProductController
}

func NewProductHandler(controller controllers.IProductController) *ProductHandler {
	return &ProductHandler{controller: controller}
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.controller.ListProducts()
	if err != nil {
		response.Error(c, errors.NotFoundError("Unable to fetch products"))
		return
	}
	response.Success(c, http.StatusOK, products)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	productId, err := strconv.ParseInt(c.Param("productId"), 10, 64)
	if err != nil {
		response.Error(c, errors.BadRequestError("Invalid product ID"))
		return
	}

	product, err := h.controller.GetProduct(productId)
	if err != nil {
		response.Error(c, errors.NotFoundError("Product not found"))
		return
	}
	response.Success(c, http.StatusOK, product)
}
