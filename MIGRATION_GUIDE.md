# 🔧 Migration Guide - How to Update Your Existing Code

## 🎯 Overview

Panduan ini menjelaskan langkah-langkah untuk memigrate kode existing Anda ke struktur yang sudah di-refactor.

## 📋 Pre-Migration Checklist

- [ ] Backup database Anda
- [ ] Backup semua kode existing
- [ ] Commit semua perubahan ke git
- [ ] Dokumentasikan semua endpoint yang sudah ada
- [ ] List semua dependencies yang digunakan

## 🚀 Step-by-Step Migration

### Step 1: Update Dependencies

Pastikan Anda memiliki dependencies yang diperlukan:

```bash
# Core dependencies
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres  # atau driver database lain
go get -u github.com/joho/godotenv
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v5

# Validation
go get -u github.com/go-playground/validator/v10

# Clean up
go mod tidy
```

### Step 2: Update Exception Handling

1. **Backup old exception files:**
```bash
mkdir backup
cp pkg/exceptions/* backup/
```

2. **Replace dengan file baru:**
   - ✅ `pkg/exceptions/app_error.go` (NEW)
   - ✅ `pkg/exceptions/error_handler.go` (UPDATED)

3. **Update semua tempat yang menggunakan old error types:**

**Sebelum:**
```go
return exceptions.NotFoundError{Message: "user not found"}
return exceptions.UnauthorizedError{Message: "invalid credentials"}
return exceptions.ConflictError{Message: "email exists"}
```

**Sesudah:**
```go
return exceptions.NewNotFoundError("user not found")
return exceptions.NewUnauthorizedError("invalid credentials")
return exceptions.NewConflictError("email exists")
```

### Step 3: Update Main.go

Replace `cmd/server/main.go` dengan versi baru yang include dependency injection:

```go
func main() {
    // Load configuration
    config := configs.New()

    // Connect to database
    configs.ConnectDatabase()

    // Initialize repositories
    repos := repositories.NewRepositories(configs.DB)

    // Setup router with dependencies
    r := routes.SetupRouter(repos, config)

    // Health check
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
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
```

### Step 4: Update Routes Setup

**Update `internal/routes/routes.go`:**

```go
func SetupRouter(repos *repositories.Repositories, config configs.Config) *gin.Engine {
    r := gin.Default()
    r.HandleMethodNotAllowed = true
    r.Use(middlewares.ErrorMiddleware())

    api := r.Group("/api")
    {
        SetupAuthRoutes(api, repos, config)
        // Add more routes here
    }

    return r
}
```

### Step 5: Update Each Feature

Untuk setiap feature yang ada (misal: Product, User, Order), lakukan:

#### 5.1 Update Repository

**Sebelum:**
```go
func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepositoryImpl{DB: db}
}
```

**Sesudah:**
```go
func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepositoryImpl{db: db}  // lowercase db
}

// Update all error handling
if errors.Is(err, gorm.ErrRecordNotFound) {
    return models.User{}, exceptions.NewNotFoundError("user not found")
}
```

#### 5.2 Update Service

**Sebelum:**
```go
type authServiceImpl struct {
    repositories repositories.Repositories  // not pointer
    config       configs.Config
}

func NewAuthService() AuthService {
    return &authServiceImpl{}  // no params
}
```

**Sesudah:**
```go
type authServiceImpl struct {
    repositories *repositories.Repositories  // pointer
    config       configs.Config
}

func NewAuthService(repositories *repositories.Repositories, config configs.Config) AuthService {
    return &authServiceImpl{
        repositories: repositories,
        config:       config,
    }
}
```

**Update service methods:**
- Return DTOs instead of models
- Use proper error types
- Add context parameter if missing

#### 5.3 Update Controller

**Update response handling:**

**Sebelum:**
```go
c.JSON(http.StatusOK, gin.H{
    "code":    http.StatusOK,
    "message": "Success",
    "data":    result,
})
```

**Sesudah:**
```go
exceptions.SuccessResponse(c, http.StatusOK, "success", result)
```

**Update error handling:**

**Sebelum:**
```go
if err != nil {
    c.JSON(500, gin.H{
        "success": false,
        "message": err.Error(),
    })
    return
}
```

**Sesudah:**
```go
if err != nil {
    c.Error(err)  // Let middleware handle it
    return
}
```

#### 5.4 Update Routes

**Sebelum:**
```go
func ProductRouter(rg *gin.RouterGroup) {
    productService := services.NewProductService()
    productController := controllers.NewProductController(productService)
    // ...
}
```

**Sesudah:**
```go
func SetupProductRoutes(rg *gin.RouterGroup, repos *repositories.Repositories, config configs.Config) {
    productService := services.NewProductService(repos)
    productController := controllers.NewProductController(productService)
    // ...
}
```

#### 5.5 Update DTOs

Tambahkan response DTOs jika belum ada:

```go
type ProductResponse struct {
    ID        uint    `json:"id"`
    Name      string  `json:"name"`
    Price     float64 `json:"price"`
    Stock     int     `json:"stock"`
    CreatedAt string  `json:"createdAt"`
    UpdatedAt string  `json:"updatedAt"`
}
```

### Step 6: Update Repositories.go

Add all your repositories to the struct:

```go
type Repositories struct {
    DB      *gorm.DB
    User    UserRepository
    Product ProductRepository
    Order   OrderRepository
    // ... add more
}

func NewRepositories(db *gorm.DB) *Repositories {
    return &Repositories{
        DB:      db,
        User:    NewUserRepository(db),
        Product: NewProductRepository(db),
        Order:   NewOrderRepository(db),
        // ... initialize all
    }
}
```

## 🧪 Testing Migration

### Test 1: Start Server
```bash
go run cmd/server/main.go
```

Pastikan server start tanpa error.

### Test 2: Test Health Check
```bash
curl http://localhost:8080/ping
```

Expected response:
```json
{
  "message": "pong"
}
```

### Test 3: Test Auth Endpoints

**Register:**
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123"
  }'
```

Expected response:
```json
{
  "success": true,
  "message": "registration successful",
  "data": null
}
```

**Login:**
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

Expected response:
```json
{
  "success": true,
  "message": "login successful",
  "data": {
    "user": {
      "id": 1,
      "name": "Test User",
      "email": "test@example.com",
      "role": "USER"
    },
    "token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

### Test 4: Test Error Responses

**Invalid Email:**
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "invalid-email",
    "password": "password123"
  }'
```

Expected response:
```json
{
  "success": false,
  "message": "Validation Error",
  "data": [
    {
      "field": "Email",
      "message": "Email must be a valid email"
    }
  ]
}
```

**Wrong Password:**
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "wrongpassword"
  }'
```

Expected response:
```json
{
  "success": false,
  "message": "invalid email or password"
}
```

## 🐛 Common Issues & Solutions

### Issue 1: Import Errors

**Error:**
```
cannot find package "github.com/go-playground/validator/v10"
```

**Solution:**
```bash
go get github.com/go-playground/validator/v10
go mod tidy
```

### Issue 2: Circular Dependency

**Error:**
```
import cycle not allowed
```

**Solution:**
- Make sure models don't import from internal packages
- Make sure DTOs are separate from models
- Check import paths are correct

### Issue 3: Interface Not Implemented

**Error:**
```
*productServiceImpl does not implement ProductService
```

**Solution:**
- Check all interface methods are implemented
- Check method signatures match exactly
- Check return types (pointer vs value)

### Issue 4: Nil Pointer Dereference

**Error:**
```
panic: runtime error: invalid memory address
```

**Solution:**
- Make sure repositories are initialized in NewRepositories()
- Check dependency injection chain
- Add nil checks where needed

### Issue 5: Validation Not Working

**Error:**
Validation errors not showing properly

**Solution:**
- Make sure gin validator is set up
- Check binding tags in DTOs
- Verify error middleware is registered

## 📝 Migration Checklist

### Core Files
- [ ] `pkg/exceptions/app_error.go` - Created
- [ ] `pkg/exceptions/error_handler.go` - Updated
- [ ] `cmd/server/main.go` - Updated
- [ ] `internal/routes/routes.go` - Updated

### For Each Feature
- [ ] Model - Reviewed (no changes needed usually)
- [ ] DTO - Added response DTOs
- [ ] Repository - Updated error handling
- [ ] Service - Updated DI & return DTOs
- [ ] Controller - Updated response handling
- [ ] Routes - Updated with DI

### Testing
- [ ] Server starts successfully
- [ ] Health check works
- [ ] All endpoints return correct format
- [ ] Error handling works
- [ ] Validation works
- [ ] Authentication works (if applicable)

## 🎯 Post-Migration Tasks

1. **Remove old code:**
```bash
# Remove old exception files if everything works
rm backup/exceptions/*.go  # after confirming everything works
```

2. **Update documentation:**
   - Update API documentation
   - Update README
   - Document new response format

3. **Add tests:**
   - Unit tests for services
   - Integration tests for endpoints
   - Test error scenarios

4. **Code review:**
   - Review all changes
   - Check for consistency
   - Ensure best practices followed

## 💡 Tips

1. **Migrate one feature at a time** - Don't try to migrate everything at once
2. **Keep backup** - Don't delete old code until new code is tested
3. **Test thoroughly** - Test both success and error scenarios
4. **Document changes** - Keep notes of what you changed
5. **Use git** - Commit after each successful migration step

## 🆘 Need Help?

If you encounter issues:

1. Check error messages carefully
2. Review the examples in `REFACTORING_GUIDE.md`
3. Look at the working auth implementation
4. Check `QUICK_REFERENCE.md` for patterns
5. Make sure all imports are correct
6. Verify dependency injection chain

## 📚 References

- `REFACTORING_GUIDE.md` - Detailed guide with examples
- `REFACTORING_SUMMARY.md` - Summary of changes
- `QUICK_REFERENCE.md` - Quick reference for common patterns
- Working code in `internal/` - Reference implementation

---

**Good luck with your migration! 🚀**

Remember: Take it slow, test often, and don't hesitate to refer back to the documentation.
