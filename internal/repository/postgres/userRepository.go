package postgres

import (
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`

	err := r.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		fmt.Errorf("error creating user: %v", err)
	}

	return nil
}

func (r *UserRepository) GetById(id int) (*User, error) {
	user := &User{}

	query := `SELECT * FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		fmt.Errorf("error getting user: %v", err)
	}

	return user, nil
}

func (r *UserRepository) GetAll() ([]*User, error) {
	query := `SELECT * FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		fmt.Errorf("error getting all users: %v", err)
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			fmt.Errorf("error getting users: %v", err)
		}
	}

	return users, nil
}

func (r *UserRepository) Update(user *User) error {
	query := `UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4`
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.Id)
	if err != nil {
		fmt.Errorf("error updating user: %v", err)
	}

	return nil
}

func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Errorf("error deleting user: %v", err)
	}

	return nil
}
