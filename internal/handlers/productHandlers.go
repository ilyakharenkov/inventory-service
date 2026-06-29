package handlers

import (
	"encoding/json"
	"inventory-service/internal/service"
	"inventory-service/internal/service/dto"
	"net/http"
)

type ProductHandler interface {
	RegisterRoutes(mux *http.ServeMux)
}

type productHttpHandler struct {
	service service.ProductService
}

func NewProductHttpHandler(service service.ProductService) ProductHandler {
	return &productHttpHandler{service: service}
}

func (handler *productHttpHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateProduct(w, r)
		case http.MethodGet:
			handler.FindProductBySku(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/products/stock", handler.AdjustStock)
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
	sku := r.URL.Query().Get("sku")
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
