package repository

import (
	"github.com/Prototype-1/loyalty-points-system/models"
	"gorm.io/gorm"
)

type LoyaltyPointsRepository interface {
	GetPointsBalance(userID int) (int, error)
	GetPointsHistory(userID int) ([]models.LoyaltyPoints, error)
}

type loyaltyPointsRepositoryImpl struct {
	db *gorm.DB
}

func NewLoyaltyPointsRepository(db *gorm.DB) LoyaltyPointsRepository {
	return &loyaltyPointsRepositoryImpl{db}
}

func (r *loyaltyPointsRepositoryImpl) GetPointsBalance(userID int) (int, error) {
	var totalPoints int
	err := r.db.Model(&models.LoyaltyPoints{}).
		Where("user_id = ?", userID).
		Select("SUM(points)").
		Scan(&totalPoints).Error
	return totalPoints, err
}

func (r *loyaltyPointsRepositoryImpl) GetPointsHistory(userID int) ([]models.LoyaltyPoints, error) {
	var history []models.LoyaltyPoints
	err := r.db.Where("user_id = ?", userID).Find(&history).Error
	return history, err
}
