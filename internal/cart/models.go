package cart

type ProductDTO struct {
	ID int `json:"id"`
	Price int `json:"price"`
	Stock int `json:"stock"`
}
