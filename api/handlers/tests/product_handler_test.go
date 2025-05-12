package handlers_test

import (
	"backend-challenge/api/handlers"
	"backend-challenge/internal/controllers/mocks"
	"backend-challenge/internal/models"
	"backend-challenge/pkg/errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestProductHandlers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("ListProducts success", func(t *testing.T) {
		expected := []models.Product{
			{ID: 10, Name: "Test Product"},
		}

		mockCtrl := new(mocks.ProductController)
		mockCtrl.On("ListProducts").Return(expected, nil)

		handler := handlers.NewProductHandler(mockCtrl)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		handler.ListProducts(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"id":10`)
		mockCtrl.AssertExpectations(t)
	})

	t.Run("GetProduct success", func(t *testing.T) {
		productID := int64(10)
		expected := &models.Product{ID: productID}

		mockCtrl := new(mocks.ProductController)
		mockCtrl.On("GetProduct", productID).Return(expected, nil)

		handler := handlers.NewProductHandler(mockCtrl)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "productId", Value: strconv.FormatInt(productID, 10)}}

		handler.GetProduct(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"id":10`)
		mockCtrl.AssertExpectations(t)
	})

	t.Run("GetProduct invalid ID", func(t *testing.T) {
		handler := handlers.NewProductHandler(new(mocks.ProductController))

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "productId", Value: "invalid"}}

		handler.GetProduct(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("GetProduct not found", func(t *testing.T) {
		productID := int64(999)

		mockCtrl := new(mocks.ProductController)
		mockCtrl.On("GetProduct", productID).Return(nil, errors.NotFoundError("not found"))

		handler := handlers.NewProductHandler(mockCtrl)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "productId", Value: strconv.FormatInt(productID, 10)}}

		handler.GetProduct(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
		mockCtrl.AssertExpectations(t)
	})
}
