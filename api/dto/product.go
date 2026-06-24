package dto

import (
	"math/big"
	"time"
)

type ProductRequest struct {
	Sku       string    `json:"sku"`
	Name      string    `json:"name"`
	Quantity  int64     `json:"quantity"`
	Reserved  int64     `json:"reserved"`
	Price     big.Rat   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ProductResponse struct {
	Sku       string    `json:"sku"`
	Name      string    `json:"name"`
	Quantity  int64     `json:"quantity"`
	Reserved  int64     `json:"reserved"`
	Price     big.Rat   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
