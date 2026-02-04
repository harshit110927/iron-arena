package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/harshit110927/iron-arena/backend/internal/config"
	"github.com/harshit110927/iron-arena/backend/internal/db"
	"github.com/harshit110927/iron-arena/backend/internal/handlers"
	appMiddleware "github.com/harshit110927/iron-arena/backend/internal/middleware"
	"github.com/harshit110927/iron-arena/backend/internal/services"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connection
	var database *db.Database
	if cfg.DatabaseURL != "" {
		var err error
		database, err = db.NewDatabase(cfg.DatabaseURL)
		if err != nil {
			log.Printf("Warning: Failed to connect to database: %v", err)
			database = nil
		} else {
			defer database.Close()
		}
	} else {
		log.Println("Warning: Running without database connection (DATABASE_URL not set)")
	}

	// Initialize services
	authService := services.NewAuthService(cfg, database)
	
	// Initialize handlers
	healthHandler := handlers.NewHealthHandler(database)
	authHandler := handlers.NewAuthHandler(authService)

	// Setup router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Timeout(60 * time.Second))

	// Public routes
	r.Get("/health", healthHandler.HealthCheck)
	
	// Auth routes
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/login", authHandler.Login)
		r.Post("/register", authHandler.Register)
	})

	// Protected routes
	r.Route("/api", func(r chi.Router) {
		r.Use(appMiddleware.JWTAuth(cfg.JWTSecret))
		r.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "This is a protected route"}`))
		})
	})

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	go func() {
		log.Printf("Starting server on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
