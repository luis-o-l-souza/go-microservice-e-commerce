package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/luis-o-l-souza/go-microservice-e-commerce/internal/user"
	"github.com/luis-o-l-souza/go-microservice-e-commerce/internal/user/repository"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "postgres://user:pass@localhost:5432/user_db?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repo := repository.NewPostgresRepository(db)
	svc := user.NewService(repo)
	handler := user.NewHandler(svc)

	r := chi.NewRouter()
	handler.RegisterRoutes(r)

	log.Println("User service starting on :8080")
	http.ListenAndServe(":8080", r)
}
