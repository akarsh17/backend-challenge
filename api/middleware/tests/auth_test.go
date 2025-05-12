package middleware_test

import (
	"backend-challenge/api/middleware"
	"backend-challenge/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mocking AppConfig for testing
	config.AppConfig = &config.Config{
		APIKey: "apitest", // Set the key directly
	}

	tests := []struct {
		name           string
		apiKey         string
		statusCode     int
		expectedErrMsg string
	}{
		{
			name:           "valid api key",
			apiKey:         "apitest",
			statusCode:     http.StatusOK,
			expectedErrMsg: "",
		},
		{
			name:           "missing api key",
			apiKey:         "",
			statusCode:     http.StatusUnauthorized,
			expectedErrMsg: "Invalid API key", // Match the error message here
		},
		{
			name:           "invalid api key",
			apiKey:         "wrongkey",
			statusCode:     http.StatusUnauthorized,
			expectedErrMsg: "Invalid API key", // Match the error message here
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)

			// Use the middleware
			r.Use(middleware.APIKeyAuthMiddleware())

			// Test endpoint
			r.GET("/test", func(c *gin.Context) {
				c.Status(http.StatusOK)
			})

			// Create the request
			req := httptest.NewRequest("GET", "/test", nil)
			if tt.apiKey != "" {
				req.Header.Set("api_key", tt.apiKey)
			}

			// Serve the HTTP request
			r.ServeHTTP(w, req)

			// Assert status and error response if unauthorized
			assert.Equal(t, tt.statusCode, w.Code)
			if tt.statusCode == http.StatusUnauthorized {
				assert.Contains(t, w.Body.String(), tt.expectedErrMsg)
			}
		})
	}
}
