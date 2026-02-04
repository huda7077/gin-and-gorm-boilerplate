# Gin and GORM Boilerplate - Refactored

## 📋 Struktur Proyek

```
gin-and-gorm-boilerplate/
├── cmd/
│   └── server/
│       └── main.go                 # Entry point aplikasi
├── configs/
│   ├── config.go                   # Configuration management
│   └── database.go                 # Database connection
├── internal/
│   ├── controllers/                # HTTP handlers
│   │   └── auth.controller.go
│   ├── dto/                        # Data Transfer Objects
│   │   └── auth_request.go
│   ├── middlewares/                # Custom middlewares
│   │   ├── error.middleware.go
│   │   └── jwt_auth.middleware.go
│   ├── repositories/               # Database operations
│   │   ├── repositories.go
│   │   └── user.repository.go
│   ├── routes/                     # Route definitions
│   │   ├── routes.go
│   │   └── auth.route.go
│   └── services/                   # Business logic
│       └── auth.service.go
├── models/                         # Database models
│   └── user.model.go
└── pkg/
    ├── exceptions/                 # Custom error handling
    │   ├── app_error.go
    │   └── error_handler.go
    └── helpers/                    # Utility functions
        └── generate_jwt.helper.go
```

## 🎯 Best Practices yang Diterapkan

### 1. **Dependency Injection**
- Dependencies dipass dari layer teratas (main.go) ke bawah
- Tidak ada global variables untuk service/repository
- Mudah untuk testing dan maintainability

### 2. **Standardized Error Handling**
- Error types yang konsisten menggunakan `AppError`
- Centralized error handler di middleware
- Response format yang uniform

### 3. **Separation of Concerns**
- **Controller**: Handle HTTP request/response
- **Service**: Business logic
- **Repository**: Database operations
- **DTO**: Data validation dan transfer

### 4. **Clean Code Principles**
- Interface-based design
- Single Responsibility Principle
- Dependency Inversion Principle

## 🚀 Cara Menggunakan

### 1. Membuat Feature Baru (Contoh: Product)

#### Step 1: Buat Model
```go
// models/product.model.go
package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	Stock       int       `gorm:"not null;default:0" json:"stock"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (p *Product) TableName() string {
	return "products"
}
```

#### Step 2: Buat DTO

**Create two files for better organization:**

```go
// internal/dto/product.request.dto.go
package dto

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required,min=3,max=255"`
	Description string  `json:"description" binding:"max=1000"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"required,gte=0"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" binding:"omitempty,min=3,max=255"`
	Description string  `json:"description" binding:"omitempty,max=1000"`
	Price       float64 `json:"price" binding:"omitempty,gt=0"`
	Stock       int     `json:"stock" binding:"omitempty,gte=0"`
}
```

```go
// internal/dto/product.response.dto.go
package dto

type ProductResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}
```

#### Step 3: Buat Repository
```go
// internal/repositories/product.repository.go
package repositories

import (
	"context"
	"errors"

	"github.com/huda7077/gin-and-gorm-boilerplate/models"
	"github.com/huda7077/gin-and-gorm-boilerplate/pkg/exceptions"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]models.Product, error)
	FindById(ctx context.Context, id uint) (models.Product, error)
	Create(ctx context.Context, product models.Product) (models.Product, error)
	Update(ctx context.Context, id uint, product models.Product) (models.Product, error)
	Delete(ctx context.Context, id uint) error
}

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{db: db}
}

func (r *productRepositoryImpl) FindAll(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	if err := r.db.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepositoryImpl) FindById(ctx context.Context, id uint) (models.Product, error) {
	var product models.Product
	if err := r.db.WithContext(ctx).First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Product{}, exceptions.NewNotFoundError("product not found")
		}
		return models.Product{}, err
	}
	return product, nil
}

func (r *productRepositoryImpl) Create(ctx context.Context, product models.Product) (models.Product, error) {
	if err := r.db.WithContext(ctx).Create(&product).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (r *productRepositoryImpl) Update(ctx context.Context, id uint, product models.Product) (models.Product, error) {
	existing, err := r.FindById(ctx, id)
	if err != nil {
		return models.Product{}, err
	}

	if err := r.db.WithContext(ctx).Model(&existing).Updates(product).Error; err != nil {
		return models.Product{}, err
	}

	return existing, nil
}

func (r *productRepositoryImpl) Delete(ctx context.Context, id uint) error {
	if _, err := r.FindById(ctx, id); err != nil {
		return err
	}

	if err := r.db.WithContext(ctx).Delete(&models.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}

// Jangan lupa update repositories.go
// Add this to Repositories struct:
// Product ProductRepository

// Add this to NewRepositories:
// Product: NewProductRepository(db),

// Add this to WithTx:
// Product: NewProductRepository(tx),
```

#### Step 4: Buat Service
```go
// internal/services/product.service.go
package services

import (
	"context"
	"time"

	"github.com/huda7077/gin-and-gorm-boilerplate/internal/dto"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/repositories"
	"github.com/huda7077/gin-and-gorm-boilerplate/models"
	"github.com/huda7077/gin-and-gorm-boilerplate/pkg/exceptions"
)

type ProductService interface {
	GetAll(ctx context.Context) ([]dto.ProductResponse, error)
	GetById(ctx context.Context, id uint) (*dto.ProductResponse, error)
	Create(ctx context.Context, req dto.CreateProductRequest) (*dto.ProductResponse, error)
	Update(ctx context.Context, id uint, req dto.UpdateProductRequest) (*dto.ProductResponse, error)
	Delete(ctx context.Context, id uint) error
}

type productServiceImpl struct {
	repositories *repositories.Repositories
}

func NewProductService(repositories *repositories.Repositories) ProductService {
	return &productServiceImpl{
		repositories: repositories,
	}
}

func (s *productServiceImpl) GetAll(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := s.repositories.Product.FindAll(ctx)
	if err != nil {
		return nil, exceptions.NewInternalServerError("failed to get products")
	}

	responses := make([]dto.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = dto.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
		}
	}

	return responses, nil
}

func (s *productServiceImpl) GetById(ctx context.Context, id uint) (*dto.ProductResponse, error) {
	product, err := s.repositories.Product.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	response := &dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}

	return response, nil
}

func (s *productServiceImpl) Create(ctx context.Context, req dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	created, err := s.repositories.Product.Create(ctx, product)
	if err != nil {
		return nil, exceptions.NewInternalServerError("failed to create product")
	}

	response := &dto.ProductResponse{
		ID:          created.ID,
		Name:        created.Name,
		Description: created.Description,
		Price:       created.Price,
		Stock:       created.Stock,
		CreatedAt:   created.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   created.UpdatedAt.Format(time.RFC3339),
	}

	return response, nil
}

func (s *productServiceImpl) Update(ctx context.Context, id uint, req dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	updated, err := s.repositories.Product.Update(ctx, id, product)
	if err != nil {
		return nil, err
	}

	response := &dto.ProductResponse{
		ID:          updated.ID,
		Name:        updated.Name,
		Description: updated.Description,
		Price:       updated.Price,
		Stock:       updated.Stock,
		CreatedAt:   updated.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   updated.UpdatedAt.Format(time.RFC3339),
	}

	return response, nil
}

func (s *productServiceImpl) Delete(ctx context.Context, id uint) error {
	if err := s.repositories.Product.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
```

#### Step 5: Buat Controller
```go
// internal/controllers/product.controller.go
package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/dto"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/services"
	"github.com/huda7077/gin-and-gorm-boilerplate/pkg/exceptions"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (ctrl *ProductController) GetAll(c *gin.Context) {
	products, err := ctrl.productService.GetAll(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	exceptions.SuccessResponse(c, http.StatusOK, "products retrieved successfully", products)
}

func (ctrl *ProductController) GetById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(exceptions.NewBadRequestError("invalid product id", nil))
		return
	}

	product, err := ctrl.productService.GetById(c.Request.Context(), uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	exceptions.SuccessResponse(c, http.StatusOK, "product retrieved successfully", product)
}

func (ctrl *ProductController) Create(c *gin.Context) {
	var reqBody dto.CreateProductRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Error(err)
		return
	}

	product, err := ctrl.productService.Create(c.Request.Context(), reqBody)
	if err != nil {
		c.Error(err)
		return
	}

	exceptions.SuccessResponse(c, http.StatusCreated, "product created successfully", product)
}

func (ctrl *ProductController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(exceptions.NewBadRequestError("invalid product id", nil))
		return
	}

	var reqBody dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.Error(err)
		return
	}

	product, err := ctrl.productService.Update(c.Request.Context(), uint(id), reqBody)
	if err != nil {
		c.Error(err)
		return
	}

	exceptions.SuccessResponse(c, http.StatusOK, "product updated successfully", product)
}

func (ctrl *ProductController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(exceptions.NewBadRequestError("invalid product id", nil))
		return
	}

	if err := ctrl.productService.Delete(c.Request.Context(), uint(id)); err != nil {
		c.Error(err)
		return
	}

	exceptions.SuccessResponse(c, http.StatusOK, "product deleted successfully", nil)
}
```

#### Step 6: Buat Routes
```go
// internal/routes/product.route.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huda7077/gin-and-gorm-boilerplate/configs"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/controllers"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/middlewares"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/repositories"
	"github.com/huda7077/gin-and-gorm-boilerplate/internal/services"
)

func SetupProductRoutes(rg *gin.RouterGroup, repos *repositories.Repositories, config configs.Config) {
	// Initialize service
	productService := services.NewProductService(repos)

	// Initialize controller
	productController := controllers.NewProductController(productService)

	// Setup routes
	products := rg.Group("/products")
	{
		// Public routes
		products.GET("", productController.GetAll)
		products.GET("/:id", productController.GetById)

		// Protected routes (require authentication)
		products.Use(middlewares.JWTAuthMiddleware(config))
		{
			products.POST("", productController.Create)
			products.PUT("/:id", productController.Update)
			products.DELETE("/:id", productController.Delete)
		}
	}
}
```

#### Step 7: Register Routes di Main Router
```go
// internal/routes/routes.go
// Add this in SetupRouter function:
SetupProductRoutes(api, repos, config)
```

## 📝 Response Format

### Success Response
```json
{
  "success": true,
  "message": "login successful",
  "data": {
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "role": "USER"
    },
    "token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

### Error Response
```json
{
  "success": false,
  "message": "Validation Error",
  "data": [
    {
      "field": "Email",
      "message": "Email is required"
    }
  ]
}
```

## 🔐 Error Types

```go
// Bad Request (400)
exceptions.NewBadRequestError("invalid input", nil)

// Unauthorized (401)
exceptions.NewUnauthorizedError("invalid credentials")

// Forbidden (403)
exceptions.NewForbiddenError("access denied")

// Not Found (404)
exceptions.NewNotFoundError("user not found")

// Conflict (409)
exceptions.NewConflictError("email already exists")

// Internal Server Error (500)
exceptions.NewInternalServerError("something went wrong")

// Validation Error (400)
exceptions.NewValidationError("validation failed", details)
```

## 🧪 Testing Tips

1. Test setiap layer secara terpisah (unit test)
2. Gunakan mock untuk dependencies
3. Test edge cases dan error handling
4. Integration test untuk flow lengkap

## 🎨 Code Style

1. Gunakan nama yang deskriptif
2. Comment untuk exported functions
3. Keep functions small and focused
4. Follow Go conventions (gofmt, golint)

## 📚 Resources

- [Gin Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [Go Best Practices](https://golang.org/doc/effective_go.html)
