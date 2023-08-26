package domain

import "time"

type Topup struct {
	TopupId       string
	AccountNumber string
	Quantity      float64
	Price         float64
	CreatedAt     time.Time
}
