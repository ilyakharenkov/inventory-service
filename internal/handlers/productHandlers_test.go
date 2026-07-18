package handlers

import (
	"inventory-service/internal/service/dto"
	"net/http"
	"testing"
)

func TestProductHttpHandler_FindAllProducts(t *testing.T) {
	testCases := struct {
		name           string
		mockResponse   []dto.Product
		mockError      error
		expectedStatus int
		expectedBody   interface{}
	}{
		name:           "success",
		mockResponse:   []dto.Product{},
		mockError:      nil,
		expectedStatus: http.StatusOK,
		expectedBody:   []dto.Product{},
	}

	t.Run(testCases.name, func(t *testing.T) {

	})
}
