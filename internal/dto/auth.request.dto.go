package dto

// AuthLoginRequest represents login request payload
type AuthLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// AuthRegisterRequest represents register request payload
type AuthRegisterRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,strongpassword"`
}

// AuthVerifyAccountRequest represents verify account request payload
type AuthVerifyAccountRequest struct {
	Email string `json:"email" binding:"required,email"`
	Otp   string `json:"otp" binding:"required"`
}

// AuthResendOTPRequest represents resend OTP request payload
type AuthResendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// AuthForgotPasswordRequest represents forgot password request payload
type AuthForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// AuthResetPasswordRequest represents reset password request payload
type AuthResetPasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Otp         string `json:"otp" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=6"`
}
