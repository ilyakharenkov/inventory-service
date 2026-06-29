package handlers

import (
	"encoding/json"
	"inventoiry-service/internal/service"
	"inventoiry-service/internal/service/dto"
	"net/http"
)

type ProductHandler interface {
	RegisterRoutes(mux *http.ServeMux)
	CreateProduct(request *dto.Product) *dto.Product
	FindProductBySku(sku string) *dto.Product
	AdjustStock(sku string, request *dto.Stock) *dto.Product
}

func NewProductHttpHandler(service service.ProductService) ProductHandler {
	return &productHttpHandler{service: service}
}

type productHttpHandler struct {
	service service.ProductService
}

func (handler *productHttpHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var requestBody dto.Product
			if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			response := handler.CreateProduct(&requestBody)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		case http.MethodGet:
			sku := r.URL.Query().Get("sku")
			response := handler.FindProductBySku(sku)

			if response == nil {
				http.Error(w, "product by sku not found", http.StatusNotFound)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/products/stock", func(w http.ResponseWriter, r *http.Request) {
		var requestBody dto.Stock
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		sku := r.URL.Query().Get("sku")
		response := handler.AdjustStock(sku, &requestBody)

		if response == nil {
			http.Error(w, "product by sku not found", http.StatusNotFound)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func (handler *productHttpHandler) CreateProduct(request *dto.Product) *dto.Product {
	return handler.service.CreateProduct(request)
}

func (handler *productHttpHandler) FindProductBySku(sku string) *dto.Product {
	product := handler.service.FindProductBySku(sku)
	if product == nil {
		return nil
	}
	return product
}

func (handler *productHttpHandler) AdjustStock(sku string, request *dto.Stock) *dto.Product {
	product := handler.service.AdjustStock(sku, request)
	if product == nil {
		return nil
	}
	return product
}
