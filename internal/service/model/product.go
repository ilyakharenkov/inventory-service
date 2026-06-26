package model

import (
	"math/big"
	"time"
)

type Product struct {
	Sku       string
	Name      string
	Quantity  int64
	Reserved  int64
	Price     big.Rat
	CreatedAt time.Time
	UpdatedAt time.Time
}
