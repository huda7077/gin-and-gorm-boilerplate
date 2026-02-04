package services

import (
	"context"
	"time"

	"github.com/huda7077/gin-and-gorm-boilerplate/configs"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/dto"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/repositories"
	"github.com/huda7077/gin-and-gorm-boilerplate/models"
	"github.com/huda7077/gin-and-gorm-boilerplate/pkg/exceptions"
	"github.com/huda7077/gin-and-gorm-boilerplate/pkg/helpers"
	"golang.org/x/crypto/bcrypt"
)

// AuthService defines methods for authentication operations
type AuthService interface {
	Login(ctx context.Context, reqBody dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	Register(ctx context.Context, reqBody dto.AuthRegisterRequest) error
}

// authServiceImpl implements AuthService interface
type authServiceImpl struct {
	repositories *repositories.Repositories
	config       configs.Config
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(repositories *repositories.Repositories, config configs.Config) AuthService {
	return &authServiceImpl{
		repositories: repositories,
		config:       config,
	}
}

// Login authenticates a user and returns user data with JWT token
func (s *authServiceImpl) Login(ctx context.Context, reqBody dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	// Find user by email
	user, err := s.repositories.User.FindByEmail(ctx, reqBody.Email)
	if err != nil {
		// Return generic error message for security (don't reveal if email exists)
		return nil, exceptions.NewUnauthorizedError("invalid email or password")
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password)); err != nil {
		return nil, exceptions.NewUnauthorizedError("invalid email or password")
	}

	// Generate JWT token
	tokenString, err := helpers.GenerateJWT(user.ID, user.Email, user.Role, time.Hour*72)
	if err != nil {
		return nil, exceptions.NewInternalServerError("failed to generate token")
	}

	// Build response
	response := &dto.AuthLoginResponse{
		User: dto.UserResponse{
			ID:         user.ID,
			Name:       user.Name,
			Email:      user.Email,
			Role:       user.Role,
			Image:      user.Image,
			VerifiedAt: user.VerifiedAt,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		},
		Token: tokenString,
	}

	// if user.Image != nil {
	// 	response.User.Image = *user.Image
	// }

	return response, nil
}

// Register creates a new user account
func (s *authServiceImpl) Register(ctx context.Context, reqBody dto.AuthRegisterRequest) error {
	// Check if email already exists
	existingUser, _ := s.repositories.User.FindByEmail(ctx, reqBody.Email)
	if existingUser.ID != 0 {
		return exceptions.NewConflictError("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return exceptions.NewInternalServerError("failed to hash password")
	}

	// Create user
	user := models.User{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Password: string(hashedPassword),
		Role:     models.RoleUser,
	}

	_, err = s.repositories.User.Create(ctx, user)
	if err != nil {
		return exceptions.NewInternalServerError("failed to create user")
	}

	return nil
}
