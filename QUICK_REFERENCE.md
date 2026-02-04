# 🚀 Quick Reference - Gin GORM Boilerplate

## 📝 Checklist untuk Feature Baru

```
[ ] 1. Buat Model di models/
[ ] 2. Buat DTO Request & Response di internal/dto/
[ ] 3. Buat Repository Interface & Implementation di internal/repositories/
[ ] 4. Update repositories.go - tambahkan ke struct & NewRepositories()
[ ] 5. Buat Service Interface & Implementation di internal/services/
[ ] 6. Buat Controller di internal/controllers/
[ ] 7. Buat Route Setup di internal/routes/
[ ] 8. Register Route di routes.go
[ ] 9. Test endpoints
```

## 🎯 Code Snippets

### 1. Error Handling
```go
// In Service or Repository
return nil, exceptions.NewNotFoundError("user not found")
return nil, exceptions.NewBadRequestError("invalid input", validationDetails)
return nil, exceptions.NewUnauthorizedError("invalid credentials")
return nil, exceptions.NewConflictError("email already exists")
return nil, exceptions.NewInternalServerError("database error")

// In Controller
if err != nil {
    c.Error(err)  // Let middleware handle it
    return
}

// Success Response
exceptions.SuccessResponse(c, http.StatusOK, "success message", data)
```

### 2. Repository Pattern
```go
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
```

### 3. Service Pattern
```go
type ProductService interface {
    GetAll(ctx context.Context) ([]dto.ProductResponse, error)
    GetById(ctx context.Context, id uint) (*dto.ProductResponse, error)
    Create(ctx context.Context, req dto.CreateProductRequest) (*dto.ProductResponse, error)
}

type productServiceImpl struct {
    repositories *repositories.Repositories
}

func NewProductService(repositories *repositories.Repositories) ProductService {
    return &productServiceImpl{repositories: repositories}
}
```

### 4. Controller Pattern
```go
type ProductController struct {
    productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
    return &ProductController{productService: productService}
}

func (ctrl *ProductController) GetAll(c *gin.Context) {
    result, err := ctrl.productService.GetAll(c.Request.Context())
    if err != nil {
        c.Error(err)
        return
    }
    exceptions.SuccessResponse(c, http.StatusOK, "success", result)
}
```

### 5. Route Setup
```go
func SetupProductRoutes(rg *gin.RouterGroup, repos *repositories.Repositories, config configs.Config) {
    productService := services.NewProductService(repos)
    productController := controllers.NewProductController(productService)
    
    products := rg.Group("/products")
    {
        products.GET("", productController.GetAll)
        products.GET("/:id", productController.GetById)
        products.POST("", productController.Create)
        products.PUT("/:id", productController.Update)
        products.DELETE("/:id", productController.Delete)
    }
}
```

### 6. DTO with Validation
```go
type CreateProductRequest struct {
    Name  string  `json:"name" binding:"required,min=3,max=100"`
    Price float64 `json:"price" binding:"required,gt=0"`
    Stock int     `json:"stock" binding:"required,gte=0"`
}

type ProductResponse struct {
    ID        uint    `json:"id"`
    Name      string  `json:"name"`
    Price     float64 `json:"price"`
    Stock     int     `json:"stock"`
    CreatedAt string  `json:"createdAt"`
}
```

## 🔐 Common Validation Tags

```go
binding:"required"              // Field is required
binding:"email"                 // Valid email format
binding:"min=6"                 // Minimum length
binding:"max=100"               // Maximum length
binding:"gt=0"                  // Greater than 0
binding:"gte=0"                 // Greater than or equal 0
binding:"lt=100"                // Less than 100
binding:"lte=100"               // Less than or equal 100
binding:"oneof=red green"       // One of specified values
binding:"alphanum"              // Alphanumeric
binding:"url"                   // Valid URL
binding:"required,min=6,max=20" // Multiple validations
```

## 🌐 HTTP Methods & Status Codes

```go
// Methods
GET    - Retrieve data
POST   - Create new resource
PUT    - Update entire resource
PATCH  - Update partial resource
DELETE - Delete resource

// Success Status Codes
http.StatusOK                  // 200 - Success
http.StatusCreated             // 201 - Resource created
http.StatusNoContent           // 204 - Success with no content

// Error Status Codes
http.StatusBadRequest          // 400 - Bad request
http.StatusUnauthorized        // 401 - Not authenticated
http.StatusForbidden           // 403 - Not authorized
http.StatusNotFound            // 404 - Resource not found
http.StatusConflict            // 409 - Resource conflict
http.StatusInternalServerError // 500 - Server error
```

## 📦 Common GORM Operations

```go
// Create
db.Create(&user)

// Find One
db.First(&user, id)
db.Where("email = ?", email).First(&user)

// Find Many
db.Find(&users)
db.Where("status = ?", "active").Find(&users)

// Update
db.Model(&user).Updates(updatedUser)
db.Model(&user).Update("name", "John")

// Delete
db.Delete(&user, id)

// With Context
db.WithContext(ctx).Create(&user)

// Pagination
db.Limit(10).Offset(20).Find(&users)

// Count
db.Model(&User{}).Count(&total)

// Transaction
db.Transaction(func(tx *gorm.DB) error {
    // operations
    return nil
})
```

## 🔄 Converting Between Model & DTO

```go
// Model to DTO (in Service)
response := &dto.UserResponse{
    ID:        user.ID,
    Name:      user.Name,
    Email:     user.Email,
    CreatedAt: user.CreatedAt.Format(time.RFC3339),
}

// DTO to Model (in Service)
user := models.User{
    Name:     req.Name,
    Email:    req.Email,
    Password: hashedPassword,
}
```

## 🧪 Testing Helper

```go
// Mock Repository
type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) FindById(ctx context.Context, id uint) (models.User, error) {
    args := m.Called(ctx, id)
    return args.Get(0).(models.User), args.Error(1)
}

// Test Service
func TestLogin(t *testing.T) {
    mockRepo := new(MockUserRepository)
    mockRepo.On("FindByEmail", mock.Anything, "test@test.com").Return(mockUser, nil)
    
    repos := &repositories.Repositories{User: mockRepo}
    service := services.NewAuthService(repos, mockConfig)
    
    result, err := service.Login(context.Background(), loginRequest)
    assert.NoError(t, err)
    assert.NotNil(t, result)
}
```

## 📁 File Naming Convention

```
models/user.model.go                      # Database model
internal/dto/user.request.dto.go          # Request DTOs
internal/dto/user.response.dto.go         # Response DTOs
internal/repositories/user.repository.go  # Repository
internal/services/user.service.go         # Service
internal/controllers/user.controller.go   # Controller
internal/routes/user.route.go             # Routes
```

**DTO Naming Convention:**
- `feature.request.dto.go` - Untuk request DTOs
- `feature.response.dto.go` - Untuk response DTOs
- Atau jika simple bisa digabung: `feature.dto.go`

## 🎨 Code Organization Rules

1. **Models** - Only database structure, no business logic
2. **DTOs** - Request/Response structures with validation
3. **Repositories** - ONLY database operations, no business logic
4. **Services** - Business logic, orchestrate repositories
5. **Controllers** - Handle HTTP, minimal logic
6. **Routes** - Just wire everything together

## 💡 Best Practices

✅ **DO:**
- Use dependency injection
- Return DTOs from services, not models
- Use pointer receivers for methods
- Use context for cancellation
- Handle errors properly
- Use meaningful variable names
- Add comments for exported functions

❌ **DON'T:**
- Don't use global variables for dependencies
- Don't put business logic in controllers
- Don't expose models directly in responses
- Don't ignore errors
- Don't use panic for regular errors
- Don't hardcode configuration values

## 🔍 Debugging Tips

```go
// Print struct
fmt.Printf("%+v\n", user)

// Print JSON
json, _ := json.MarshalIndent(user, "", "  ")
fmt.Println(string(json))

// GORM Debug Mode
db.Debug().Create(&user)

// Log Error Details
log.Printf("Error: %v\n", err)
```

## 📚 Useful Commands

```bash
# Run server
go run cmd/server/main.go

# With hot reload (using air)
air

# Format code
go fmt ./...

# Run tests
go test ./...

# Run specific test
go test -v ./internal/services -run TestLogin

# Install dependency
go get github.com/package/name

# Update dependencies
go mod tidy

# Build
go build -o bin/server cmd/server/main.go
```

## 🌟 Example Request/Response

### Register
```bash
POST /api/auth/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}

Response: 201 Created
{
  "success": true,
  "message": "registration successful",
  "data": null
}
```

### Login
```bash
POST /api/auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}

Response: 200 OK
{
  "success": true,
  "message": "login successful",
  "data": {
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "role": "USER",
      "createdAt": "2024-01-01T00:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

---

**Need more help?** Check:
- `REFACTORING_GUIDE.md` - Detailed guide with examples
- `REFACTORING_SUMMARY.md` - Summary of changes
- Existing code in `internal/` for reference
