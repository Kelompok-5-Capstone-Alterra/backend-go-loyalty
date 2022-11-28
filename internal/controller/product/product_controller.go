package controller

import (
	"backend-go-loyalty/internal/entity"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entity.Product
	json.NewDecoder(r.Body).Decode(&product)
	Instance.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	var products []entity.Product
	Instance.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entity.Product
	Instance.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entity.Product
	Instance.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)
	Instance.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entity.Product
	Instance.Delete(&product, productId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}

func checkIfProductExists(productId string) bool {
	var product entity.Product
	Instance.First(&product, productId)
	if product.ProductID == 0 {
		return false
	}
	return true
}
