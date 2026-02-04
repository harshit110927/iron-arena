package com.ironarena.app.viewmodel

import androidx.lifecycle.ViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow

/**
 * MainViewModel follows MVVM pattern
 * Manages UI-related data and business logic for MainActivity
 */
class MainViewModel : ViewModel() {
    
    private val _welcomeMessage = MutableStateFlow("Mobile-first product built with Kotlin & Jetpack Compose")
    val welcomeMessage: String
        get() = _welcomeMessage.value

    private val _isLoading = MutableStateFlow(false)
    val isLoading: StateFlow<Boolean> = _isLoading.asStateFlow()

    init {
        // Initialize any data here
    }

    fun onGetStarted() {
        // Handle get started action
        _isLoading.value = true
        // TODO: Navigate to next screen or perform action
    }
}
