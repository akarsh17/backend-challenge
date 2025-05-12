package errors_test

import (
	"backend-challenge/pkg/errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIErrors(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		code    int
		message string
	}{
		{
			name:    "not found",
			err:     errors.ErrNotFound,
			code:    http.StatusNotFound,
			message: "not_found",
		},
		{
			name:    "unauthorized",
			err:     errors.ErrUnauthorized,
			code:    http.StatusUnauthorized,
			message: "unauthorized",
		},
		{
			name:    "invalid input",
			err:     errors.ErrInvalidInput,
			code:    http.StatusBadRequest,
			message: "invalid_input",
		},
		{
			name:    "items required",
			err:     errors.ErrItemsRequired,
			code:    http.StatusUnprocessableEntity,
			message: "items_required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiErr := tt.err.(errors.APIError)
			assert.Equal(t, tt.code, apiErr.Code)
			assert.Equal(t, tt.message, apiErr.Message)
			assert.Equal(t, tt.message, apiErr.Error())
		})
	}
}

func TestErrorFunctions(t *testing.T) {
	t.Run("NotFoundError", func(t *testing.T) {
		err := errors.NotFoundError("custom not found")
		assert.Equal(t, http.StatusNotFound, err.Code)
		assert.Equal(t, "custom not found", err.Message)
	})

	t.Run("UnauthorizedError", func(t *testing.T) {
		err := errors.UnauthorizedError("custom unauthorized")
		assert.Equal(t, http.StatusUnauthorized, err.Code)
		assert.Equal(t, "custom unauthorized", err.Message)
	})
}
