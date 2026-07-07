package service

import (
	"inventory-service/internal/repository"
	"inventory-service/internal/repository/model"
	"inventory-service/internal/service/dto"
	"time"
)

type ProductService interface {
	FindAllProducts() []dto.Product
	CreateProduct(product *dto.Product) *dto.Product
	FindProductBySku(sku string) *dto.Product
	AdjustStock(sku string, stock *dto.Stock) *dto.Product
}

type productCrudService struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) ProductService {
	return &productCrudService{repository: repository}
}

func (service *productCrudService) FindAllProducts() []dto.Product {
	products := service.repository.FindAllProducts()

	productsResponse := make([]dto.Product, len(products))

	for i, v := range products {
		productsResponse[i] = dto.Product{
			Sku:       v.Sku,
			Name:      v.Name,
			Quantity:  v.Quantity,
			Reserved:  v.Reserved,
			Price:     v.Price,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return productsResponse
}

func (service *productCrudService) CreateProduct(product *dto.Product) *dto.Product {
	productEntity := model.Product{
		Sku:       product.Sku,
		Name:      product.Name,
		Quantity:  product.Quantity,
		Reserved:  product.Reserved,
		Price:     product.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}

	p := service.repository.CreateProduct(&productEntity)

	return &dto.Product{
		Sku:       p.Sku,
		Name:      p.Name,
		Quantity:  p.Quantity,
		Reserved:  p.Reserved,
		Price:     p.Price,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (service *productCrudService) FindProductBySku(sku string) *dto.Product {
	product := service.repository.FindProductBySku(sku)

	if product == nil {
		return nil
	}

	return &dto.Product{
		Sku:       product.Sku,
		Name:      product.Name,
		Quantity:  product.Quantity,
		Reserved:  product.Reserved,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func (service *productCrudService) AdjustStock(sku string, stock *dto.Stock) *dto.Product {
	product := service.repository.FindProductBySku(sku)
	if product == nil {
		return nil
	}

	switch stock.Action {
	case "ADD":
		product.Quantity += stock.Quantity
	case "SUBJECT":
		product.Quantity -= stock.Quantity
	}

	return &dto.Product{
		Sku:       product.Sku,
		Name:      product.Name,
		Quantity:  product.Quantity,
		Reserved:  product.Reserved,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}
