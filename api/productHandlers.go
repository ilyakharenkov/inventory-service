package api

import (
	"fmt"
	"inventoiry-service/api/dto"
	"inventoiry-service/internal/service"
	"math/big"
	"time"
)

type ProductHandler interface {
	CreateProduct(request *dto.ProductRequest) *dto.ProductResponse
	FindProductBySku(sku string) *dto.ProductResponse
	AdjustStock(sku string, request *dto.StockRequest) *dto.ProductResponse
}

func NewProductHttpHandler(service service.ProductService) ProductHandler {
	return &productHttpHandler{service: service}
}

type productHttpHandler struct {
	service service.ProductService
}

func (handler *productHttpHandler) CreateProduct(request *dto.ProductRequest) *dto.ProductResponse {
	fmt.Printf("Body: %v\n", request)
	return &dto.ProductResponse{
		Sku:       request.Sku,
		Name:      request.Name,
		Quantity:  request.Quantity,
		Reserved:  request.Reserved,
		Price:     request.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (handler *productHttpHandler) FindProductBySku(sku string) *dto.ProductResponse {
	fmt.Printf("Sku: %v\n", sku)
	return &dto.ProductResponse{
		Sku:       sku,
		Name:      "Sku name",
		Quantity:  0,
		Reserved:  0,
		Price:     *big.NewRat(9, 99),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (handler *productHttpHandler) AdjustStock(sku string, request *dto.StockRequest) *dto.ProductResponse {
	fmt.Printf("Sku: %v\n", sku)
	fmt.Printf("Body: %v\n", request)
	return &dto.ProductResponse{
		Sku:       sku,
		Name:      "Sku name",
		Quantity:  request.Quantity,
		Reserved:  0,
		Price:     *big.NewRat(9, 99),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
