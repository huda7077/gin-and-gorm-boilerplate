package services

import (
	"github.com/huda7077/gin-and-gorm-boilerplate/configs"
	"github.com/huda7077/gin-and-gorm-boilerplate/models"
)

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (ps *ProductService) GetAll() ([]models.Product, error) {
	var products []models.Product

	if err := configs.DB.Find(&products); err != nil {
		return nil, err.Error
	}
	return products, nil

}

func (ps *ProductService) Create(name string, price int, stock *int) error {
	product := &models.Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}
	if err := configs.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}
