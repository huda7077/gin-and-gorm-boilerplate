package exceptions

import "net/http"

// AppError represents a custom application error with status code and message
type AppError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Details    any    `json:"details,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

// NewAppError creates a new AppError
func NewAppError(statusCode int, message string, details any) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		Details:    details,
	}
}

// Common error constructors
func NewBadRequestError(message string, details any) *AppError {
	if message == "" {
		message = "Bad Request"
	}
	return NewAppError(http.StatusBadRequest, message, details)
}

func NewUnauthorizedError(message string) *AppError {
	if message == "" {
		message = "Unauthorized"
	}
	return NewAppError(http.StatusUnauthorized, message, nil)
}

func NewForbiddenError(message string) *AppError {
	if message == "" {
		message = "Forbidden"
	}
	return NewAppError(http.StatusForbidden, message, nil)
}

func NewNotFoundError(message string) *AppError {
	if message == "" {
		message = "Resource Not Found"
	}
	return NewAppError(http.StatusNotFound, message, nil)
}

func NewConflictError(message string) *AppError {
	if message == "" {
		message = "Resource Conflict"
	}
	return NewAppError(http.StatusConflict, message, nil)
}

func NewInternalServerError(message string) *AppError {
	if message == "" {
		message = "Internal Server Error"
	}
	return NewAppError(http.StatusInternalServerError, message, nil)
}

func NewValidationError(message string, details any) *AppError {
	if message == "" {
		message = "Validation Error"
	}
	return NewAppError(http.StatusBadRequest, message, details)
}
