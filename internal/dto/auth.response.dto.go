package dto

import (
	"time"

	"github.com/huda7077/gin-and-gorm-boilerplate/models"
)

// AuthLoginResponse represents login response
type AuthLoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

// AuthRegisterResponse represents register response (if needed)
type AuthRegisterResponse struct {
	Message string `json:"message"`
}

type UserResponse struct {
	ID         uint        `json:"id"`
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	Role       models.Role `json:"role"`
	Image      *string     `json:"image,omitempty"`
	VerifiedAt *time.Time  `json:"verifiedAt"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
}
