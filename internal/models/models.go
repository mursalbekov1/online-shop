package models

import (
	"database/sql"
	"online-shop/internal/repository/postgres"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Models struct {
	User interface {
		Create(user *User) error
		GetById(id int) (*User, error)
		GetAll() (*[]User, error)
		Update(user *User) error
		Delete(id int) error
	}

	Users postgres.UserRepository
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users: postgres.UserRepository{DB: db},
	}
}
