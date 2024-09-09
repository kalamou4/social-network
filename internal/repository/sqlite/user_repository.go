package sqlite

import (
	"database/sql"
	"social-network/internal/models"
	"social-network/internal/repository"
)

type UserRepositoryImpl struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
    return &UserRepositoryImpl{
        db: db,
    }
}

// CreateUser inserts a new user into the database
func (r *UserRepositoryImpl) CreateUser(user *models.User) error {
    query := "INSERT INTO users (username, email, password, first_name, last_name, date_of_birth, avatar_url, nickname, about_me, is_public, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
    _, err := r.db.Exec(query, user.Nickname, user.Email, user.Password, user.CreatedAt)
    return err
}

// GetUserByID retrieves a user by their ID
func (r *UserRepositoryImpl) GetUserByID(userID string) (*models.User, error) {
    var user models.User
    query := "SELECT id, nickname, email, created_at FROM users WHERE id = ?"
    err := r.db.QueryRow(query, userID).Scan(&user.ID, &user.Nickname, &user.Email, &user.CreatedAt)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// GetUserByEmail retrieves a user by their email
func (r *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    query := "SELECT id, nickname, email, password, created_at FROM users WHERE email = ?"
    err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// UpdateUser updates user profile information
func (r *UserRepositoryImpl) UpdateUser(user *models.User) error {
    query := "UPDATE users SET nickname = ?, email = ? WHERE id = ?"
    _, err := r.db.Exec(query, user.Nickname, user.Email, user.ID)
    return err
}

// DeleteUser deletes a user from the database
func (r *UserRepositoryImpl) DeleteUser(userID int) error {
    query := "DELETE FROM users WHERE id = ?"
    _, err := r.db.Exec(query, userID)
    return err
}

// FollowUser inserts a follow relationship between two users
func (r *UserRepositoryImpl) FollowUser(followerID, followingID int) error {
    query := "INSERT INTO followers (follower_id, following_id) VALUES (?, ?)"
    _, err := r.db.Exec(query, followerID, followingID)
    return err
}

// UnfollowUser removes a follow relationship between two users
func (r *UserRepositoryImpl) UnfollowUser(followerID, followingID int) error {
    query := "DELETE FROM followers WHERE follower_id = ? AND following_id = ?"
    _, err := r.db.Exec(query, followerID, followingID)
    return err
}

// GetFollowers retrieves the list of users who follow a specific user
func (r *UserRepositoryImpl) GetFollowers(userID int) ([]*models.User, error) {
    query := `
        SELECT u.id, u.nickname, u.email, u.created_at
        FROM users u
        JOIN followers f ON u.id = f.follower_id
        WHERE f.following_id = ?`
    
    rows, err := r.db.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var followers []*models.User
    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.ID, &user.Nickname, &user.Email, &user.CreatedAt)
        if err != nil {
            return nil, err
        }
        followers = append(followers, &user)
    }

    return followers, nil
}

// GetFollowing retrieves the list of users that a specific user is following
func (r *UserRepositoryImpl) GetFollowing(userID int) ([]*models.User, error) {
    query := `
        SELECT u.id, u.nickname, u.email, u.created_at
        FROM users u
        JOIN followers f ON u.id = f.following_id
        WHERE f.follower_id = ?`

    rows, err := r.db.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var following []*models.User
    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.ID, &user.Nickname, &user.Email, &user.CreatedAt)
        if err != nil {
            return nil, err
        }
        following = append(following, &user)
    }

    return following, nil
}

// CheckIfFollowing checks if a user is following another user
func (r *UserRepositoryImpl) CheckIfFollowing(followerID, followingID int) (bool, error) {
    var count int
    query := "SELECT COUNT(*) FROM followers WHERE follower_id = ? AND following_id = ?"
    err := r.db.QueryRow(query, followerID, followingID).Scan(&count)
    if err != nil {
        return false, err
    }
    return count > 0, nil
}
