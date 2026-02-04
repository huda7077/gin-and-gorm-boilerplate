package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huda7077/gin-and-gorm-boilerplate/configs"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/middlewares"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/repositories"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/validators"
)

// SetupRouter initializes all routes for the application
func SetupRouter(repos *repositories.Repositories, config configs.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// Enable CORS
	r.Use(middlewares.CORSMiddleware())

	validators.SetupCustomValidators()

	// Enable method not allowed handler
	r.HandleMethodNotAllowed = true

	// Apply global error handling middleware
	r.Use(middlewares.ErrorMiddleware())

	// Setup API routes
	api := r.Group("/api")
	{
		// Setup feature routes with dependency injection
		AuthRoutes(api, repos, config)
		// Add more routes here as needed
		// ProductRoutes(api, repos, config)
		// UserRoutes(api, repos, config)
		ProtectedRoutes(api, repos, config)
	}

	return r
}
