package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Models struct {
	Users interface {
		Create(user *User) error
		GetById(id int) (*User, error)
		GetAll() (*[]User, error)
		Update(user *User) error
		Delete(id int) error
	}

	//Users postgres.UserRepository
}

//func NewModels(db *sql.DB) Models {
//	return Models{
//		Users: postgres.UserRepository{DB: db},
//	}
//}
