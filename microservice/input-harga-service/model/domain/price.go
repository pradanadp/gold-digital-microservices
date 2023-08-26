package domain

import (
	"time"
)

type Price struct {
	PriceId      string
	AdminId      string
	TopupPrice   float64
	BuybackPrice float64
	CreatedAt    time.Time
}
