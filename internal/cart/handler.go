package cart

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

type AddToCartRequest struct {
	ProductId int `json:"product_id"`
	Amount int `json:"amount"`
	UserId int `json:"user_id"`
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/", h.addToCart)
}

func (h *Handler) addToCart(w http.ResponseWriter, r *http.Request) {

	var body AddToCartRequest

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if body.ProductId <= 0 || body.Amount <= 0 || body.UserId <= 0 {
		http.Error(w, fmt.Sprintf("Invalid payload: %+v", body), http.StatusBadRequest)
		return
	}

	price, stock, err := h.service.g.CheckProductExists(r.Context(), body.ProductId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if stock < body.Amount {
		http.Error(w, "Theres not enough stock for the product.", http.StatusBadRequest)
		return
	}

	err = h.service.u.CheckUserExists(r.Context(), body.UserId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.AddToCart(&body, price)

	if err != nil {
		http.Error(w, "Error when saving the cart", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product added to the cart!"))
}
