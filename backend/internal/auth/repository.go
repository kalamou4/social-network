package auth

import (
	"database/sql"
	"social-network/backend/internal/user"
	"social-network/backend/pkg/utils"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (r *AuthRepository) CreateUser(user user.User) error {
	var err error
	user.ID, err = utils.GenerateUUID()
	if err != nil {
		return err
	}
	_, err = r.db.Exec("INSERT INTO users (id, firstname, lastname, dateOfbirth, gender, password) VALUES (?, ?)", user.Nick_name, user.Password)
	return err
}

func (r *AuthRepository) GetUserByEmailOrUsername(email, usernameOrEmail string) (*user.User, error) {
	var user user.User
	rows, err := r.db.Query("SELECT nickname, email FROM Users WHERE email = ? OR username = ?", usernameOrEmail, email)
	if err != nil {
		return nil, err
	}

	for rows.Next(){

	}
	return &user, nil
}
