package web

type TopupRequest struct {
	AccountNumber string  `json:"norek"`
	Price         float64 `json:"harga"`
	Quantity      float64 `json:"gram"`
}
