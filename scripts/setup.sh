#!/bin/bash

# Development setup script for Iron Arena monorepo

set -e

echo "🚀 Setting up Iron Arena development environment..."

# Backend setup
echo ""
echo "📦 Setting up Backend (Go)..."
cd backend
if [ ! -f ".env" ]; then
    echo "Creating .env file from .env.example..."
    cp .env.example .env
    echo "⚠️  Please update .env with your configuration"
fi

echo "Installing Go dependencies..."
go mod download
echo "✅ Backend setup complete"

cd ..

# Android setup
echo ""
echo "📱 Setting up Android..."
cd android
if command -v ./gradlew &> /dev/null; then
    echo "Gradle wrapper found, building project..."
    ./gradlew build --no-daemon || echo "⚠️  Android build may require Android SDK setup"
else
    echo "⚠️  Gradle wrapper not found. Please setup Android Studio and Gradle"
fi
cd ..

echo ""
echo "✅ Setup complete!"
echo ""
echo "Next steps:"
echo "1. Update backend/.env with your database credentials"
echo "2. Run 'cd backend && go run cmd/server/main.go' to start the backend"
echo "3. Open android/ in Android Studio to run the Android app"
echo "4. Open ios/IronArena.xcodeproj in Xcode to run the iOS app"
