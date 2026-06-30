package main

import (
	"fmt"
	"inventory-service/internal/handlers"
	"inventory-service/internal/pkg/utils"
	"inventory-service/internal/repository"
	"inventory-service/internal/service"
	"net/http"
)

func main() {
	cv := utils.NewCustomValidator()

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository)
	productHandler := handlers.NewProductHttpHandler(productService, cv)

	http.HandleFunc("POST /products", productHandler.CreateProduct)
	http.HandleFunc("GET /products", productHandler.FindAllProducts)
	http.HandleFunc("PATCH /products/stock", productHandler.AdjustStock)
	http.HandleFunc("GET /products/{sku}", productHandler.FindProductBySku)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Println(err)
	}
}
