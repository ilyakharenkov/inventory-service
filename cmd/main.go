package main

import (
	"encoding/json"
	"fmt"
	"inventoiry-service/api"
	"inventoiry-service/api/dto"
	"net/http"
)

func main() {
	productHandler := api.NewProductHttpHandler()

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var requestBody dto.ProductRequest
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
		var requestBody dto.StockRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		sku := r.URL.Query().Get("sku")
		response := productHandler.AdjustStock(sku, &requestBody)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Println(err)
	}
}
