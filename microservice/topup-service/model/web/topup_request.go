package web

type TopupRequest struct {
	AccountNumber string  `json:"norek" validate:"required"`
	Price         float64 `json:"harga" validate:"required"`
	Quantity      float64 `json:"gram" validate:"required"`
}
