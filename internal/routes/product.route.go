package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/controllers"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/services"
)

func ProductRouter(rg *gin.RouterGroup) {
	productService := services.NewProductService()
	productController := controllers.NewProductController(productService)

	product := rg.Group("/products")
	{
		product.GET("/", productController.GetAll)
	}
}
