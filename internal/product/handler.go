package product

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

type CreateProductRequest struct {
	Name string `json:"name"`
	Price int `json:"price"`
	Stock int `json:"stock"`
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/create", h.createProduct)
	r.Get("/", h.getProducts)
	r.Get("/name", h.getProductsByName)
}

func (h *Handler) createProduct(w http.ResponseWriter, r *http.Request) {
	var req CreateProductRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateProduct(req.Name, req.Price, req.Stock)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.repo.GetProducts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(products)
}

func (h *Handler) getProductsByName(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := r.Form.Get("name")

	if name == "" {
		http.Error(w, "Required query param 'name' not found", http.StatusBadRequest)
		return
	}

	products, err := h.service.repo.GetByName(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
