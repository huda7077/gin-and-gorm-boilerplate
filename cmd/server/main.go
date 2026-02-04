package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/huda7077/gin-and-gorm-boilerplate/configs"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/repositories"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/routes"
)

func main() {
	// Load configuration
	config := configs.New()

	// Connect to database
	configs.ConnectDatabase()

	// Initialize repositories
	repos := repositories.NewRepositories(configs.DB)

	// Setup router with dependencies
	r := routes.SetupRouter(repos, config)

	// Health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start server
	port := config.Get("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
