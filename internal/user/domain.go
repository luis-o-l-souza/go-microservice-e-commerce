package user


import "time"

type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}


type Repository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
}
