package exceptions

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Response represents a standardized API response
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// ErrorHandler handles all errors and returns appropriate responses
func ErrorHandler(c *gin.Context, err error) {
	// Handle custom AppError
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.StatusCode, Response{
			Success: false,
			Message: appErr.Message,
			Data:    appErr.Details,
		})
		c.Abort()
		return
	}

	// Handle validator errors (from Gin's ShouldBindJSON)
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		details := make([]map[string]string, 0)
		for _, fieldErr := range validationErrs {
			details = append(details, map[string]string{
				"field":   fieldErr.Field(),
				"message": getValidationErrorMessage(fieldErr),
			})
		}

		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Validation Error",
			Data:    details,
		})
		c.Abort()
		return
	}

	// Handle legacy error types for backward compatibility
	switch e := err.(type) {
	case ValidationError:
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Validation Error",
			Data:    e.Message,
		})
	case NotFoundError:
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: e.Message,
		})
	case UnauthorizedError:
		c.JSON(http.StatusUnauthorized, Response{
			Success: false,
			Message: e.Message,
		})
	case ConflictError:
		c.JSON(http.StatusConflict, Response{
			Success: false,
			Message: e.Message,
		})
	case BadRequestError:
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: e.Message,
		})
	default:
		// Default internal server error
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
	}

	c.Abort()
}

// getValidationErrorMessage returns a human-readable error message for validation errors
func getValidationErrorMessage(fieldErr validator.FieldError) string {
	field := fieldErr.Field()
	param := fieldErr.Param()

	switch fieldErr.Tag() {
	case "required":
		return field + " is required"

	case "email":
		return field + " must be a valid email address"

	case "min":
		// Different message for string vs number
		if fieldErr.Type().String() == "string" {
			return fmt.Sprintf("%s must be at least %s characters", field, param)
		}
		return fmt.Sprintf("%s must be at least %s", field, param)

	case "max":
		if fieldErr.Type().String() == "string" {
			return fmt.Sprintf("%s must be at most %s characters", field, param)
		}
		return fmt.Sprintf("%s must be at most %s", field, param)

	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", field, param)

	case "gt":
		return fmt.Sprintf("%s must be greater than %s", field, param)

	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, param)

	case "lt":
		return fmt.Sprintf("%s must be less than %s", field, param)

	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, param)

	case "eq":
		return fmt.Sprintf("%s must be equal to %s", field, param)

	case "ne":
		return fmt.Sprintf("%s must not be equal to %s", field, param)

	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", field, param)

	case "alphanum":
		return field + " must contain only alphanumeric characters"

	case "alpha":
		return field + " must contain only alphabetic characters"

	case "numeric":
		return field + " must contain only numeric characters"

	case "url":
		return field + " must be a valid URL"

	case "uri":
		return field + " must be a valid URI"

	case "uuid":
		return field + " must be a valid UUID"

	case "uuid3":
		return field + " must be a valid UUID v3"

	case "uuid4":
		return field + " must be a valid UUID v4"

	case "uuid5":
		return field + " must be a valid UUID v5"

	case "ascii":
		return field + " must contain only ASCII characters"

	case "lowercase":
		return field + " must be in lowercase"

	case "uppercase":
		return field + " must be in uppercase"

	case "datetime":
		return fmt.Sprintf("%s must be a valid datetime in format: %s", field, param)

	case "contains":
		return fmt.Sprintf("%s must contain '%s'", field, param)

	case "containsany":
		return fmt.Sprintf("%s must contain at least one of: %s", field, param)

	case "excludes":
		return fmt.Sprintf("%s must not contain '%s'", field, param)

	case "excludesall":
		return fmt.Sprintf("%s must not contain any of: %s", field, param)

	case "startswith":
		return fmt.Sprintf("%s must start with '%s'", field, param)

	case "endswith":
		return fmt.Sprintf("%s must end with '%s'", field, param)

	case "ip":
		return field + " must be a valid IP address"

	case "ipv4":
		return field + " must be a valid IPv4 address"

	case "ipv6":
		return field + " must be a valid IPv6 address"

	case "json":
		return field + " must be valid JSON"

	case "latitude":
		return field + " must be a valid latitude"

	case "longitude":
		return field + " must be a valid longitude"

	case "strongpassword":
		return field + " must contain at least 8 characters with uppercase, lowercase, and numbers"

	default:
		// Fallback for unknown validation tags
		return fmt.Sprintf("%s is invalid", field)
	}
}

// SuccessResponse returns a standardized success response
func SuccessResponse(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}
