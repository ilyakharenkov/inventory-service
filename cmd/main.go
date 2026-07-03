package main

import (
	"database/sql"
	"fmt"
	"inventory-service/configs"
	"inventory-service/internal/handlers"
	"inventory-service/internal/repository"
	"inventory-service/internal/service"
	"inventory-service/pkg/utils"
	"log"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	//m, err := migrate.New(
	//	"file://migrations",
	//	"postgres://user:pass@localhost:5432/dbname?sslmode=disable",
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer func(m *migrate.Migrate) {
	//	err, _ := m.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(m)
	//// Применить все миграции
	//if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
	//	log.Fatal(err)
	//}

	productRepository := repository.NewProductRepository(db)
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
