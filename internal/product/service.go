package product

import (
	"errors"
	"time"
)


type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateProduct(name string, price, stock int) (*Product, error) {
	if name == "" || price == 0 || stock == 0 {
		return nil, errors.New("invalid input")
	}

	p := &Product{
		Name: name,
		Price: price,
		Stock: stock,
		CreatedAt: time.Now(),
	}

	return p, s.repo.Create(p)
}
