package main

import (
	"database/sql"
	"fmt"
	"inventory-service/configs"
	"inventory-service/internal/handlers"
	"inventory-service/internal/pkg/utils"
	"inventory-service/internal/repository"
	"inventory-service/internal/service"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	postgresConfig := configs.PostgresConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgresConfig.DBHost, postgresConfig.DBPort, postgresConfig.DBUser, postgresConfig.DBPassword, postgresConfig.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository)
	productHandler := handlers.NewProductHttpHandler(productService, utils.NewCustomValidator())

	http.HandleFunc("POST /products", productHandler.CreateProduct)
	http.HandleFunc("GET /products", productHandler.FindAllProducts)
	http.HandleFunc("PATCH /products/stock", productHandler.AdjustStock)
	http.HandleFunc("GET /products/{sku}", productHandler.FindProductBySku)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Println(err)
	}
}
