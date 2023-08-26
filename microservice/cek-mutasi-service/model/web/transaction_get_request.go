package web

type TransactionGetRequest struct {
	AccountNumber string `json:"norek"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
}
