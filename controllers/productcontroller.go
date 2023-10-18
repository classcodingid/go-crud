package controllers

import (
	"encoding/json"
	"go-crud/database"
	"go-crud/models"
	"net/http"

	"github.com/gorilla/mux"
)

// insert
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Create(&product)
	json.NewEncoder(w).Encode(product)
}

// get by id
func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product models.Product
	database.Instance.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// get all
func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	database.Instance.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// update
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product models.Product
	database.Instance.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// delete
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product models.Product
	database.Instance.Delete(&product, productId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}

// cek jika product sudah ada
func checkIfProductExists(productId string) bool {
	var product models.Product
	database.Instance.First(&product, productId)
	if product.ID == 0 {
		return false
	}
	return true
}
