package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/huda7077/gin-and-gorm-boilerplate/configs"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	port := 8080

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	configs.ConnectDatabase()

	r := routes.SetupRouter()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(fmt.Sprintf(":%d", port))
}
