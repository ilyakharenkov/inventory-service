package repository

import (
	"inventory-service/internal/repository/model"
	"sync"
)

type ProductRepository interface {
	CreateProduct(product *model.Product) *model.Product
	FindProductBySku(sku string) *model.Product
}

func NewProductRepository() ProductRepository {
	return &productRepositoryPostgres{
		products: make([]model.Product, 0),
	}
}

type productRepositoryPostgres struct {
	products []model.Product
	mu       sync.RWMutex
}

func (repository *productRepositoryPostgres) CreateProduct(product *model.Product) *model.Product {
	repository.mu.Lock()
	defer repository.mu.Unlock()

	repository.products = append(repository.products, *product)
	return product
}

func (repository *productRepositoryPostgres) FindProductBySku(sku string) *model.Product {
	repository.mu.RLock()
	defer repository.mu.RUnlock()

	for index := range repository.products {
		if repository.products[index].Sku == sku {
			return &repository.products[index]
		}
	}
	return nil
}
