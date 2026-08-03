package handlers

import (
	"context"
	"inventory-service/internal/service/dto"
	"inventory-service/pkg/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductHttpHandler_FindAllProducts(t *testing.T) {
	testCases := struct {
		name           string
		mockResponse   []dto.Product
		mockError      error
		expectedStatus int
	}{
		name:           "success",
		mockResponse:   []dto.Product{},
		mockError:      nil,
		expectedStatus: http.StatusOK,
	}

	t.Run(testCases.name, func(t *testing.T) {

		service := &mockService{products: testCases.mockResponse, err: testCases.mockError}
		validator := utils.NewCustomValidator()

		handler := NewProductHttpHandler(service, validator)

		// Создаем HTTP запрос
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		req = req.WithContext(context.Background())

		// Создаем ResponseRecorder для записи ответа
		w := httptest.NewRecorder()

		handler.FindAllProducts(w, req)

		if w.Code != testCases.expectedStatus {
			t.Errorf("Expected status %d, got %d", testCases.expectedStatus, w.Code)
		}

		if w.Body.Len() == 0 {
			t.Error("response body is empty")
		}

		contentType := w.Header().Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("Content-Type: got %q, want %q", contentType, "application/json")
		}
	})
}

type mockService struct {
	products []dto.Product
	product  dto.Product
	err      error
}

func (s *mockService) FindAllProducts(ctx context.Context) ([]dto.Product, error) {
	return s.products, s.err
}

func (s *mockService) CreateProduct(ctx context.Context, product *dto.Product) (*dto.Product, error) {
	return &s.product, s.err
}

func (s *mockService) FindProductBySku(ctx context.Context, sku string) (*dto.Product, error) {
	return &s.product, s.err
}

func (s *mockService) AdjustStock(ctx context.Context, sku string, stock *dto.Stock) (*dto.Product, error) {
	return &s.product, s.err
}
