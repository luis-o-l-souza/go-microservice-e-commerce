package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/luis-o-l-souza/go-microservice-e-commerce/internal/cart"
	"github.com/luis-o-l-souza/go-microservice-e-commerce/internal/cart/gateway"
	"github.com/luis-o-l-souza/go-microservice-e-commerce/internal/cart/repository"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	defer rdb.Close()

	repo := repository.NewRedisRepository(rdb)

	client := new(http.Client)
	productGateway := gateway.NewHttpProductGateway("http://localhost:8081", client)
	userGateway := gateway.NewHttpUserGateway("http://localhost:8080", client)

	svc := cart.NewService(productGateway, repo, userGateway)
	handler := cart.NewHandler(svc)

	r := chi.NewRouter()
	handler.RegisterRoutes(r)

	log.Println("Cart service starting on :8082")
	http.ListenAndServe(":8082", r)
}
