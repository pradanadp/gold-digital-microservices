package web

type BalanceCheckResponse struct {
	AccountNumber string  `json:"norek"`
	Balance       float64 `json:"saldo"`
}
