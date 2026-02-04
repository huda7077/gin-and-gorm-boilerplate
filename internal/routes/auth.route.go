package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huda7077/gin-and-gorm-boilerplate/configs"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/controllers"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/repositories"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/services"
)

// AuthRoutes configures authentication routes
func AuthRoutes(rg *gin.RouterGroup, repos *repositories.Repositories, config configs.Config) {
	// Initialize service with dependencies
	authService := services.NewAuthService(repos, config)

	// Initialize controller with service
	authController := controllers.NewAuthController(authService)

	// Setup routes
	auth := rg.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
	}
}
