package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/luis-o-l-souza/go-microservice-e-commerce/internal/cart"
	"github.com/luis-o-l-souza/go-microservice-e-commerce/internal/cart/gateway"
)

func main() {
	client := new(http.Client)
	gateway := gateway.NewHttpProductGateway("http://localhost:8081", client)
	svc := cart.NewService(gateway)
	handler := cart.NewHandler(svc)

	r := chi.NewRouter()
	handler.RegisterRoutes(r)

	log.Println("Cart service starting on :8082")
	http.ListenAndServe(":8082", r)
}
