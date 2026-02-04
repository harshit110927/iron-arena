# API Documentation

## Base URL

```
Development: http://localhost:8080
Production: TBD
```

## Authentication

All authenticated endpoints require a JWT token in the Authorization header:

```
Authorization: Bearer <token>
```

## Endpoints

### Health Check

**GET** `/health`

Returns the health status of the API and database connection.

**Response**
```json
{
  "status": "ok",
  "timestamp": "2024-01-01T00:00:00Z",
  "database": "healthy"
}
```

### Authentication

#### Register

**POST** `/api/auth/register`

Register a new user account.

**Request Body**
```json
{
  "email": "user@example.com",
  "password": "securepassword",
  "username": "johndoe"
}
```

**Response**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_at": "2024-01-02T00:00:00Z",
  "user": {
    "id": "user-123",
    "email": "user@example.com",
    "username": "johndoe",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### Login

**POST** `/api/auth/login`

Authenticate and receive a JWT token.

**Request Body**
```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

**Response**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_at": "2024-01-02T00:00:00Z",
  "user": {
    "id": "user-123",
    "email": "user@example.com",
    "username": "johndoe",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### Protected Routes

#### Get Protected Resource

**GET** `/api/protected`

Example protected endpoint that requires authentication.

**Headers**
```
Authorization: Bearer <token>
```

**Response**
```json
{
  "message": "This is a protected route"
}
```

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid request body"
}
```

### 401 Unauthorized
```json
{
  "error": "Invalid or expired token"
}
```

### 500 Internal Server Error
```json
{
  "error": "Internal server error"
}
```

## Rate Limiting

TBD

## Versioning

TBD
