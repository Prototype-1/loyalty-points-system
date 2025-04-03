package repository

import (
	"github.com/Prototype-1/loyalty-points-system/models"
	"gorm.io/gorm"
)

type LoyaltyRepository interface {
	AddLoyaltyPoints(points *models.LoyaltyPoints) error
	GetUserTotalPoints(userID int) (int, error)
}

type loyaltyRepositoryImpl struct {
	db *gorm.DB
}

func NewLoyaltyRepository(db *gorm.DB) LoyaltyRepository {
	return &loyaltyRepositoryImpl{db: db}
}

func (r *loyaltyRepositoryImpl) AddLoyaltyPoints(points *models.LoyaltyPoints) error {
	return r.db.Create(points).Error
}

func (r *loyaltyRepositoryImpl) GetUserTotalPoints(userID int) (int, error) {
	var total int
	err := r.db.Model(&models.LoyaltyPoints{}).Where("user_id = ?", userID).Select("SUM(points)").Scan(&total).Error
	return total, err
}
