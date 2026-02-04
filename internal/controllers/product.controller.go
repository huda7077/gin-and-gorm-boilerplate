package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/services"
	"github.com/huda7077/gin-and-gorm-boilerplate/models"
	"github.com/huda7077/gin-and-gorm-boilerplate/pkg/exceptions"
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
		c.Error(exceptions.NotFoundError{Message: "Products not found"})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Get products successfully",
		"data":    products,
	})
}

func (ctr *ProductController) Create(c *gin.Context) {
	var req *models.Product

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(exceptions.ValidationError{Message: err.Error()})
		return
	}
	if err := ctr.Service.Create(req.Name, req.Price, req.Stock); err != nil {
		c.Error(exceptions.ConflictError{Message: err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Create product successfully",
		"data": gin.H{
			"name":  req.Name,
			"price": req.Price,
			"stock": req.Stock,
		},
	})
}
