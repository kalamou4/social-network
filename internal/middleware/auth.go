package middleware

import (
	"context"
	"errors"
	"net/http"
	"social-network/internal/models"
	"social-network/internal/repository"
	"social-network/internal/services"
	"strconv"
	"strings"
	"time"
)

const userContextKey = "user"

// AuthenticationMiddleware checks for a valid session token and extracts user information
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// Extract and validate the token
		sessionID, err := extractSessionID(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Retrieve user information from session ID
		user, err := getUserFromSession(sessionID)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Set user information in context
		ctx := context.WithValue(r.Context(), userContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// extractSessionID parses the token and returns the session ID
func extractSessionID(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", errors.New("invalid token format")
	}

	// Assuming token structure is <sessionID>.<userID>.<expiration>
	sessionID := parts[0]
	expiration, err := parseExpiration(parts[2])
	if err != nil {
		return "", err
	}

	if time.Now().After(time.Unix(expiration, 0)) {
		return "", errors.New("token expired")
	}

	return sessionID, nil
}

// parseExpiration parses the expiration string to int64
func parseExpiration(expirationStr string) (int64, error) {
	return strconv.ParseInt(expirationStr, 10, 64)
}

// getUserFromSession retrieves the user based on the session ID
func getUserFromSession(sessionID string) (*models.User, error) {
	// Use your service to get the user based on the session ID
	// Example with a hypothetical session service
	var repo repository.SessionRepository
	sessionService := services.NewSessionService(repo)
	session, err := sessionService.GetSessionByID(sessionID)
	if err != nil || session == nil {
		return nil, err
	}

	var repo2 repository.UserRepository
	userService := services.NewUserService(repo2)
	return userService.GetUserByID(session.UserID)
}

// UserFromContext retrieves the user from the request context
func UserFromContext(ctx context.Context) (*models.User, bool) {
	user, ok := ctx.Value(userContextKey).(*models.User)
	return user, ok
}
