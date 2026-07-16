package service

import (
	"context"
	"inventory-service/internal/repository"
	"inventory-service/internal/repository/model"
	"inventory-service/internal/service/dto"
	"time"
)

type ProductService interface {
	FindAllProducts(ctx context.Context) ([]dto.Product, error)
	CreateProduct(ctx context.Context, product *dto.Product) (*dto.Product, error)
	FindProductBySku(ctx context.Context, sku string) (*dto.Product, error)
	AdjustStock(ctx context.Context, sku string, stock *dto.Stock) (*dto.Product, error)
}

type productCrudService struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) ProductService {
	return &productCrudService{repository: repository}
}

func (service *productCrudService) FindAllProducts(ctx context.Context) ([]dto.Product, error) {
	products, err := service.repository.FindAllProducts(ctx)
	if err != nil {
		return nil, err
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

	return productsResponse, nil
}

func (service *productCrudService) CreateProduct(ctx context.Context, product *dto.Product) (*dto.Product, error) {
	productEntity := model.Product{
		Sku:       product.Sku,
		Name:      product.Name,
		Quantity:  product.Quantity,
		Reserved:  product.Reserved,
		Price:     product.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}

	p, err := service.repository.CreateProduct(ctx, &productEntity)

	if err != nil {
		return nil, err
	}

	return &dto.Product{
		Sku:       p.Sku,
		Name:      p.Name,
		Quantity:  p.Quantity,
		Reserved:  p.Reserved,
		Price:     p.Price,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, nil
}

func (service *productCrudService) FindProductBySku(ctx context.Context, sku string) (*dto.Product, error) {
	product, err := service.repository.FindProductBySku(ctx, sku)

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

func (service *productCrudService) AdjustStock(ctx context.Context, sku string, stock *dto.Stock) (*dto.Product, error) {
	switch stock.Action {
	case "SUBTRACT":
		stock.Quantity = -stock.Quantity
	}

	product, err := service.repository.AdjustStock(ctx, sku, stock.Quantity)
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
