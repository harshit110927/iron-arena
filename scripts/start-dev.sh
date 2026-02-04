#!/bin/bash

# Start all services for local development

set -e

echo "🚀 Starting Iron Arena services..."

# Start backend
echo "Starting backend server..."
cd backend
go run cmd/server/main.go &
BACKEND_PID=$!
echo "Backend started with PID: $BACKEND_PID"

cd ..

echo ""
echo "✅ Services started!"
echo "Backend: http://localhost:8080"
echo "Health check: http://localhost:8080/health"
echo ""
echo "Press Ctrl+C to stop all services"

# Wait for interrupt
trap "echo 'Stopping services...'; kill $BACKEND_PID 2>/dev/null; exit" INT TERM

wait
