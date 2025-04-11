package utils

import (
	"encoding/json"
	"log"
	"time"
	"github.com/Prototype-1/loyalty-points-system/models"
	"gorm.io/gorm"
)

func LogAudit(db *gorm.DB, userID uint, action string, details interface{}) {
	jsonDetails, err := json.Marshal(details)
	if err != nil {
		log.Println("Failed to marshal audit details:", err)
		return
	}

	audit := models.AuditLog{
		UserID:    userID,
		Action:    action,
		Details:   string(jsonDetails),
		CreatedAt: time.Now(),
	}

	if err := db.Create(&audit).Error; err != nil {
		log.Println("Failed to write audit log:", err)
	}
}
