package cart

import "context"

type CartItem struct {
	ProductId int `json:"product_id"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
}

type Cart struct {
	UserID int `json:"user_id"`
	Items []CartItem `json:"items"`
}


type Repository interface {
	Get(ctx context.Context, userId int) (*Cart, error)
	Save(ctx context.Context, cart *Cart) error
}
