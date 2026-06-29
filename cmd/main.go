package main

import (
	"encoding/json"
	"fmt"
	"inventoiry-service/api"
	"inventoiry-service/internal/repository"
	"inventoiry-service/internal/service"
	"inventoiry-service/internal/service/dto"
	"net/http"
)

func main() {
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository)
	productHandler := api.NewProductHttpHandler(productService)

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var requestBody dto.Product
			if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			response := productHandler.CreateProduct(&requestBody)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		case http.MethodGet:
			sku := r.URL.Query().Get("sku")
			response := productHandler.FindProductBySku(sku)

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

	http.HandleFunc("/products/stock", func(w http.ResponseWriter, r *http.Request) {
		var requestBody dto.Stock
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		sku := r.URL.Query().Get("sku")
		response := productHandler.AdjustStock(sku, &requestBody)

		if response == nil {
			http.Error(w, "product by sku not found", http.StatusNotFound)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Println(err)
	}
}
