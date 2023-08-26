package domain

import (
	"time"
)

type Transaction struct {
	TransactionId string
	AccountNumber string
	Quantity      float64
	TopupPrice    float64
	BuybackPrice  float64
	Type          string
	Balance       float64
	CreatedAt     time.Time
}
