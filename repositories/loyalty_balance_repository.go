package repository

import (
	"github.com/Prototype-1/loyalty-points-system/models"
	"gorm.io/gorm"
	"errors"
	"time"
)

type LoyaltyPointsRepository interface {
	GetPointsBalance(userID int) (int, error)
	GetPointsHistory(userID int, startDate, endDate, pointType string) ([]models.LoyaltyPoints, error) 
	RedeemPoints(userID int, points int) error
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

func (r *loyaltyPointsRepositoryImpl) GetPointsHistory(userID int, startDate, endDate, status string) ([]models.LoyaltyPoints, error) {
    var history []models.LoyaltyPoints
    query := r.db.Where("user_id = ?", userID)

    if startDate != "" && endDate != "" {
        layout := "2006-01-02 15:04:05"
        formattedStart := startDate + " 00:00:00"
        formattedEnd := endDate + " 23:59:59"

        loc, _ := time.LoadLocation("Asia/Kolkata") 
        startTime, _ := time.ParseInLocation(layout, formattedStart, loc)
        endTime, _ := time.ParseInLocation(layout, formattedEnd, loc)

        query = query.Where("created_at BETWEEN ? AND ?", startTime, endTime)
    }

    if status != "" {
        query = query.Where("status = ?", status)
    }

    err := query.Find(&history).Error
    return history, err
}

func (r *loyaltyPointsRepositoryImpl) RedeemPoints(userID int, points int) error {
	var totalPoints int

	err := r.db.Model(&models.LoyaltyPoints{}).
		Where("user_id = ?", userID).
		Select("SUM(points)").
		Scan(&totalPoints).Error

	if err != nil {
		return err
	}
	if totalPoints < points {
		return errors.New("insufficient points")
	}

	newRedemption := models.LoyaltyPoints{
		UserID: userID,
		//-ve value
		Points: -points,
		Status: "redeemed",
	}
	return r.db.Create(&newRedemption).Error
}

