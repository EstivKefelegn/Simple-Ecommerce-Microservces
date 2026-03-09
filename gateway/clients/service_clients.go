package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

var authServiceURL = "http://localhost:8081"
var orderServiceURL = "http://localhost:8082"
var productServiceURL = "http://localhost:8083"

// LoginRequest represents login payload
type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"password123"`
}

// RegisterRequest represents register payload
type RegisterRequest struct {
	Username string `json:"username" example:"john"`
	Email    string `json:"email" example:"john@example.com"`
	Password string `json:"password" example:"password123"`
}

type ValidateRequest struct {
	Token string `json:"token"`
}

type ValidateResponse struct {
	UserID string `json:"user_id"`
	Error  bool   `json:"error"`
}
func Login(data LoginRequest) (*http.Response, error) {

	jsonData, _ := json.Marshal(data)

	return http.Post(
		authServiceURL+"/auth/login",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
}

func Register(data RegisterRequest) (*http.Response, error) {

	jsonData, _ := json.Marshal(data)

	return http.Post(
		authServiceURL+"/auth/register",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
}

func ValidateTokenAndGetUserID(token string) (string, bool) {
	req, _ := http.NewRequest("GET", authServiceURL+"/validate", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", false
	}

	userIDBytes, _ := io.ReadAll(resp.Body)
	userID := string(userIDBytes)
	if userID == "" {
		return "", false
	}

	return userID, true
}
func CreateOrder(body []byte) (*http.Response, error) {

	return http.Post(
		orderServiceURL+"/orders",
		"application/json",
		bytes.NewBuffer(body),
	)
}

func GetProducts() (*http.Response, error) {

	return http.Get(productServiceURL + "/products")
}

func CreateProduct(body []byte) (*http.Response, error) {

	return http.Post(
		productServiceURL+"/products",
		"application/json",
		bytes.NewBuffer(body),
	)
}

func GetProduct(id uuid.UUID) (*http.Response, error) {
	url := fmt.Sprintf("%s/products/%s", productServiceURL, id)
	return http.Get(url)
}
