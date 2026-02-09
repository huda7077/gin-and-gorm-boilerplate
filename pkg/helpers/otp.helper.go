package helpers

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GenerateOTP generates a random 6-digit OTP code
func GenerateOTP() (string, error) {
	// Generate 6 random digits
	otp := ""
	for i := 0; i < 6; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		otp += fmt.Sprintf("%d", num)
	}
	return otp, nil
}
