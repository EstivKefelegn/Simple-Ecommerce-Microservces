package http

import (
	"encoding/json"
	"net/http"

	"github/orderService/internal/application"

	"github.com/google/uuid"
)

type Handler struct {
	service *application.OrderService
}

func NewHandler(s *application.OrderService) *Handler {
	return &Handler{s}
}

type CreateOrderRequest struct {
	ProductID string `json:"product_id"`
	UserID    string `json:"user_id"`
	Quantity  int    `json:"quantity"`
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return
	}

	productUUID, err := uuid.Parse(req.ProductID)
	if err != nil {
		http.Error(w, "invalid product_id", http.StatusBadRequest)
		return
	}

	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}

	orderID, err := h.service.CreateOrder(r.Context(), productUUID, userUUID, req.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := map[string]interface{}{
		"message":  "order created successfully",
		"order_id": orderID.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
