package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huda7077/gin-and-gorm-boilerplate/configs"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/middlewares"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/repositories"
	"github.com/huda7077/gin-and-gorm-boilerplate/pkg/exceptions"
)

func ProtectedRoutes(rg *gin.RouterGroup, repos *repositories.Repositories, config configs.Config) {
	protected := rg.Group("/protected")
	{
		protected.GET("/profile", middlewares.AuthMiddleware(), func(c *gin.Context) {
			exceptions.SuccessResponse(c, http.StatusOK, "token valid", nil)

		})
	}
}
