package errors

import "net/http"

var (
	ErrNotFound      = Error(http.StatusNotFound, "Product not found")
	ErrInvalidInput  = Error(http.StatusBadRequest, "Input is invalid")
	ErrUnauthorized  = Error(http.StatusUnauthorized, "User is unauthorized to perform this request")
	ErrItemsRequired = Error(http.StatusUnprocessableEntity, "At least one item is required")
)

// APIError represents an error response
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e APIError) Error() string {
	return e.Message
}

// NewAPIError creates a new APIError
func Error(code int, message string) APIError {
	return APIError{
		Code:    code,
		Message: message,
	}
}

// Helper functions to create specific error types
func NotFoundError(message string) APIError {
	return Error(http.StatusNotFound, message)
}

func InvalidInputError(message string) APIError {
	return Error(http.StatusInternalServerError, message)
}

func BadRequestError(message string) APIError {
	return Error(http.StatusBadRequest, message)
}

func UnauthorizedError(message string) APIError {
	return Error(http.StatusUnauthorized, message)
}

func ValidationError(message string) APIError {
	return Error(http.StatusUnprocessableEntity, message)
}
