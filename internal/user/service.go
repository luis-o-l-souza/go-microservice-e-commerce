package user

import (
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)


type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(email, password string) (*User, error) {
	if email == "" || password == "" {
		return nil, errors.New("invalid input")
	}
	_, err := s.repo.GetByEmail(email)

	if err != nil && err != sql.ErrNoRows {
		return nil, errors.New("error validating the user email")
	}

	if err == nil {
		return nil, errors.New("user already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	u := &User{
		Email: email,
		Password: string(hashed),
		CreatedAt: time.Now(),
	}

	return u, s.repo.Create(u)
}

func (s *Service) CheckUserExists(userId int) error {
	if userId <= 0 {
		return errors.New("Invalid ID")
	}

	_, err := s.repo.GetById(userId)

	if err == sql.ErrNoRows {
		return errors.New("User does not exist")
	}

	return nil
}
