package model

import (
	"time"
)

type Product struct {
	Id        int64     `json:"id"`
	Sku       string    `json:"sku"`
	Name      string    `json:"name"`
	Quantity  int64     `json:"quantity"`
	Reserved  int64     `json:"reserved"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
