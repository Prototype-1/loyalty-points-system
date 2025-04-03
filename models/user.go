package models

import "time"

type User struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Username     string    `gorm:"unique;not null" json:"username"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string   `gorm:"column:password_hash;not null" json:"-"` 
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
