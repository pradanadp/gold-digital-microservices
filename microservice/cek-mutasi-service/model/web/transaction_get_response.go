package web

import "time"

type TransactionGetResponse struct {
	TransactionId string    `json:"transaction_id,omitempty"`
	AccountNumber string    `json:"norek,omitempty"`
	Quantity      float64   `json:"gram,omitempty"`
	Type          string    `json:"type,omitempty"`
	TopupPrice    float64   `json:"harga_topup,omitempty"`
	BuybackPrice  float64   `json:"harga_buyback,omitempty"`
	Balance       float64   `json:"saldo,omitempty"`
	CreatedAt     time.Time `json:"date,omitempty"`
}
