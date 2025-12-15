package cart

type CartItem struct {
	ProductId int `json:"product_id"`
	Quantity int `json:"quantity"`
}

type Cart struct {
	UserID int `json:"user_id"`
	Items []CartItem `json:"items"`
}
