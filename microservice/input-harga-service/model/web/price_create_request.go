package web

type PriceCreateRequest struct {
	AdminId      string  `json:"admin_id" validate:"required"`
	TopupPrice   float64 `json:"topup_price" validate:"required"`
	BuybackPrice float64 `json:"buyback_price" validate:"required"`
}
