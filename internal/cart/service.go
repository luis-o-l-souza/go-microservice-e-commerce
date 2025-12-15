package cart

import (
	"context"

	"github.com/luis-o-l-souza/go-microservice-e-commerce/internal/cart/gateway"
)

type Service struct {
	g gateway.ProductGateway
	u gateway.UserGateway
	r Repository
}

func NewService(g gateway.ProductGateway, r Repository, u gateway.UserGateway) *Service {
	return &Service{g: g, r: r, u: u}
}

func (s *Service) AddToCart(payload *AddToCartRequest, productPrice int) error {
	cart := &Cart{
		UserID: payload.UserId,
		Items: []CartItem{
			{
				ProductId: payload.ProductId,
				Quantity: payload.Amount,
				Price: productPrice,
			},
		},
	}
	return s.r.Save(context.Background(), cart)
}
