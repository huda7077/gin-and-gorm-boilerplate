package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/dto"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/services"
	"github.com/huda7077/gin-and-gorm-boilerplate/pkg/exceptions"
)

// AuthController handles authentication related requests
type AuthController struct {
	authService services.AuthService
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Login handles user login request
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dto.AuthLoginRequest true "Login credentials"
// @Success 200 {object} exceptions.Response
// @Failure 400 {object} exceptions.Response
// @Failure 401 {object} exceptions.Response
// @Router /auth/login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var reqBody dto.AuthLoginRequest

	// Bind and validate request body
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Error(err)
		return
	}

	// Call service
	result, err := ctrl.authService.Login(c.Request.Context(), reqBody)
	if err != nil {
		c.Error(err)
		return
	}

	// Return success response
	exceptions.SuccessResponse(c, http.StatusOK, "login successful", result)
}

// Register handles user registration request
// @Summary Register user
// @Description Create a new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dto.AuthRegisterRequest true "Registration data"
// @Success 201 {object} exceptions.Response
// @Failure 400 {object} exceptions.Response
// @Failure 409 {object} exceptions.Response
// @Router /auth/register [post]
func (ctrl *AuthController) Register(c *gin.Context) {
	var reqBody dto.AuthRegisterRequest

	// Bind and validate request body
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Error(err)
		return
	}

	// Call service
	if err := ctrl.authService.Register(c.Request.Context(), reqBody); err != nil {
		c.Error(err)
		return
	}

	// Return success response
	exceptions.SuccessResponse(c, http.StatusCreated, "registration successful", nil)
}
