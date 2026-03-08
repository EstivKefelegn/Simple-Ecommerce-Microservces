package handlers

import (
	"encoding/json"
	"github/ecommerceMSCGateway/clients"
	"io"
	"net/http"
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
