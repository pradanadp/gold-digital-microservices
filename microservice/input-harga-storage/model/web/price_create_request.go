package web

type PriceCreateRequest struct {
	AdminId      string  `json:"admin_id"`
	TopupPrice   float64 `json:"topup_price"`
	BuybackPrice float64 `json:"buyback_price"`
}
