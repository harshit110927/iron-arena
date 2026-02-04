import Foundation
import Combine

/// MainViewModel follows MVVM pattern
/// Manages UI-related data and business logic for ContentView
class MainViewModel: ObservableObject {
    @Published var welcomeMessage = "Mobile-first product built with SwiftUI"
    @Published var isLoading = false
    
    init() {
        // Initialize any data here
    }
    
    func onGetStarted() {
        isLoading = true
        // TODO: Navigate to next screen or perform action
        DispatchQueue.main.asyncAfter(deadline: .now() + 0.5) {
            self.isLoading = false
        }
    }
}
