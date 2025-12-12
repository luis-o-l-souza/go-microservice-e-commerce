package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/luis-o-l-souza/go-microservice-e-commerce/internal/product"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{DB: db}
}

func (p *PostgresRepository) Create(product *product.Product) error {
	query := `INSERT INTO products (name, price, stock, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	return p.DB.QueryRow(query, &product.Name, &product.Price, &product.Stock, time.Now()).Scan(&product.ID)
}

func (p *PostgresRepository) GetProducts() ([]product.Product, error) {
	var products []product.Product

	query := `SELECT id, name, price, stock, created_at FROM products WHERE 1=1`
	rows, err := p.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product product.Product

		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CreatedAt); err != nil {
			return products, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
        return products, err
    }

    return products, nil
}

func (p *PostgresRepository) GetById(id int) (*product.Product, error) {
	product := &product.Product{}

	query := `SELECT id, name, price, stock, created_at FROM products WHERE 1=1 AND id = $1`
	err := p.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CreatedAt)
	return product, err
}

func (p *PostgresRepository) GetByName(name string) ([]product.Product, error) {
	var products []product.Product

	query := `SELECT id, name, price, stock, created_at FROM products WHERE 1=1 AND name LIKE '%' || $1 || '%'`
	rows, err := p.DB.Query(query, name)

	log.Printf("search by name. %v", err)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product product.Product
		log.Printf("search by 2. %v", products)

		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CreatedAt); err != nil {
			return products, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
        return products, err
    }
    log.Printf("search by name2. %v", products)

    return products, nil
}
