package models

import "time"

type PointsHistory struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	UserID      int       `gorm:"index;not null" json:"user_id"`
	Points      int       `gorm:"not null" json:"points"`
	Transaction string    `gorm:"not null" json:"transaction"`
	Reason      string    `gorm:"not null" json:"reason"`
	Date        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date"`
}
