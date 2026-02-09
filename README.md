# 🚀 Gin & GORM Boilerplate

A production-ready, feature-rich REST API boilerplate built with **Gin Framework** and **GORM** following best practices and clean architecture principles.

[![Go Version](https://img.shields.io/badge/Go-1.25.4-blue.svg)](https://golang.org)
[![Gin](https://img.shields.io/badge/Gin-v1.11.0-green.svg)](https://gin-gonic.com/)
[![GORM](https://img.shields.io/badge/GORM-v1.31.1-red.svg)](https://gorm.io/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

---

## 📋 Table of Contents

- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Project Structure](#-project-structure)
- [Prerequisites](#-prerequisites)
- [Installation](#-installation)
- [Configuration](#-configuration)
- [Database Migration](#-database-migration)
- [Running the Application](#-running-the-application)
- [API Documentation](#-api-documentation)
- [Architecture Overview](#-architecture-overview)
- [Best Practices](#-best-practices)
- [Testing](#-testing)
- [Deployment](#-deployment)
- [Contributing](#-contributing)
- [License](#-license)

---

## ✨ Features

### Core Features

- ✅ **RESTful API** - Clean and well-structured REST endpoints
- ✅ **Clean Architecture** - Repository → Service → Controller pattern
- ✅ **Dependency Injection** - Proper DI throughout the application
- ✅ **JWT Authentication** - Secure token-based authentication
- ✅ **Email Verification** - OTP-based email verification system
- ✅ **Password Reset** - Secure password reset with OTP
- ✅ **Database Transactions** - ACID-compliant operations
- ✅ **Error Handling** - Centralized error handling with custom errors
- ✅ **Input Validation** - Request validation using binding tags
- ✅ **CORS Support** - Configurable CORS middleware
- ✅ **Environment Config** - Environment-based configuration
- ✅ **Hot Reload** - Development with Air for hot reload
- ✅ **Database Migrations** - Schema versioning with Atlas

### Security Features

- 🔐 Password hashing with bcrypt
- 🔐 JWT token-based authentication
- 🔐 Email verification required for login
- 🔐 OTP expiry (10 minutes)
- 🔐 One-time use verification codes
- 🔐 Generic error messages (security by obscurity)
- 🔐 CORS protection

### Developer Experience

- 📝 Comprehensive documentation
- 🎯 Consistent code structure
- 🔄 Hot reload in development
- 🧪 Easy testing setup
- 📊 Beautiful email templates
- 🎨 Clean and readable code

---

## 🛠 Tech Stack

### Backend

- **[Go](https://golang.org/)** (v1.25.4) - Programming language
- **[Gin](https://gin-gonic.com/)** (v1.11.0) - HTTP web framework
- **[GORM](https://gorm.io/)** (v1.31.1) - ORM library
- **[PostgreSQL](https://www.postgresql.org/)** - Primary database
- **[JWT](https://github.com/golang-jwt/jwt)** (v5.3.1) - JSON Web Tokens
- **[bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)** - Password hashing

### Tools & Libraries

- **[Air](https://github.com/cosmtrek/air)** - Hot reload for Go
- **[Atlas](https://atlasgo.io/)** - Database schema migration
- **[gomail](https://github.com/go-gomail/gomail)** - Email sending
- **[godotenv](https://github.com/joho/godotenv)** - Environment variables
- **[validator](https://github.com/go-playground/validator)** - Request validation

---

## 📁 Project Structure

```
gin-and-gorm-boilerplate/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── configs/
│   ├── config.go                   # Configuration loader
│   ├── database.go                 # Database connection
│   └── email.go                    # Email configuration
├── internal/
│   ├── controllers/                # HTTP handlers
│   │   ├── auth.controller.go
│   │   └── product.controller.go
│   ├── dto/                        # Data Transfer Objects
│   │   ├── auth.request.dto.go
│   │   ├── auth.response.dto.go
│   │   ├── product.request.dto.go
│   │   └── product.response.dto.go
│   ├── middlewares/                # HTTP middlewares
│   │   ├── auth.middleware.go
│   │   ├── cors.middleware.go
│   │   └── error.middleware.go
│   ├── provider/                   # External service providers
│   │   └── mail/
│   │       ├── mail.go
│   │       └── templates/          # Email templates
│   │           ├── verify-account.html
│   │           ├── reset-password.html
│   │           └── welcome.html
│   ├── repositories/               # Data access layer
│   │   ├── repositories.go
│   │   ├── user.repository.go
│   │   └── verification_code.repository.go
│   ├── routes/                     # Route definitions
│   │   ├── routes.go
│   │   ├── auth.route.go
│   │   ├── product.route.go
│   │   └── protected.route.go
│   ├── services/                   # Business logic layer
│   │   ├── auth.service.go
│   │   └── product.service.go
│   └── validators/                 # Custom validators
│       └── validators.go
├── migrations/                     # Database migrations
│   └── *.sql
├── models/                         # Database models
│   ├── auth.model.go
│   ├── user.model.go
│   ├── verification_code.model.go
│   └── product.model.go
├── pkg/
│   ├── exceptions/                 # Custom error types
│   │   ├── app_error.go
│   │   ├── bad_request.go
│   │   ├── conflict.go
│   │   ├── error_handler.go
│   │   ├── not_found.go
│   │   ├── unauthorized.go
│   │   └── validation.go
│   └── helpers/                    # Utility functions
│       ├── jwt.helper.go
│       └── otp.helper.go
├── .air.toml                       # Air configuration
├── .env                            # Environment variables
├── .env.example                    # Example environment variables
├── .gitignore
├── atlas.hcl                       # Atlas configuration
├── go.mod                          # Go module dependencies
├── go.sum
├── Makefile                        # Build commands
└── README.md                       # This file
```

---

## 📋 Prerequisites

Before you begin, ensure you have the following installed:

- **Go** (v1.25.4 or higher) - [Download](https://golang.org/dl/)
- **PostgreSQL** (v14 or higher) - [Download](https://www.postgresql.org/download/)
- **Air** (optional, for hot reload) - [Install](https://github.com/cosmtrek/air#installation)
- **Atlas** (optional, for migrations) - [Install](https://atlasgo.io/getting-started/)
- **Make** (optional, for Makefile commands)

---

## 🚀 Installation

### 1. Clone the Repository

```bash
git clone https://github.com/huda7077/gin-and-gorm-boilerplate.git
cd gin-and-gorm-boilerplate
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Install Development Tools (Optional)

```bash
# Install Air for hot reload
go install github.com/cosmtrek/air@latest

# Install Atlas for migrations
curl -sSf https://atlasgo.sh | sh
```

---

## ⚙️ Configuration

### 1. Create Environment File

```bash
cp .env.example .env
```

### 2. Configure Environment Variables

Edit `.env` file:

```env
# Application
APP_PORT=8080

# JWT
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production

# Database
DATABASE_URL=postgres://username:password@localhost:5432/dbname?sslmode=disable&search_path=public

# Email (Gmail Example)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your-email@gmail.com
SMTP_PASS=your-gmail-app-password
SMTP_FROM=your-email@gmail.com
```

### 3. Gmail Setup for Email

**For Gmail users:**

1. Enable 2-Factor Authentication in your Google Account
2. Generate App Password:
   - Go to: https://myaccount.google.com/apppasswords
   - Select "Mail" and your device
   - Copy the 16-character password
3. Use the app password in `SMTP_PASS`

**For other email providers:**

- Update `SMTP_HOST` and `SMTP_PORT` accordingly
- Use your email credentials

---

## 🗄️ Database Migration

### Using Atlas (Recommended)

**1. Create a new migration:**

```bash
atlas migrate diff <migration_name> \
  --env local \
  --to "file://migrations"
```

**2. Apply migrations:**

```bash
atlas migrate apply \
  --env local \
  --url "postgres://user:pass@localhost:5432/dbname?sslmode=disable"
```

**3. Check migration status:**

```bash
atlas migrate status \
  --env local \
  --url "postgres://user:pass@localhost:5432/dbname?sslmode=disable"
```

### Manual Migration

You can also run migrations manually:

```bash
psql -U username -d dbname -f migrations/20260205104634_init.sql
```

---

## 🏃 Running the Application

### Development Mode (with Hot Reload)

```bash
# Using Air
air

# Or using Makefile
make dev
```

### Production Mode

```bash
# Build the application
go build -o bin/server cmd/server/main.go

# Run the binary
./bin/server

# Or using Makefile
make build
make run
```

### Using Make Commands

```bash
# Run in development mode
make dev

# Build for production
make build

# Run built binary
make run

# Run tests
make test

# Clean build artifacts
make clean
```

The server will start on `http://localhost:8080` (or your configured `APP_PORT`)

---

## 📚 API Documentation

### Base URL

```
http://localhost:8080/api
```

### Authentication Endpoints

#### 1. Register User

```http
POST /api/auth/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:** `201 Created`

```json
{
  "success": true,
  "message": "registration successful, please check your email for verification code",
  "data": null
}
```

---

#### 2. Verify Email

```http
POST /api/auth/verify-email
Content-Type: application/json

{
  "email": "john@example.com",
  "otp": "123456"
}
```

**Response:** `200 OK`

```json
{
  "success": true,
  "message": "email verified successfully",
  "data": null
}
```

---

#### 3. Resend Verification Code

```http
POST /api/auth/resend-verification
Content-Type: application/json

{
  "email": "john@example.com"
}
```

**Response:** `200 OK`

```json
{
  "success": true,
  "message": "verification code sent successfully",
  "data": null
}
```

---

#### 4. Login

```http
POST /api/auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:** `200 OK`

```json
{
  "success": true,
  "message": "login successful",
  "data": {
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "role": "USER",
      "verifiedAt": "2024-01-01T00:00:00Z",
      "createdAt": "2024-01-01T00:00:00Z",
      "updatedAt": "2024-01-01T00:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

---

#### 5. Forgot Password

```http
POST /api/auth/forgot-password
Content-Type: application/json

{
  "email": "john@example.com"
}
```

**Response:** `200 OK`

```json
{
  "success": true,
  "message": "if email exists, reset code has been sent",
  "data": null
}
```

---

#### 6. Reset Password

```http
POST /api/auth/reset-password
Content-Type: application/json

{
  "email": "john@example.com",
  "otp": "789012",
  "newPassword": "newpassword123"
}
```

**Response:** `200 OK`

```json
{
  "success": true,
  "message": "password reset successful",
  "data": null
}
```

---

### Error Response Format

All errors follow this structure:

```json
{
  "success": false,
  "message": "error message here",
  "data": {
    "field": "Email",
    "message": "Email is required"
  }
}
```

### Common HTTP Status Codes

| Code | Description                           |
| ---- | ------------------------------------- |
| 200  | Success                               |
| 201  | Created                               |
| 400  | Bad Request (validation error)        |
| 401  | Unauthorized                          |
| 404  | Not Found                             |
| 409  | Conflict (e.g., email already exists) |
| 500  | Internal Server Error                 |

---

## 🏛️ Architecture Overview

### Clean Architecture Layers

```
┌─────────────────────────────────────────┐
│           HTTP Layer (Gin)              │
│         Controllers & Routes             │
└────────────────┬────────────────────────┘
                 │
┌────────────────▼────────────────────────┐
│         Business Logic Layer            │
│              Services                    │
└────────────────┬────────────────────────┘
                 │
┌────────────────▼────────────────────────┐
│         Data Access Layer               │
│           Repositories                   │
└────────────────┬────────────────────────┘
                 │
┌────────────────▼────────────────────────┐
│            Database (PostgreSQL)         │
└──────────────────────────────────────────┘
```

### Request Flow

```
1. HTTP Request
   ↓
2. Middleware (CORS, Error Handling, Auth)
   ↓
3. Controller (Input Validation)
   ↓
4. Service (Business Logic + Transaction)
   ↓
5. Repository (Database Operations)
   ↓
6. Database
   ↓
7. Response (JSON)
```

### Transaction Flow Example

```go
// Service Layer
func (s *service) Register(ctx context.Context, req dto.RegisterRequest) error {
    // Start transaction
    err := s.repositories.DB.Transaction(func(tx *gorm.DB) error {
        txRepos := s.repositories.WithTx(tx)

        // 1. Create user
        user, err := txRepos.User.Create(ctx, user)
        if err != nil {
            return err // Rollback
        }

        // 2. Create verification code
        _, err = txRepos.VerificationCode.Create(ctx, code)
        if err != nil {
            return err // Rollback
        }

        return nil // Commit
    })

    // 3. Send email (outside transaction)
    s.mailProvider.SendMail(...)

    return err
}
```

---

## 💡 Best Practices

### 1. **Dependency Injection**

All dependencies are injected through constructors:

```go
// Service
func NewAuthService(repos *repositories.Repositories, config configs.Config) AuthService {
    return &authServiceImpl{
        repositories: repos,
        config: config,
    }
}

// Controller
func NewAuthController(service services.AuthService) *AuthController {
    return &AuthController{
        authService: service,
    }
}
```

### 2. **Error Handling**

Use custom error types for consistent error responses:

```go
// Service
if existingUser.ID != 0 {
    return exceptions.NewConflictError("email already registered")
}

// Controller
if err := ctrl.authService.Register(c.Request.Context(), reqBody); err != nil {
    c.Error(err) // Middleware will handle formatting
    return
}
```

### 3. **Database Transactions**

Always use transactions for operations that modify multiple tables:

```go
err := s.repositories.DB.Transaction(func(tx *gorm.DB) error {
    txRepos := s.repositories.WithTx(tx)

    // Multiple operations...

    return nil // Commit or return error to rollback
})
```

### 4. **Email Sending**

Send emails outside transactions to avoid holding DB connections:

```go
// Inside transaction: Create user + verification code
err := s.repositories.DB.Transaction(func(tx *gorm.DB) error {
    // DB operations
})

// Outside transaction: Send email
s.mailProvider.SendMail(...)
```

### 5. **Security**

- Never expose internal errors to clients
- Use generic messages for security-sensitive operations
- Hash passwords with bcrypt
- Validate all inputs
- Use JWT with expiry
- Require email verification

---

## 🧪 Testing

### Run All Tests

```bash
go test ./...

# With coverage
go test -cover ./...

# With verbose output
go test -v ./...
```

### Test Structure

```
internal/
├── controllers/
│   └── auth.controller_test.go
├── services/
│   └── auth.service_test.go
└── repositories/
    └── user.repository_test.go
```

### Example Test

```go
func TestRegister(t *testing.T) {
    // Setup
    // ...

    // Execute
    err := service.Register(ctx, request)

    // Assert
    assert.NoError(t, err)
}
```

---

## 🚢 Deployment

### Docker Deployment

**1. Create Dockerfile:**

```dockerfile
FROM golang:1.25.4-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server cmd/server/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
COPY --from=builder /app/.env .
EXPOSE 8080
CMD ["./server"]
```

**2. Build and run:**

```bash
docker build -t gin-boilerplate .
docker run -p 8080:8080 gin-boilerplate
```

### Production Checklist

- [ ] Update `JWT_SECRET` to a strong random value
- [ ] Use production email service (SendGrid, AWS SES)
- [ ] Enable HTTPS/TLS
- [ ] Set up database backups
- [ ] Configure logging and monitoring
- [ ] Set up CI/CD pipeline
- [ ] Add rate limiting
- [ ] Configure firewall rules
- [ ] Review and update CORS settings
- [ ] Set up health check endpoints

---

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Coding Standards

- Follow Go best practices and conventions
- Write meaningful commit messages
- Add comments for complex logic
- Ensure tests pass before submitting PR
- Update documentation as needed

---

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 👨‍💻 Author

**Huda**

- GitHub: [@huda7077](https://github.com/huda7077)

---

## 🙏 Acknowledgments

- [Gin Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [Atlas](https://atlasgo.io/)
- Go Community

---

## 📞 Support

If you have any questions or need help, please:

- Open an issue on GitHub

---

**Happy Coding! 🚀**
