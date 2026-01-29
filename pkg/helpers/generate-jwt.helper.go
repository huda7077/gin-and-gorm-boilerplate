package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(Id string, role string, JWTExpiryTime time.Duration, JWTSecretKey string) (string, error) {
	claims := jwt.MapClaims{
		"userId": Id,
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
