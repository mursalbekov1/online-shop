package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"online-shop/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`

	err := r.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		fmt.Errorf("error creating user: %v", err)
	}

	return nil
}

func (r *UserRepository) GetById(id int) (*models.User, error) {
	user := &models.User{}

	query := `SELECT * FROM users WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		fmt.Errorf("error getting user: %v", err)
	}

	return user, nil
}

func (r *UserRepository) GetAll() ([]*models.User, error) {
	query := `SELECT * FROM users`
	rows, err := r.DB.Query(query)
	if err != nil {
		fmt.Errorf("error getting all users: %v", err)
	}
	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			fmt.Errorf("error getting users: %v", err)
		}
	}

	return users, nil
}

func (r *UserRepository) Update(user *models.User) error {
	query := `UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4`
	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password, user.Id)
	if err != nil {
		fmt.Errorf("error updating user: %v", err)
	}

	return nil
}

func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		fmt.Errorf("error deleting user: %v", err)
	}

	return nil
}
