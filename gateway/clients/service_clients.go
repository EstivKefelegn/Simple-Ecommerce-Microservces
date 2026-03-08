package clients

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var authServiceURL = "http://localhost:8081"
var orderServiceURL = "http://localhost:8082"
var productServiceURL = "http://localhost:8083"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(data LoginRequest) (*http.Response, error) {

	jsonData, _ := json.Marshal(data)

	return http.Post(
		authServiceURL+"/login",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
}

func ValidateToken(token string) bool {

	req, _ := http.NewRequest(
		"GET",
		authServiceURL+"/validate",
		nil,
	)

	req.Header.Set("Authorization", token)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
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