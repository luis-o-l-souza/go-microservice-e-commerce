package product

import "time"

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Stock int `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
}

type Repository interface {
	Create(p *Product) error
	GetProducts() ([]Product, error)
	GetById(id int) (*Product, error)
	GetByName(name string) ([]Product, error)
}
