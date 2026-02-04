package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/huda7077/gin-and-gorm-boilerplate/pkg/exceptions"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		exceptions.ErrorHandler(c, err)
	}
}
