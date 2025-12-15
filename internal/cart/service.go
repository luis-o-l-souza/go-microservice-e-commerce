package cart

import "github.com/luis-o-l-souza/go-microservice-e-commerce/internal/cart/gateway"

type Service struct {
	g gateway.ProductGateway
}

func NewService(g gateway.ProductGateway) *Service {
	return &Service{g: g}
}
