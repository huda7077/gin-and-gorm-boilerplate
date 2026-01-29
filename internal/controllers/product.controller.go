package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/services"
	"github.com/huda7077/gin-and-gorm-boilerplate/models"
)

type ProductController struct {
	Service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{
		Service: service,
	}
}

func (ctr *ProductController) GetAll(c *gin.Context) {
	products, err := ctr.Service.GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})

	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "Get products successfully",
		"data":    products,
	})
}

func (ctr *ProductController) Create(c *gin.Context) {
	var req *models.Product

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if err := ctr.Service.Create(req.Name, req.Price, req.Stock); err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "Create product successfully",
		"data": gin.H{
			"name":  req.Name,
			"price": req.Price,
			"stock": req.Stock,
		},
	})
}
