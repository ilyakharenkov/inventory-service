package handlers

import (
	"encoding/json"
	"inventory-service/internal/service"
	"inventory-service/internal/service/dto"
	"net/http"
)

type ProductHandler interface {
	FindAllProducts(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	FindProductBySku(w http.ResponseWriter, r *http.Request)
	AdjustStock(w http.ResponseWriter, r *http.Request)
}

type productHttpHandler struct {
	service service.ProductService
}

func NewProductHttpHandler(service service.ProductService) ProductHandler {
	return &productHttpHandler{service: service}
}

func (handler *productHttpHandler) FindAllProducts(w http.ResponseWriter, r *http.Request) {
	response := handler.service.FindAllProducts()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *productHttpHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var requestBody dto.Product
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := handler.service.CreateProduct(&requestBody)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *productHttpHandler) FindProductBySku(w http.ResponseWriter, r *http.Request) {
	sku := r.PathValue("sku")
	response := handler.service.FindProductBySku(sku)
	if response == nil {
		http.Error(w, "product by sku not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *productHttpHandler) AdjustStock(w http.ResponseWriter, r *http.Request) {
	var requestBody dto.Stock
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sku := r.URL.Query().Get("sku")
	response := handler.service.AdjustStock(sku, &requestBody)

	if response == nil {
		http.Error(w, "product by sku not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
