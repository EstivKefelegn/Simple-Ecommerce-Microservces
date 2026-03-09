package handlers

import (
	"encoding/json"
	"github/ecommerceMSCGateway/clients"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)


// Login godoc
// @Summary Login user
// @Description Authenticate user and return token
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body clients.LoginRequest true "Login credentials"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {string} string
// @Router /auth/login [post]
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


// Register godoc
// @Summary Register new user
// @Description Create new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param register body clients.RegisterRequest true "Register data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {string} string
// @Router /auth/register [post]
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


// CreateOrder godoc
// @Summary Create order
// @Description Create new order for authenticated user
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body map[string]interface{} true "Order payload"
// @Security BearerAuth
// @Success 201 {object} map[string]interface{}
// @Failure 401 {string} string
// @Router /orders [post]
func CreateOrder(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)

	var order map[string]interface{}
	if err := json.Unmarshal(body, &order); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("user_id").(string)
	if !ok || userID == "" {
		http.Error(w, "unauthenticated user", http.StatusUnauthorized)
		return
	}

	order["user_id"] = userID

	newBody, _ := json.Marshal(order)

	resp, err := clients.CreateOrder(newBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	w.WriteHeader(resp.StatusCode)
	w.Write(data)
}


// GetProducts godoc
// @Summary Get all products
// @Description Retrieve all available products
// @Tags Products
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /products [get]
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


// CreateProduct godoc
// @Summary Create product
// @Description Add new product (authenticated)
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product body map[string]interface{} true "Product data"
// @Success 201 {object} map[string]interface{}
// @Router /products [post]
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


// GetProduct godoc
// @Summary Get product by ID
// @Description Retrieve product details
// @Tags Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Router /product/{id} [get]
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