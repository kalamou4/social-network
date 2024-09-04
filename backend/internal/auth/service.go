package auth

import (
	"database/sql"
	"errors"
	"social-network/backend/internal/user"
	"social-network/backend/pkg/utils"
)

type AuthService struct {
	repo *AuthRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		repo: NewAuthRepository(),
	}
}

func (s *AuthService) Register(newUser user.User) error {
	// Check if a user with the same email or username already exists
	existingUser, err := s.repo.GetUserByEmailOrUsername(newUser.Email, newUser.Nick_name)
	if err != nil && err != sql.ErrNoRows {
		return err // Return if there's any error other than "no rows found"
	}

	if existingUser != nil {
		return errors.New("user with the same email or username already exists")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return err
	}

	// Update the user object with the hashed password
	newUser.Password = hashedPassword

	// Create the new user
	return s.repo.CreateUser(newUser)
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}
