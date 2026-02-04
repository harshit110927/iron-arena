package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/harshit110927/iron-arena/backend/internal/db"
)

type HealthHandler struct {
	db *db.Database
}

func NewHealthHandler(database *db.Database) *HealthHandler {
	return &HealthHandler{db: database}
}

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Database  string `json:"database"`
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	dbStatus := "healthy"
	if h.db != nil {
		if err := h.db.HealthCheck(ctx); err != nil {
			dbStatus = "unhealthy"
		}
	} else {
		dbStatus = "not_connected"
	}

	response := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Database:  dbStatus,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
