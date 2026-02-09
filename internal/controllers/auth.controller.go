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

// Register handles user registration request
// @Summary Register user
// @Description Create a new user account and send verification email
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

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Error(err)
		return
	}

	if err := ctrl.authService.Register(c.Request.Context(), reqBody); err != nil {
		c.Error(err)
		return
	}

	exceptions.SuccessResponse(c, http.StatusCreated, "registration successful, please check your email for verification code", nil)
}

// VerifyEmail handles email verification request
// @Summary Verify email
// @Description Verify user email with OTP code
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dto.AuthVerifyAccountRequest true "Verification data"
// @Success 200 {object} exceptions.Response
// @Failure 400 {object} exceptions.Response
// @Router /auth/verify-email [post]
func (ctrl *AuthController) VerifyEmail(c *gin.Context) {
	var reqBody dto.AuthVerifyAccountRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Error(err)
		return
	}

	if err := ctrl.authService.VerifyEmail(c.Request.Context(), reqBody); err != nil {
		c.Error(err)
		return
	}

	exceptions.SuccessResponse(c, http.StatusOK, "email verified successfully", nil)
}

// ResendVerificationCode handles resend verification code request
// @Summary Resend verification code
// @Description Resend verification email with new OTP
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dto.AuthResendOTPRequest true "Email data"
// @Success 200 {object} exceptions.Response
// @Failure 400 {object} exceptions.Response
// @Router /auth/resend-verification [post]
func (ctrl *AuthController) ResendVerificationCode(c *gin.Context) {
	var reqBody dto.AuthResendOTPRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Error(err)
		return
	}

	if err := ctrl.authService.ResendVerificationCode(c.Request.Context(), reqBody); err != nil {
		c.Error(err)
		return
	}

	exceptions.SuccessResponse(c, http.StatusOK, "verification code sent successfully", nil)
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

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Error(err)
		return
	}

	result, err := ctrl.authService.Login(c.Request.Context(), reqBody)
	if err != nil {
		c.Error(err)
		return
	}

	exceptions.SuccessResponse(c, http.StatusOK, "login successful", result)
}

// ForgotPassword handles forgot password request
// @Summary Forgot password
// @Description Send password reset code to email
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dto.AuthForgotPasswordRequest true "Email data"
// @Success 200 {object} exceptions.Response
// @Failure 400 {object} exceptions.Response
// @Router /auth/forgot-password [post]
func (ctrl *AuthController) ForgotPassword(c *gin.Context) {
	var reqBody dto.AuthForgotPasswordRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Error(err)
		return
	}

	if err := ctrl.authService.ForgotPassword(c.Request.Context(), reqBody); err != nil {
		c.Error(err)
		return
	}

	exceptions.SuccessResponse(c, http.StatusOK, "if email exists, reset code has been sent", nil)
}

// ResetPassword handles reset password request
// @Summary Reset password
// @Description Reset password with OTP code
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dto.AuthResetPasswordRequest true "Reset password data"
// @Success 200 {object} exceptions.Response
// @Failure 400 {object} exceptions.Response
// @Router /auth/reset-password [post]
func (ctrl *AuthController) ResetPassword(c *gin.Context) {
	var reqBody dto.AuthResetPasswordRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Error(err)
		return
	}

	if err := ctrl.authService.ResetPassword(c.Request.Context(), reqBody); err != nil {
		c.Error(err)
		return
	}

	exceptions.SuccessResponse(c, http.StatusOK, "password reset successful", nil)
}
