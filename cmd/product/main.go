package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/luis-o-l-souza/go-microservice-e-commerce/internal/product"
	"github.com/luis-o-l-souza/go-microservice-e-commerce/internal/product/repository"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "postgres://user:pass@localhost:5433/product_db?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repo := repository.NewPostgresRepository(db)
	svc := product.NewService(repo)
	handler := product.NewHandler(svc)

	r := chi.NewRouter()
	handler.RegisterRoutes(r)

	log.Println("User service starting on :8081")
	http.ListenAndServe(":8081", r)
}
