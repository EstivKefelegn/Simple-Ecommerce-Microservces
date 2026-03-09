package handlers

import (
	"encoding/json"
	"github/ecommerceMSCGateway/clients"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var login clients.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := clients.Login(login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var register clients.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&register)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := clients.Register(register)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)

	resp, err := clients.CreateOrder(body)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	w.WriteHeader(resp.StatusCode)
	w.Write(data)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {

	resp, err := clients.GetProducts()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	w.WriteHeader(resp.StatusCode)
	w.Write(data)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)

	resp, err := clients.CreateProduct(body)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	w.WriteHeader(resp.StatusCode)
	w.Write(data)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return
	}
	resp, err := clients.GetProduct(uuidID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	w.WriteHeader(resp.StatusCode)
	w.Write(data)
}
