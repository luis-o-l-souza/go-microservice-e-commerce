package repository

import (
	"database/sql"
	"strconv"

	"github.com/luis-o-l-souza/go-microservice-e-commerce/internal/user"
)


type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{DB: db}
}

func (r *PostgresRepository) Create(u *user.User) error {
	query := `INSERT INTO users (email, password, created_at)
				VALUES ($1, $2, $3) RETURNING id`

	return r.DB.QueryRow(query, u.Email, u.Password, u.CreatedAt).Scan(&u.ID)
}

func (r *PostgresRepository) GetByEmail(email string) (*user.User, error) {
	u := &user.User{}

	query := `SELECT id, email, password, created_at FROM users where email = $1`

	err := r.DB.QueryRow(query, email).Scan(&u.ID, &u.Email, &u.Password, &u.CreatedAt)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *PostgresRepository) GetById(userId int) (*user.User, error) {
	u := &user.User{}

	query := `SELECT id FROM users WHERE id = $1`

	err := r.DB.QueryRow(query, strconv.Itoa(userId)).Scan(&u.ID)

	if err != nil {
		return nil, err
	}

	return u, nil
}
