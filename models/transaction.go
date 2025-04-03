package models

import "time"

type Transaction struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	TransactionID   string    `json:"transaction_id"`
	Amount          float64   `json:"amount"`
	Category        string    `json:"category"`
	TransactionDate time.Time `json:"transaction_date"`
	ProductCode     string    `json:"product_code"`
}
