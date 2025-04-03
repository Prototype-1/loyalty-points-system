package models

import "time"

type LoyaltyPoints struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int       `gorm:"index;not null" json:"user_id"`
	Points    int       `gorm:"not null" json:"points"`
	Status    string    `gorm:"type:varchar(20);not null" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
