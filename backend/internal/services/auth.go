package services

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/harshit110927/iron-arena/backend/internal/config"
	"github.com/harshit110927/iron-arena/backend/internal/db"
	"github.com/harshit110927/iron-arena/backend/internal/middleware"
	"github.com/harshit110927/iron-arena/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	cfg *config.Config
	db  *db.Database
}

func NewAuthService(cfg *config.Config, database *db.Database) *AuthService {
	return &AuthService{
		cfg: cfg,
		db:  database,
	}
}

// Login authenticates a user and returns a JWT token
func (s *AuthService) Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error) {
	// This is a placeholder implementation
	// In a real application, you would query the database for the user
	// For now, we'll return a mock response
	
	if req.Email == "" || req.Password == "" {
		return nil, fmt.Errorf("email and password are required")
	}

	// Mock user for demonstration
	user := models.User{
		ID:       "user-123",
		Email:    req.Email,
		Username: "demo_user",
	}

	token, expiresAt, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User:      user,
	}, nil
}

// Register creates a new user account
func (s *AuthService) Register(ctx context.Context, req models.RegisterRequest) (*models.AuthResponse, error) {
	// This is a placeholder implementation
	// In a real application, you would:
	// 1. Validate the request
	// 2. Check if user already exists
	// 3. Hash the password
	// 4. Insert into database
	
	if req.Email == "" || req.Password == "" || req.Username == "" {
		return nil, fmt.Errorf("email, password, and username are required")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Mock user creation
	user := models.User{
		ID:           "user-" + fmt.Sprintf("%d", time.Now().Unix()),
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	token, expiresAt, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User:      user,
	}, nil
}

func (s *AuthService) generateToken(user models.User) (string, time.Time, error) {
	expiresAt := time.Now().Add(time.Duration(s.cfg.JWTExpiryHours) * time.Hour)

	claims := &middleware.Claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, expiresAt, nil
}
