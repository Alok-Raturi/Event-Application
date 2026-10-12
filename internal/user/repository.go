package user

import (
	"database/sql"
	"errors"
)

type UserRepository interface {
	GetUsers() ([]User, error)
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUsers() ([]User, error) {
	rows, err := r.db.Query("SELECT id, name, email, password_hash, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userList []User

	for rows.Next() {

		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		userList = append(userList, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return userList, nil
}

func (r *userRepository) CreateUser(user *User) error {
	_, err := r.db.Exec("INSERT INTO users (id, name, email, password_hash, created_at) VALUES ($1, $2, $3, $4, $5)", user.ID, user.Name, user.Email, user.PasswordHash, user.CreatedAt)
	return err
}

func (r *userRepository) GetUserByEmail(email string) (*User, error) {
	row := r.db.QueryRow("Select id, email, name, password_hash, created_at FROM users WHERE email = $1", email)
	var u User
	err := row.Scan(&u.ID, &u.Email, &u.Name, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No user found with the given email
		}
		return nil, err
	}
	return &u, nil
}
