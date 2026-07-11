package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"inventory-service/internal/repository/model"
	"log"
	"sync"
)

type ProductRepository interface {
	CreateProduct(product *model.Product) *model.Product
	FindProductBySku(sku string) *model.Product
	FindAllProducts() []model.Product
	AdjustStock(sku string, quantity int64) (*model.Product, error)
}

type productRepositoryPostgres struct {
	db       *sql.DB
	mu       sync.RWMutex
	products []model.Product
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepositoryPostgres{
		db: db,
	}
}

func (repository *productRepositoryPostgres) CreateProduct(product *model.Product) *model.Product {
	query := `INSERT INTO product_t (sku, name, quantity, reserved, price, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) 
              RETURNING id`

	var id int64
	err := repository.db.QueryRow(query,
		product.Sku,
		product.Name,
		product.Quantity,
		product.Reserved,
		product.Price,
		product.CreatedAt,
		product.UpdatedAt,
	).Scan(&id)

	if err != nil {
		log.Printf("Failed to create product: %v", err)
		return nil
	}

	product.Id = id
	return product
}

func (repository *productRepositoryPostgres) FindProductBySku(sku string) *model.Product {
	query := `SELECT * FROM product_t WHERE product_t.sku = $1`
	product := &model.Product{}
	err := repository.db.QueryRow(query, sku).Scan(
		&product.Id,
		&product.Sku,
		&product.Name,
		&product.Quantity,
		&product.Reserved,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("Product not found: %s\n", sku)
		}
		return nil
	}

	return product
}

func (repository *productRepositoryPostgres) FindAllProducts() []model.Product {
	query := "SELECT * FROM product_t"
	rows, err := repository.db.Query(query)
	if err != nil {
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)
	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(
			&product.Id,
			&product.Sku,
			&product.Name,
			&product.Quantity,
			&product.Reserved,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt)
		if err != nil {
			return nil
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Products not found")
		}
		log.Printf("Rows error: %v", err)
		return nil
	}

	return products
}

func (repository *productRepositoryPostgres) AdjustStock(sku string, quantity int64) (*model.Product, error) {
	query := `UPDATE product_t
			  SET quantity = quantity + $1, updated_at = NOW()
			  WHERE sku = $2
			  AND quantity + $1 >= 0
			  RETURNING id, sku, name, quantity, reserved, price, created_at, updated_at`

	product := &model.Product{}
	err := repository.db.QueryRow(query, quantity, sku).Scan(
		&product.Id,
		&product.Sku,
		&product.Name,
		&product.Quantity,
		&product.Reserved,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Product not found: %s\n", sku)
		}
		return nil, fmt.Errorf("Product not found: %s\n", sku)
	}

	return product, nil
}
