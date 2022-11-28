package controller

import (
	"backend-go-loyalty/internal/dto"
	service "backend-go-loyalty/internal/service/product"
	"encoding/json"
	"net/http"
)

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product dto.ProductRequest
	json.NewDecoder(r.Body).Decode(&product)
	service.ProductService.InsertProduct(&product)
	json.NewEncoder(w).Encode(product)
}
