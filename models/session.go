package models

import "time"

type Session struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	UserID       int       `gorm:"index;not null" json:"user_id"`
	Token        string    `gorm:"not null;unique" json:"token"`
	RefreshToken string    `gorm:"not null;unique" json:"refresh_token"`
	ExpiresAt    time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
