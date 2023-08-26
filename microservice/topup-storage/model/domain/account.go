package domain

import "time"

type Account struct {
	AccountNumber string
	Balance       float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
