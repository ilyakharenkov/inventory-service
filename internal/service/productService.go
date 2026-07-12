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
	AdjustStock(sku string, stock *dto.Stock) (*dto.Product, error)
}

type productCrudService struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) ProductService {
	return &productCrudService{repository: repository}
}

func (service *productCrudService) FindAllProducts() []dto.Product {
	products, err := service.repository.FindAllProducts()
	if err != nil {
		return nil
	}

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

	p, err := service.repository.CreateProduct(&productEntity)

	if err != nil {
		return nil
	}

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
	product, err := service.repository.FindProductBySku(sku)

	if err != nil {
		return nil
	}

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

func (service *productCrudService) AdjustStock(sku string, stock *dto.Stock) (*dto.Product, error) {
	switch stock.Action {
	case "SUBTRACT":
		stock.Quantity = -stock.Quantity
	}

	product, err := service.repository.AdjustStock(sku, stock.Quantity)
	if err != nil {
		return nil, err
	}

	return &dto.Product{
		Sku:       product.Sku,
		Name:      product.Name,
		Quantity:  product.Quantity,
		Reserved:  product.Reserved,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}, nil
}
