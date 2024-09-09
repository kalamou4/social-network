package handlers

import (
	"encoding/json"
	"net/http"
	"social-network/internal/models"
	"social-network/internal/services"
	"social-network/pkg/auth"

	"time"
)

var s services.UserService

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := s.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginData models.LoginData
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := s.AuthenticateUser(loginData.Email, loginData.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Generate a session token
	sessionToken, err := auth.GenerateToken(user.ID)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	// Set the session token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",
		Expires:  time.Now().Add(72 * time.Hour), // Example expiration time
		HttpOnly: true,                           // Prevent access to cookie from JavaScript
		Secure:   true,                           // Ensure cookie is sent only over HTTPS (useful for production)
	})

	// Send a response with user information and session token
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Login successful",
		"user": map[string]interface{}{
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"nickname":   user.Nickname,
			"gender":     user.Gender,                                 // Add gender if present in the user model
			"age":        time.Now().Year() - user.DateOfBirth.Year(), // Calculate age
		},
	}
	json.NewEncoder(w).Encode(response)
}
