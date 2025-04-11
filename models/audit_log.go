package models

import (
	"time"
)

type AuditLog struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index"` 
	Action    string    `gorm:"size:50"` 
	Details   string    `gorm:"type:json"` 
	CreatedAt time.Time
}
