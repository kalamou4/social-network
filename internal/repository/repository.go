package repository

import "social-network/internal/models"

type UserRepository interface {
    CreateUser(user *models.User) error
    GetUserByID(userID string) (*models.User, error)
    GetUserByEmail(email string) (*models.User, error)
    UpdateUser(user *models.User) error
    DeleteUser(userID int) error
    FollowUser(followerID, followingID int) error
    UnfollowUser(followerID, followingID int) error
    GetFollowers(userID int) ([]*models.User, error)
    GetFollowing(userID int) ([]*models.User, error)
    CheckIfFollowing(followerID, followingID int) (bool, error)
}

type SessionRepository interface {
	GetByID(id string) (*models.Session, error)
}
