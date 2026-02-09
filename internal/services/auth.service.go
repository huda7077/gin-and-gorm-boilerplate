package services

import (
	"context"
	"time"

	"github.com/huda7077/gin-and-gorm-boilerplate/configs"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/dto"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/provider/mail"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/repositories"
	"github.com/huda7077/gin-and-gorm-boilerplate/models"
	"github.com/huda7077/gin-and-gorm-boilerplate/pkg/exceptions"
	"github.com/huda7077/gin-and-gorm-boilerplate/pkg/helpers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthService defines methods for authentication operations
type AuthService interface {
	Register(ctx context.Context, reqBody dto.AuthRegisterRequest) error
	VerifyEmail(ctx context.Context, reqBody dto.AuthVerifyAccountRequest) error
	ResendVerificationCode(ctx context.Context, reqBody dto.AuthResendOTPRequest) error
	Login(ctx context.Context, reqBody dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	ForgotPassword(ctx context.Context, reqBody dto.AuthForgotPasswordRequest) error
	ResetPassword(ctx context.Context, reqBody dto.AuthResetPasswordRequest) error
}

// authServiceImpl implements AuthService interface
type authServiceImpl struct {
	repositories *repositories.Repositories
	config       configs.Config
	mailProvider *mail.Provider
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(repositories *repositories.Repositories, config configs.Config) AuthService {
	emailConfig := configs.NewEmail(config)
	mailProvider := mail.NewMailProvider(
		emailConfig.Host,
		emailConfig.Port,
		emailConfig.Username,
		emailConfig.Password,
		emailConfig.From,
	)

	return &authServiceImpl{
		repositories: repositories,
		config:       config,
		mailProvider: mailProvider,
	}
}

// Register creates a new user account and sends verification email
// Uses DB transaction to ensure atomicity
func (s *authServiceImpl) Register(ctx context.Context, reqBody dto.AuthRegisterRequest) error {
	// Check if email already exists (outside transaction for performance)
	existingUser, _ := s.repositories.User.FindByEmail(ctx, reqBody.Email)
	if existingUser.ID != 0 {
		return exceptions.NewConflictError("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return exceptions.NewInternalServerError("failed to hash password")
	}

	// Generate OTP
	otp, err := helpers.GenerateOTP()
	if err != nil {
		return exceptions.NewInternalServerError("failed to generate verification code")
	}

	var createdUser models.User

	// Start transaction
	err = s.repositories.DB.Transaction(func(tx *gorm.DB) error {
		// Create transaction-aware repositories
		txRepos := s.repositories.WithTx(tx)

		// Create user
		user := models.User{
			Name:     reqBody.Name,
			Email:    reqBody.Email,
			Password: string(hashedPassword),
			Role:     models.RoleUser,
		}

		createdUser, err = txRepos.User.Create(ctx, user)
		if err != nil {
			return err
		}

		// Save verification code
		expiredMinutes := 10
		verificationCode := models.VerificationCode{
			UserID:    createdUser.ID,
			Code:      otp,
			Purpose:   "EMAIL_VERIFICATION",
			ExpiredAt: time.Now().Add(time.Duration(expiredMinutes) * time.Minute),
		}

		_, err = txRepos.VerificationCode.Create(ctx, verificationCode)
		if err != nil {
			return err
		}

		// Transaction will be committed if no error is returned
		return nil
	})

	if err != nil {
		return exceptions.NewInternalServerError("failed to create user")
	}

	// Send verification email (outside transaction - non-critical operation)
	// If email fails, user is still created and can resend verification
	emailData := struct {
		OTP            string
		ExpiredMinutes int
	}{
		OTP:            otp,
		ExpiredMinutes: 10,
	}

	_ = s.mailProvider.SendMail(
		createdUser.Email,
		"Verify Your Email Address",
		"verify-account.html",
		emailData,
	)
	// Note: We don't return error if email fails, as user is already created
	// User can use resend verification endpoint

	return nil
}

// VerifyEmail verifies user email with OTP code
// Uses DB transaction to ensure atomicity
func (s *authServiceImpl) VerifyEmail(ctx context.Context, reqBody dto.AuthVerifyAccountRequest) error {
	// Find user by email (outside transaction)
	user, err := s.repositories.User.FindByEmail(ctx, reqBody.Email)
	if err != nil {
		return exceptions.NewNotFoundError("user not found")
	}

	// Check if already verified
	if user.VerifiedAt != nil {
		return exceptions.NewBadRequestError("email already verified", nil)
	}

	// Start transaction
	err = s.repositories.DB.Transaction(func(tx *gorm.DB) error {
		txRepos := s.repositories.WithTx(tx)

		// Find and validate verification code
		_, err = txRepos.VerificationCode.FindValidCode(ctx, user.ID, reqBody.Otp, "EMAIL_VERIFICATION")
		if err != nil {
			return exceptions.NewBadRequestError("invalid or expired verification code", nil)
		}

		// Update user verified_at
		now := time.Now()
		user.VerifiedAt = &now
		_, err = txRepos.User.Update(ctx, int(user.ID), user)
		if err != nil {
			return err
		}

		// Delete used verification code
		err = txRepos.VerificationCode.DeleteByUser(ctx, user.ID, "EMAIL_VERIFICATION")
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		// Check if it's already our custom error
		if _, ok := err.(*exceptions.AppError); ok {
			return err
		}
		return exceptions.NewInternalServerError("failed to verify email")
	}

	return nil
}

// ResendVerificationCode resends verification email
// Uses DB transaction to ensure atomicity
func (s *authServiceImpl) ResendVerificationCode(ctx context.Context, reqBody dto.AuthResendOTPRequest) error {
	// Find user by email
	user, err := s.repositories.User.FindByEmail(ctx, reqBody.Email)
	if err != nil {
		return exceptions.NewNotFoundError("user not found")
	}

	// Check if already verified
	if user.VerifiedAt != nil {
		return exceptions.NewBadRequestError("email already verified", nil)
	}

	// Generate new OTP
	otp, err := helpers.GenerateOTP()
	if err != nil {
		return exceptions.NewInternalServerError("failed to generate verification code")
	}

	// Start transaction
	err = s.repositories.DB.Transaction(func(tx *gorm.DB) error {
		txRepos := s.repositories.WithTx(tx)

		// Delete old verification codes
		err = txRepos.VerificationCode.DeleteByUser(ctx, user.ID, "EMAIL_VERIFICATION")
		if err != nil {
			return err
		}

		// Save new verification code
		expiredMinutes := 10
		verificationCode := models.VerificationCode{
			UserID:    user.ID,
			Code:      otp,
			Purpose:   "EMAIL_VERIFICATION",
			ExpiredAt: time.Now().Add(time.Duration(expiredMinutes) * time.Minute),
		}

		_, err = txRepos.VerificationCode.Create(ctx, verificationCode)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return exceptions.NewInternalServerError("failed to save verification code")
	}

	// Send verification email (outside transaction)
	emailData := struct {
		OTP            string
		ExpiredMinutes int
	}{
		OTP:            otp,
		ExpiredMinutes: 10,
	}

	err = s.mailProvider.SendMail(
		user.Email,
		"Verify Your Email Address",
		"verify-account.html",
		emailData,
	)
	if err != nil {
		return exceptions.NewInternalServerError("failed to send verification email")
	}

	return nil
}

// Login authenticates a user and returns user data with JWT token
func (s *authServiceImpl) Login(ctx context.Context, reqBody dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	// Find user by email
	user, err := s.repositories.User.FindByEmail(ctx, reqBody.Email)
	if err != nil {
		return nil, exceptions.NewUnauthorizedError("invalid email or password")
	}

	// Check if email is verified
	if user.VerifiedAt == nil {
		return nil, exceptions.NewUnauthorizedError("please verify your email first")
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

	return response, nil
}

// ForgotPassword sends password reset code to user email
// Uses DB transaction to ensure atomicity
func (s *authServiceImpl) ForgotPassword(ctx context.Context, reqBody dto.AuthForgotPasswordRequest) error {
	// Find user by email
	user, err := s.repositories.User.FindByEmail(ctx, reqBody.Email)
	if err != nil {
		// Return success even if user not found (security best practice)
		// Don't reveal if email exists or not
		return nil
	}

	// Generate OTP
	otp, err := helpers.GenerateOTP()
	if err != nil {
		return exceptions.NewInternalServerError("failed to generate reset code")
	}

	// Start transaction
	err = s.repositories.DB.Transaction(func(tx *gorm.DB) error {
		txRepos := s.repositories.WithTx(tx)

		// Delete old reset codes
		err = txRepos.VerificationCode.DeleteByUser(ctx, user.ID, "RESET_PASSWORD")
		if err != nil {
			return err
		}

		// Save reset code
		expiredMinutes := 10
		verificationCode := models.VerificationCode{
			UserID:    user.ID,
			Code:      otp,
			Purpose:   "RESET_PASSWORD",
			ExpiredAt: time.Now().Add(time.Duration(expiredMinutes) * time.Minute),
		}

		_, err = txRepos.VerificationCode.Create(ctx, verificationCode)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return exceptions.NewInternalServerError("failed to save reset code")
	}

	// Send reset email (outside transaction)
	emailData := struct {
		OTP            string
		ExpiredMinutes int
	}{
		OTP:            otp,
		ExpiredMinutes: 10,
	}

	err = s.mailProvider.SendMail(
		user.Email,
		"Reset Your Password",
		"reset-password.html",
		emailData,
	)
	if err != nil {
		return exceptions.NewInternalServerError("failed to send reset email")
	}

	return nil
}

// ResetPassword resets user password with OTP code
// Uses DB transaction to ensure atomicity
func (s *authServiceImpl) ResetPassword(ctx context.Context, reqBody dto.AuthResetPasswordRequest) error {
	// Find user by email
	user, err := s.repositories.User.FindByEmail(ctx, reqBody.Email)
	if err != nil {
		return exceptions.NewNotFoundError("user not found")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return exceptions.NewInternalServerError("failed to hash password")
	}

	// Start transaction
	err = s.repositories.DB.Transaction(func(tx *gorm.DB) error {
		txRepos := s.repositories.WithTx(tx)

		// Find and validate reset code
		_, err = txRepos.VerificationCode.FindValidCode(ctx, user.ID, reqBody.Otp, "RESET_PASSWORD")
		if err != nil {
			return exceptions.NewBadRequestError("invalid or expired reset code", nil)
		}

		// Update password
		user.Password = string(hashedPassword)
		_, err = txRepos.User.Update(ctx, int(user.ID), user)
		if err != nil {
			return err
		}

		// Delete used reset code
		err = txRepos.VerificationCode.DeleteByUser(ctx, user.ID, "RESET_PASSWORD")
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		// Check if it's already our custom error
		if _, ok := err.(*exceptions.AppError); ok {
			return err
		}
		return exceptions.NewInternalServerError("failed to reset password")
	}

	return nil
}
