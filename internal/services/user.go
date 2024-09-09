package services

import (
    "errors"
    "social-network/internal/models"
    "social-network/internal/repository"
    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{
        repo: repo,
    }
}

// CreateUser creates a new user and hashes the password
func (s *UserService) CreateUser(user *models.User) error {
    // Check if email already exists
    existingUser, _ := s.repo.GetUserByEmail(user.Email)
    if existingUser != nil {
        return errors.New("email already in use")
    }

    // Hash the password
    hashedPassword, err := s.HashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashedPassword

    return s.repo.CreateUser(user)
}

// GetUserByID retrieves user information by user ID
func (s *UserService) GetUserByID(userID string) (*models.User, error) {
    return s.repo.GetUserByID(userID)
}

// UpdateUser updates user profile information
func (s *UserService) UpdateUser(user *models.User) error {
    return s.repo.UpdateUser(user)
}

// DeleteUser deletes a user and all associated data (posts, followers, etc.)
func (s *UserService) DeleteUser(userID int) error {
    return s.repo.DeleteUser(userID)
}

// AuthenticateUser checks user credentials and returns the user if correct
func (s *UserService) AuthenticateUser(email, password string) (*models.User, error) {
    user, err := s.repo.GetUserByEmail(email)
    if err != nil || user == nil {
        return nil, errors.New("invalid email or password")
    }

    if !s.ComparePassword(user.Password, password) {
        return nil, errors.New("invalid email or password")
    }

    return user, nil
}

// FollowUser allows a user to follow another user
func (s *UserService) FollowUser(followerID, followingID int) error {
    // Check if user is already following
    isFollowing, _ := s.repo.CheckIfFollowing(followerID, followingID)
    if isFollowing {
        return errors.New("already following the user")
    }
    return s.repo.FollowUser(followerID, followingID)
}

// UnfollowUser allows a user to unfollow another user
func (s *UserService) UnfollowUser(followerID, followingID int) error {
    isFollowing, _ := s.repo.CheckIfFollowing(followerID, followingID)
    if !isFollowing {
        return errors.New("not following the user")
    }
    return s.repo.UnfollowUser(followerID, followingID)
}

// GetFollowers retrieves the list of users who follow the specified user
func (s *UserService) GetFollowers(userID int) ([]*models.User, error) {
    return s.repo.GetFollowers(userID)
}

// GetFollowing retrieves the list of users the specified user is following
func (s *UserService) GetFollowing(userID int) ([]*models.User, error) {
    return s.repo.GetFollowing(userID)
}

// HashPassword hashes a password using bcrypt
func (s *UserService) HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// ComparePassword compares a hashed password with its plaintext version
func (s *UserService) ComparePassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}

// CheckIfFollowing checks if a follower-following relationship exists
func (s *UserService) CheckIfFollowing(followerID, followingID int) (bool, error) {
    return s.repo.CheckIfFollowing(followerID, followingID)
}
