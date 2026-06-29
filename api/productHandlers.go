package api

import (
	"inventoiry-service/internal/service"
	"inventoiry-service/internal/service/dto"
)

type ProductHandler interface {
	CreateProduct(request *dto.Product) *dto.Product
	FindProductBySku(sku string) *dto.Product
	AdjustStock(sku string, request *dto.Stock) *dto.Product
}

func NewProductHttpHandler(service service.ProductService) ProductHandler {
	return &productHttpHandler{service: service}
}

type productHttpHandler struct {
	service service.ProductService
}

func (handler *productHttpHandler) CreateProduct(request *dto.Product) *dto.Product {
	return handler.service.CreateProduct(request)
}

func (handler *productHttpHandler) FindProductBySku(sku string) *dto.Product {
	product := handler.service.FindProductBySku(sku)
	if product == nil {
		return nil
	}
	return product
}

func (handler *productHttpHandler) AdjustStock(sku string, request *dto.Stock) *dto.Product {
	product := handler.service.AdjustStock(sku, request)
	if product == nil {
		return nil
	}
	return product
}
