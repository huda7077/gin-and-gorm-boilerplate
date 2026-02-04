# ✅ Testing Checklist

## 🎯 Pre-Testing Setup

```bash
# 1. Pastikan dependencies terinstall
go mod tidy

# 2. Setup environment variables
cp .env.example .env  # jika ada
# Edit .env sesuai konfigurasi lokal

# 3. Setup database
# Create database jika belum ada
# Run migrations

# 4. Start server
go run cmd/server/main.go
# atau dengan hot reload:
air
```

## 📋 Basic Functionality Tests

### ✅ Server Health

**Test: Ping Endpoint**

```bash
curl http://localhost:8080/ping
```

**Expected Response:**

```json
{
  "message": "pong"
}
```

**Status:** [✅] Pass [ ] Fail

---

## 🔐 Authentication Tests

### ✅ Register New User

**Test: Valid Registration**

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

**Expected Response:**

```json
{
  "success": true,
  "message": "registration successful",
  "data": null
}
```

**Status:** [ ] Pass [ ] Fail

---

**Test: Duplicate Email**

```bash
# Register same email again
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "email": "john@example.com",
    "password": "password456"
  }'
```

**Expected Response:**

```json
{
  "success": false,
  "message": "email already registered"
}
```

**Status:** [ ] Pass [ ] Fail

---

**Test: Invalid Email Format**

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "invalid-email",
    "password": "password123"
  }'
```

**Expected Response:**

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

**Status:** [ ] Pass [ ] Fail

---

**Test: Password Too Short**

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "12345"
  }'
```

**Expected Response:**

```json
{
  "success": false,
  "message": "Validation Error",
  "data": [
    {
      "field": "Password",
      "message": "Password must be at least 6 characters"
    }
  ]
}
```

**Status:** [ ] Pass [ ] Fail

---

**Test: Missing Required Fields**

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User"
  }'
```

**Expected Response:**

```json
{
  "success": false,
  "message": "Validation Error",
  "data": [
    {
      "field": "Email",
      "message": "Email is required"
    },
    {
      "field": "Password",
      "message": "Password is required"
    }
  ]
}
```

**Status:** [ ] Pass [ ] Fail

---

### ✅ Login

**Test: Valid Login**

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

**Expected Response:**

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
      "createdAt": "2024-01-01T00:00:00Z",
      "updatedAt": "2024-01-01T00:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

**Status:** [ ] Pass [ ] Fail

---

**Test: Wrong Password**

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "wrongpassword"
  }'
```

**Expected Response:**

```json
{
  "success": false,
  "message": "invalid email or password"
}
```

**Status:** [ ] Pass [ ] Fail

---

**Test: Non-existent Email**

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "notexist@example.com",
    "password": "password123"
  }'
```

**Expected Response:**

```json
{
  "success": false,
  "message": "invalid email or password"
}
```

**Status:** [ ] Pass [ ] Fail

---

**Test: Invalid Email Format**

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "invalid-email",
    "password": "password123"
  }'
```

**Expected Response:**

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

**Status:** [ ] Pass [ ] Fail

---

## 🔍 Edge Cases Tests

### ✅ HTTP Method Tests

**Test: GET on POST endpoint**

```bash
curl -X GET http://localhost:8080/api/auth/login
```

**Expected:** 405 Method Not Allowed

**Status:** [ ] Pass [ ] Fail

---

**Test: Empty Body**

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{}'
```

**Expected:** Validation Error

**Status:** [ ] Pass [ ] Fail

---

**Test: Malformed JSON**

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{invalid json'
```

**Expected:** 400 Bad Request

**Status:** [ ] Pass [ ] Fail

---

### ✅ Security Tests

**Test: Password Not in Response**

```bash
# After login, check response
# Password field should NOT be present in user object
```

**Status:** [ ] Pass [ ] Fail

---

**Test: SQL Injection Attempt**

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com'\''; DROP TABLE users; --",
    "password": "password"
  }'
```

**Expected:** Should fail safely, no SQL error

**Status:** [ ] Pass [ ] Fail

---

**Test: XSS Attempt**

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "<script>alert(1)</script>",
    "email": "xss@example.com",
    "password": "password123"
  }'
```

**Expected:** Should save safely (HTML escaped)

**Status:** [ ] Pass [ ] Fail

---

## 📊 Response Format Tests

### ✅ Success Response Structure

**Verify all success responses have:**

- [ ] `success: true`
- [ ] `message: string`
- [ ] `data: object/array/null`

---

### ✅ Error Response Structure

**Verify all error responses have:**

- [ ] `success: false`
- [ ] `message: string`
- [ ] `data: object/null` (optional error details)

---

## 🎯 Status Code Tests

**Verify correct HTTP status codes:**

| Scenario                       | Expected Status           | Status |
| ------------------------------ | ------------------------- | ------ |
| Successful GET                 | 200 OK                    | [ ]    |
| Successful POST (create)       | 201 Created               | [ ]    |
| Successful PUT/PATCH           | 200 OK                    | [ ]    |
| Successful DELETE              | 200 OK                    | [ ]    |
| Validation Error               | 400 Bad Request           | [ ]    |
| Unauthorized (no token)        | 401 Unauthorized          | [ ]    |
| Invalid Credentials            | 401 Unauthorized          | [ ]    |
| Forbidden (insufficient perms) | 403 Forbidden             | [ ]    |
| Resource Not Found             | 404 Not Found             | [ ]    |
| Conflict (duplicate)           | 409 Conflict              | [ ]    |
| Server Error                   | 500 Internal Server Error | [ ]    |

---

## 🧪 Database Tests

### ✅ Data Persistence

**Test: User Persisted**

```sql
-- Check database directly
SELECT * FROM users WHERE email = 'john@example.com';
```

**Verify:**

- [ ] User exists
- [ ] Password is hashed
- [ ] Timestamps are correct
- [ ] Default role is set

---

**Test: No Duplicate Emails**

```sql
-- Check for duplicates
SELECT email, COUNT(*) as count
FROM users
GROUP BY email
HAVING count > 1;
```

**Expected:** No results

**Status:** [ ] Pass [ ] Fail

---

## 🔐 JWT Token Tests

### ✅ Token Generation

**Test: Token is Valid JWT**

```bash
# Save token from login response
TOKEN="eyJhbGciOiJIUzI1NiIs..."

# Decode at https://jwt.io or use jwt-cli
```

**Verify token contains:**

- [ ] User ID
- [ ] Email
- [ ] Role
- [ ] Expiration (72 hours from now)

---

**Test: Token Works for Protected Routes**

```bash
# If you have protected routes:
curl http://localhost:8080/api/protected \
  -H "Authorization: Bearer $TOKEN"
```

**Expected:** 200 OK with data

**Status:** [ ] Pass [ ] Fail

---

**Test: Invalid Token Rejected**

```bash
curl http://localhost:8080/api/protected \
  -H "Authorization: Bearer invalid-token"
```

**Expected:** 401 Unauthorized

**Status:** [ ] Pass [ ] Fail

---

**Test: No Token Rejected**

```bash
curl http://localhost:8080/api/protected
```

**Expected:** 401 Unauthorized

**Status:** [ ] Pass [ ] Fail

---

## 📈 Performance Tests

### ✅ Response Time

**Test: Response under 100ms (local)**

```bash
# Use curl with timing
curl -w "@curl-format.txt" -o /dev/null -s http://localhost:8080/api/auth/login \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"password123"}'
```

**Create curl-format.txt:**

```
time_namelookup:  %{time_namelookup}s\n
time_connect:     %{time_connect}s\n
time_total:       %{time_total}s\n
```

**Status:** [ ] Pass [ ] Fail

---

### ✅ Concurrent Requests

**Test: Handle 100 concurrent logins**

```bash
# Use Apache Bench or similar
ab -n 100 -c 10 -p login.json -T 'application/json' \
  http://localhost:8080/api/auth/login
```

**Expected:** All requests succeed

**Status:** [ ] Pass [ ] Fail

---

## 📱 Integration Tests

### ✅ Full User Flow

**Test: Complete Registration → Login Flow**

1. [ ] Register new user
2. [ ] Login with credentials
3. [ ] Receive valid token
4. [ ] Token contains correct user info
5. [ ] Can access protected resources with token

---

## 🐛 Error Handling Tests

### ✅ Database Errors

**Test: Database Connection Lost**

```bash
# Stop database temporarily
# Try to make request
```

**Expected:** 500 Internal Server Error with generic message

**Status:** [ ] Pass [ ] Fail

---

### ✅ Application Errors

**Test: Panic Recovery**

```bash
# Create endpoint that panics
# Should be caught by recovery middleware
```

**Expected:** 500 Internal Server Error, no crash

**Status:** [ ] Pass [ ] Fail

---

## 📝 Documentation Tests

### ✅ API Documentation

- [ ] All endpoints documented
- [ ] Request examples provided
- [ ] Response examples provided
- [ ] Error responses documented
- [ ] Authentication requirements clear

---

## ✅ Final Checklist

### Code Quality

- [ ] No commented out code
- [ ] No console.logs or debug prints
- [ ] Proper error messages (no technical details leaked)
- [ ] Code follows project conventions
- [ ] All TODOs addressed

### Security

- [ ] Passwords are hashed
- [ ] No sensitive data in logs
- [ ] No sensitive data in responses
- [ ] Input validation working
- [ ] SQL injection protected
- [ ] XSS protected

### Functionality

- [ ] All endpoints work
- [ ] All validations work
- [ ] All error cases handled
- [ ] Database operations work
- [ ] Transactions work (if applicable)

### Performance

- [ ] Responses are fast
- [ ] No N+1 queries
- [ ] Database indexes used
- [ ] Connection pooling configured

### Testing

- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Edge cases tested
- [ ] Security tests pass

---

## 📊 Test Summary

**Date:** ******\_\_\_\_******

**Tester:** ******\_\_\_\_******

**Total Tests:** **\_\_**

**Passed:** **\_\_** **Failed:** **\_\_**

**Pass Rate:** **\_\_**%

---

## 🐛 Issues Found

| #   | Test | Issue | Severity | Status |
| --- | ---- | ----- | -------- | ------ |
| 1   |      |       |          |        |
| 2   |      |       |          |        |
| 3   |      |       |          |        |

---

## 📝 Notes

```
Add any additional notes, observations, or comments here:




```

---

## ✅ Sign-off

- [ ] All critical tests passed
- [ ] All issues documented
- [ ] Ready for production / next phase

**Signed:** ******\_\_\_\_****** **Date:** ******\_\_\_\_******
