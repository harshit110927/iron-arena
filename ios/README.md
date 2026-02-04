# iOS App

SwiftUI-based iOS application using MVVM architecture.

## Architecture

- **Views**: SwiftUI views for UI
- **ViewModels**: Manages UI state and business logic
- **Models**: Data models

## Structure

```
IronArena/
├── IronArenaApp.swift      # App entry point
├── Views/
│   └── ContentView.swift   # Main view
├── ViewModels/
│   └── MainViewModel.swift # Main view model
├── Models/
│   └── User.swift         # Data models
└── Resources/             # Assets and resources
```

## Build

Open `IronArena.xcodeproj` in Xcode and build.

## Requirements

- Xcode 15 or later
- iOS 17.0+
- Swift 5.9+

## Running

1. Open the project in Xcode
2. Select a simulator or device
3. Press Cmd+R to build and run
