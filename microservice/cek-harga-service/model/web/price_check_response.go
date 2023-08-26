package web

type PriceCheckResponse struct {
	TopupPrice   float64 `json:"topup_price"`
	BuybackPrice float64 `json:"buyback_price"`
}
