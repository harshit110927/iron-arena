package com.ironarena.app.data.model

/**
 * User data model
 * Represents a user in the application
 */
data class User(
    val id: String,
    val email: String,
    val username: String
)

/**
 * Auth response model
 * Response from authentication endpoints
 */
data class AuthResponse(
    val token: String,
    val expiresAt: String,
    val user: User
)

/**
 * Login request model
 */
data class LoginRequest(
    val email: String,
    val password: String
)

/**
 * Register request model
 */
data class RegisterRequest(
    val email: String,
    val password: String,
    val username: String
)
