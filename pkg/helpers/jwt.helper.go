package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/huda7077/gin-and-gorm-boilerplate/models"
)

func GenerateJWT(Id uint, email string, role models.Role, JWTExpiryTime time.Duration) (string, error) {
	JWTSecretKey := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"userId": Id,
		"email":  email,
		"role":   role,
		"exp":    time.Now().Add(JWTExpiryTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// How to pass JWTExpiryTime Argument:
// JWTExpiryTime := time.Hour * 24
// GenerateJWT(Id, role, JWTExpiryTime, JWTSecretKey)
