package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"inventory-service/internal/repository/model"
	"log"
	"time"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
	FindProductBySku(ctx context.Context, sku string) (*model.Product, error)
	FindAllProducts(ctx context.Context) ([]model.Product, error)
	AdjustStock(ctx context.Context, sku string, quantity int64) (*model.Product, error)
}

type productRepositoryPostgres struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepositoryPostgres{
		db: db,
	}
}

func (repository *productRepositoryPostgres) CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `INSERT INTO product_t (sku, name, quantity, reserved, price, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) 
              RETURNING id`

	err := repository.db.QueryRowContext(queryCtx, query,
		product.Sku,
		product.Name,
		product.Quantity,
		product.Reserved,
		product.Price,
		product.CreatedAt,
		product.UpdatedAt,
	).Scan(&product.Id)

	if err != nil {
		log.Printf("Failed to create product: %v", err)
		return nil, err
	}

	return product, nil
}

func (repository *productRepositoryPostgres) FindProductBySku(ctx context.Context, sku string) (*model.Product, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT id, sku, name, quantity, reserved, price, created_at, updated_at FROM product_t WHERE product_t.sku = $1`
	product, err := scanProduct(repository.db.QueryRowContext(queryCtx, query, sku))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("product with sku %s not found: %w", sku, err)
		}
		return nil, fmt.Errorf("failed to find product: %w", err)
	}

	return product, nil
}

func (repository *productRepositoryPostgres) FindAllProducts(ctx context.Context) ([]model.Product, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "SELECT id, sku, name, quantity, reserved, price, created_at, updated_at FROM product_t"
	rows, err := repository.db.QueryContext(queryCtx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("Failed to close rows: %v", err)
		}
	}(rows)

	products, err := scanProducts(rows)
	if err = rows.Err(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []model.Product{}, nil
		}
		return nil, fmt.Errorf("failed to scan products: %w", err)
	}

	return products, nil
}

func (repository *productRepositoryPostgres) AdjustStock(ctx context.Context, sku string, quantity int64) (*model.Product, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if quantity < 0 {
		return nil, fmt.Errorf("quantity adjustment cannot be negative: %d", quantity)
	}

	query := `UPDATE product_t
			  SET quantity = quantity + $1, updated_at = NOW()
			  WHERE sku = $2
			  AND quantity + $1 >= 0
			  RETURNING id, sku, name, quantity, reserved, price, created_at, updated_at`

	product, err := scanProduct(repository.db.QueryRowContext(queryCtx, query, quantity, sku))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("product %s not found or insufficient stock: %w", sku, err)
		}
		return nil, fmt.Errorf("failed to adjust stock: %w", err)
	}

	return product, nil
}

func scanProduct(row *sql.Row) (*model.Product, error) {
	product := &model.Product{}
	err := row.Scan(
		&product.Id,
		&product.Sku,
		&product.Name,
		&product.Quantity,
		&product.Reserved,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	return product, err
}

func scanProducts(rows *sql.Rows) ([]model.Product, error) {
	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.Id,
			&product.Sku,
			&product.Name,
			&product.Quantity,
			&product.Reserved,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, rows.Err()
}

func cancelContext(ctx context.Context) (context.Context, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return queryCtx, nil
}
