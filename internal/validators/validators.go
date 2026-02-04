package validators

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func SetupCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("strongpassword", func(fl validator.FieldLevel) bool {
			password := fl.Field().String()
			// Check: min 8 chars, contains number, uppercase, lowercase
			hasMinLen := len(password) >= 8
			hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
			hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
			hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
			return hasMinLen && hasNumber && hasUpper && hasLower
		})
	}
}
