package http

import (
	"context"
	"encoding/json"
	"github/productMCS/internal/domain"
	"github/productMCS/internal/ports/api"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	service api.ProductService
}

func NewHandler(s api.ProductService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Stock       int64   `json:"stock"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// basic validation
	if req.Name == "" || req.Price <= 0 {
		http.Error(w, "invalid product data", http.StatusBadRequest)
		return
	}

	product := &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	ctx := context.Background()

	err = h.service.CreateProduct(ctx, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(product)
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProduct(context.Background(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(product)
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {

	products, err := h.service.ListProducts(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(products)
}