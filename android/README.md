# Android App

Kotlin-based Android application using Jetpack Compose and MVVM architecture.

## Architecture

- **UI Layer**: Jetpack Compose screens
- **ViewModel**: Manages UI state and business logic
- **Model**: Data models and repository patterns

## Structure

```
app/src/main/java/com/ironarena/app/
├── MainActivity.kt          # Main entry point
├── ui/
│   └── theme/              # Material Design theme
├── viewmodel/              # ViewModels for MVVM
└── data/
    └── model/              # Data models
```

## Build

```bash
./gradlew assembleDebug
```

## Run

```bash
./gradlew installDebug
```

## Requirements

- Android Studio Hedgehog or later
- Android SDK 24+
- Kotlin 1.9+
