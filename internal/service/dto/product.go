package dto

import (
	"math/big"
	"time"
)

type Product struct {
	Sku       string `json:"sku" validate:"required,min=3,max=50"`
	Name      string `json:"name" validate:"required,min=2,max=100"`
	Quantity  int64  `json:"quantity" validate:"min=0,max=500"`
	Reserved  int64  `json:"reserved" validate:"min=0"`
	Price     big.Rat
	CreatedAt time.Time
	UpdatedAt time.Time
}
