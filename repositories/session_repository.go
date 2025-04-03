package repository

import (
	"gorm.io/gorm"
	"github.com/Prototype-1/loyalty-points-system/models"
	"fmt"
)

type SessionRepository interface {
	CreateSession(session *models.Session) error
	DeleteSessionByTokenAndUserID(token string, userID int) error 
}

type sessionRepositoryImpl struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepositoryImpl{db}
}

func (r *sessionRepositoryImpl) CreateSession(session *models.Session) error {
	return r.db.Create(session).Error
}

func (r *sessionRepositoryImpl) DeleteSessionByTokenAndUserID(token string, userID int) error {
	result := r.db.Where("token = ? AND user_id = ?", token, userID).Delete(&models.Session{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no session found with the given token and user ID")
	}
	return nil
}

