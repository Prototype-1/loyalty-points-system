package models

import "time"

type LoyaltyPoints struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Points    int       `json:"points"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
