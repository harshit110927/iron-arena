# Backend Service

Go backend service using clean architecture with Chi router, PostgreSQL, and JWT authentication.

## Architecture

```
backend/
├── cmd/
│   └── server/          # Application entry point
├── internal/
│   ├── config/          # Configuration management
│   ├── handlers/        # HTTP handlers
│   ├── services/        # Business logic
│   ├── middleware/      # HTTP middleware (auth, logging, etc.)
│   ├── db/              # Database connection and queries
│   └── models/          # Data models
```

## Setup

1. Copy `.env.example` to `.env` and update the values
2. Install dependencies: `go mod download`
3. Run the server: `go run cmd/server/main.go`

## Environment Variables

- `PORT`: Server port (default: 8080)
- `DATABASE_URL`: PostgreSQL connection string
- `JWT_SECRET`: Secret key for JWT signing
- `JWT_EXPIRY_HOURS`: JWT token expiry in hours
- `ENVIRONMENT`: Application environment (development/production)

## API Endpoints

- `GET /health` - Health check endpoint
- `POST /api/auth/login` - User login
- `POST /api/auth/register` - User registration

## Development

Run with hot reload (requires air):
```bash
air
```

Run tests:
```bash
go test ./...
```
