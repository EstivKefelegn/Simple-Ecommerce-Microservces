package http

import "net/http"

func Router(h *Handler) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/orders", h.CreateOrder)

	return mux
}