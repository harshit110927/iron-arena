import Foundation

/// User data model
/// Represents a user in the application
struct User: Codable, Identifiable {
    let id: String
    let email: String
    let username: String
}

/// Auth response model
/// Response from authentication endpoints
struct AuthResponse: Codable {
    let token: String
    let expiresAt: String
    let user: User
}

/// Login request model
struct LoginRequest: Codable {
    let email: String
    let password: String
}

/// Register request model
struct RegisterRequest: Codable {
    let email: String
    let password: String
    let username: String
}
