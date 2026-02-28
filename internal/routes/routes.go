package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huda7077/gin-and-gorm-boilerplate/configs"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/middlewares"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/repositories"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/validators"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// Serve OpenAPI documentation
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/api/docs/openapi.yaml"),
	))
	r.StaticFile("/api/docs/openapi.yaml", "./openapi.yaml")

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
