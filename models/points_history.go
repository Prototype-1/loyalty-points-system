package models

import "time"

type PointsHistory struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Points      int       `json:"points"`
	Transaction string    `json:"transaction"` 
	Reason      string    `json:"reason"`
	Date        time.Time `json:"date"`
}
