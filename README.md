# Iron Arena

Mobile-first product monorepo with backend, Android, and iOS applications.

## Project Structure

```
iron-arena/
├── backend/          # Go backend service (Chi + PostgreSQL + JWT)
├── android/          # Android app (Kotlin + Jetpack Compose + MVVM)
├── ios/              # iOS app (SwiftUI + MVVM)
├── docs/             # API and database documentation
└── scripts/          # Development and deployment scripts
```

## Quick Start

### Prerequisites

- **Backend**: Go 1.21+, PostgreSQL 14+
- **Android**: Android Studio Hedgehog+, JDK 17+
- **iOS**: Xcode 15+, macOS

### Setup

Run the setup script:
```bash
./scripts/setup.sh
```

Or manually setup each component:

#### Backend
```bash
cd backend
cp .env.example .env
# Update .env with your configuration
go mod download
go run cmd/server/main.go
```

#### Android
```bash
cd android
./gradlew build
# Open in Android Studio to run
```

#### iOS
```bash
cd ios
# Open IronArena.xcodeproj in Xcode to build and run
```

## Development

### Start Backend Server
```bash
cd backend
go run cmd/server/main.go
```

Server will be available at `http://localhost:8080`

### Run Tests
```bash
./scripts/test.sh
```

## Architecture

### Backend (Go)
- **Framework**: Chi router with net/http
- **Architecture**: Clean architecture pattern
- **Database**: PostgreSQL with pgx driver
- **Authentication**: JWT-based auth
- **API Documentation**: See [docs/api.md](docs/api.md)

### Android (Kotlin)
- **UI**: Jetpack Compose
- **Architecture**: MVVM pattern
- **Networking**: Retrofit + OkHttp
- **State Management**: ViewModel + StateFlow

### iOS (Swift)
- **UI**: SwiftUI
- **Architecture**: MVVM pattern
- **State Management**: @Published properties + Combine
- **Networking**: URLSession

## Documentation

- [API Documentation](docs/api.md)
- [Database Schema](docs/schema.md)
- [Backend README](backend/README.md)
- [Android README](android/README.md)
- [iOS README](ios/README.md)

## Contributing

1. Create a feature branch
2. Make your changes
3. Run tests: `./scripts/test.sh`
4. Submit a pull request

## License

TBD