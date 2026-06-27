package service

import (
	"inventoiry-service/internal/service/dto"
	"math/big"
	"time"
)

func NewProductService() ProductService {
	return &productCrudService{}
}

type ProductService interface {
	CreateProduct(product *dto.Product) *dto.Product
	FindProductBySku(sku string) *dto.Product
	AdjustStock(sku string, stock *dto.Stock) *dto.Product
}

type productCrudService struct {
}

func (service *productCrudService) CreateProduct(product *dto.Product) *dto.Product {
	return &dto.Product{
		Sku:       product.Sku,
		Name:      product.Name,
		Quantity:  product.Quantity,
		Reserved:  product.Reserved,
		Price:     product.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (service *productCrudService) FindProductBySku(sku string) *dto.Product {
	return &dto.Product{
		Sku:       sku,
		Name:      "Sku name",
		Quantity:  0,
		Reserved:  0,
		Price:     *big.NewRat(9, 99),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (service *productCrudService) AdjustStock(sku string, stock *dto.Stock) *dto.Product {
	return &dto.Product{
		Sku:       sku,
		Name:      "Sku name",
		Quantity:  stock.Quantity,
		Reserved:  0,
		Price:     *big.NewRat(9, 99),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
