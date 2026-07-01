package handlers

import (
	"encoding/json"
	"inventory-service/internal/pkg/utils"
	"inventory-service/internal/service"
	"inventory-service/internal/service/dto"
	"net/http"
)

type ProductHandler interface {
	FindAllProducts(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	FindProductBySku(w http.ResponseWriter, r *http.Request)
	AdjustStock(w http.ResponseWriter, r *http.Request)
}

type productHttpHandler struct {
	service  service.ProductService
	validate utils.CustomValidator
}

func NewProductHttpHandler(service service.ProductService, cv *utils.CustomValidator) ProductHandler {
	return &productHttpHandler{
		service:  service,
		validate: *cv,
	}
}

func (handler *productHttpHandler) FindAllProducts(w http.ResponseWriter, r *http.Request) {
	response := handler.service.FindAllProducts()
	BaseResponse(response, w, http.StatusOK)
}

func (handler *productHttpHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var requestBody dto.Product
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.validate.Validate(requestBody); err != nil {
		http.Error(w, "Error validation", http.StatusBadRequest)
		return
	}

	response := handler.service.CreateProduct(&requestBody)
	BaseResponse(response, w, http.StatusCreated)
}

func (handler *productHttpHandler) FindProductBySku(w http.ResponseWriter, r *http.Request) {
	sku := r.PathValue("sku")
	response := handler.service.FindProductBySku(sku)
	if response == nil {
		http.Error(w, "product by sku not found", http.StatusNotFound)
		return
	}

	BaseResponse(response, w, http.StatusOK)
}

func (handler *productHttpHandler) AdjustStock(w http.ResponseWriter, r *http.Request) {
	var requestBody dto.Stock
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sku := r.URL.Query().Get("sku")
	response := handler.service.AdjustStock(sku, &requestBody)

	if response == nil {
		http.Error(w, "product by sku not found", http.StatusNotFound)
		return
	}

	BaseResponse(response, w, http.StatusOK)
}
