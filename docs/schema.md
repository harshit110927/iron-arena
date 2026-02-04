# Database Schema

## Overview

PostgreSQL database schema for Iron Arena application.

## Tables

### users

Stores user account information.

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);
```

**Columns:**
- `id`: Unique identifier (UUID)
- `email`: User's email address (unique)
- `password_hash`: Bcrypt hashed password
- `username`: User's display name (unique)
- `created_at`: Account creation timestamp
- `updated_at`: Last modification timestamp

### sessions (optional)

Stores active user sessions.

```sql
CREATE TABLE sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_sessions_user_id ON sessions(user_id);
CREATE INDEX idx_sessions_token_hash ON sessions(token_hash);
CREATE INDEX idx_sessions_expires_at ON sessions(expires_at);
```

**Columns:**
- `id`: Unique identifier (UUID)
- `user_id`: Reference to users table
- `token_hash`: Hashed JWT token for validation
- `expires_at`: Session expiration timestamp
- `created_at`: Session creation timestamp

## Migrations

Use a migration tool like [golang-migrate](https://github.com/golang-migrate/migrate) or [goose](https://github.com/pressly/goose) for managing database schema changes.

### Example Migration (Up)

```sql
-- migrations/001_create_users_table.up.sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);
```

### Example Migration (Down)

```sql
-- migrations/001_create_users_table.down.sql
DROP TABLE IF EXISTS users;
```

## SQLC Configuration

Example `sqlc.yaml` for code generation:

```yaml
version: "2"
sql:
  - schema: "migrations/*.sql"
    queries: "internal/db/queries/*.sql"
    engine: "postgresql"
    gen:
      go:
        package: "db"
        out: "internal/db"
        emit_json_tags: true
        emit_prepared_queries: false
        emit_interface: true
        emit_exact_table_names: false
```

## Queries

Example queries for SQLC:

```sql
-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (email, password_hash, username)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET username = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
```
