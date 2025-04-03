package models

import "time"

type Transaction struct {
	TransactionID   string    `gorm:"primaryKey" json:"transaction_id"`
	UserID          int       `gorm:"index" json:"user_id"`          
	Amount          float64   `gorm:"column:transaction_amount;not null" json:"amount"` 
	Category        string    `gorm:"not null" json:"category"`
	TransactionDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"transaction_date"`
	ProductCode     string    `json:"product_code"`
}

