#!/bin/bash

# Test script to run tests across all platforms

set -e

echo "🧪 Running tests for Iron Arena..."

# Backend tests
echo ""
echo "Testing Backend..."
cd backend
go test ./... -v
cd ..

# Android tests (if gradle wrapper exists)
if [ -f "android/gradlew" ]; then
    echo ""
    echo "Testing Android..."
    cd android
    ./gradlew test --no-daemon
    cd ..
fi

echo ""
echo "✅ All tests passed!"
