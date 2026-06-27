package service

import (
	"inventoiry-service/internal/service/model"
)

func NewProductService() ProductService {
	return &ProductCrudService{}
}

type ProductService interface {
	CreateProduct(product *model.Product) *model.Product
	FindProductBySku(sku string) *model.Product
	AdjustStock(sku string, stock *model.Stock) *model.Product
}

type ProductCrudService struct {
}

func (service *ProductCrudService) CreateProduct(product *model.Product) *model.Product {
	return &model.Product{}
}

func (service *ProductCrudService) FindProductBySku(sku string) *model.Product {
	return &model.Product{}
}

func (service *ProductCrudService) AdjustStock(sku string, stock *model.Stock) *model.Product {
	return &model.Product{}
}
