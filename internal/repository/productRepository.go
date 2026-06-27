package repository

import "inventoiry-service/internal/repository/model"

var products []model.Product

type ProductRepository interface {
	CreateProduct(product *model.Product) *model.Product
	FindProductBySku(sku string) *model.Product
	AdjustStock(sku string, quantity int64) *model.Product
}

func NewProductRepository() ProductRepository {
	return &productRepositoryPostgres{}
}

type productRepositoryPostgres struct {
}

func (repository *productRepositoryPostgres) CreateProduct(product *model.Product) *model.Product {
	return &model.Product{}
}

func (repository *productRepositoryPostgres) FindProductBySku(sku string) *model.Product {
	return &model.Product{}
}

func (repository *productRepositoryPostgres) AdjustStock(sku string, quantity int64) *model.Product {
	return &model.Product{}
}
