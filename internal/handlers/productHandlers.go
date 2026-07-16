package handlers

import (
	"encoding/json"
	"inventory-service/internal/service"
	"inventory-service/internal/service/dto"
	"inventory-service/pkg/utils"
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
	response, err := handler.service.FindAllProducts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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

	response, err := handler.service.CreateProduct(r.Context(), &requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	BaseResponse(response, w, http.StatusCreated)
}

func (handler *productHttpHandler) FindProductBySku(w http.ResponseWriter, r *http.Request) {
	sku := r.PathValue("sku")
	response, err := handler.service.FindProductBySku(r.Context(), sku)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	response, err := handler.service.AdjustStock(r.Context(), sku, &requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	BaseResponse(response, w, http.StatusOK)
}
