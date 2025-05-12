package handlers_test

import (
	"backend-challenge/api/handlers"
	"backend-challenge/internal/controllers/mocks"
	"backend-challenge/internal/models"
	"backend-challenge/pkg/errors"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOrderHandlers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("PlaceOrder success", func(t *testing.T) {
		reqBody := `{"items":[{"productId":10,"quantity":1}]}`
		expected := &models.Order{ID: "test-order"}

		mockCtrl := new(mocks.OrderController)
		mockCtrl.On("PlaceOrder", mock.Anything).Return(expected, nil)

		// Create a new handler instance
		handler := handlers.NewOrderHandler(mockCtrl)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/order", bytes.NewBufferString(reqBody))
		c.Request.Header.Set("Content-Type", "application/json")

		handler.PlaceOrder(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"id":"test-order"`)
		mockCtrl.AssertExpectations(t)
	})

	t.Run("PlaceOrder invalid input", func(t *testing.T) {
		reqBody := `invalid json`

		// Create a new handler instance
		handler := handlers.NewOrderHandler(new(mocks.OrderController))

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/order", bytes.NewBufferString(reqBody))
		c.Request.Header.Set("Content-Type", "application/json")

		handler.PlaceOrder(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("PlaceOrder validation error", func(t *testing.T) {
		reqBody := `{"items":[]}`

		mockCtrl := new(mocks.OrderController)
		mockCtrl.On("PlaceOrder", mock.Anything).Return(nil, errors.ValidationError("items required"))

		// Create a new handler instance
		handler := handlers.NewOrderHandler(mockCtrl)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/order", bytes.NewBufferString(reqBody))
		c.Request.Header.Set("Content-Type", "application/json")

		handler.PlaceOrder(c)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		mockCtrl.AssertExpectations(t)
	})
}
