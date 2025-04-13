package repository

import (
	"github.com/Prototype-1/loyalty-points-system/models"
	"gorm.io/gorm"
	"errors"
	"time"
	"gorm.io/gorm/clause"
	"fmt"
)

type LoyaltyPointsRepository interface {
	GetPointsBalance(userID int) (int, error)
	GetPointsHistory(userID int, startDate, endDate, status string, page, limit int) ([]models.LoyaltyPoints, int64, error) 
	RedeemPoints(userID int, points int) error
}

type loyaltyPointsRepositoryImpl struct {
	db *gorm.DB
}

func NewLoyaltyPointsRepository(db *gorm.DB) LoyaltyPointsRepository {
	return &loyaltyPointsRepositoryImpl{db}
}

func (r *loyaltyPointsRepositoryImpl) GetPointsBalance(userID int) (int, error) {
	var earned int
	var redeemed int

	err := r.db.Model(&models.LoyaltyPoints{}).
		Where("user_id = ? AND status = ?", userID, "earned").
		Select("COALESCE(SUM(points), 0)").Scan(&earned).Error
	if err != nil {
		return 0, err
	}

	err = r.db.Model(&models.LoyaltyPoints{}).
		Where("user_id = ? AND status = ?", userID, "redeemed").
		Select("COALESCE(SUM(points), 0)").Scan(&redeemed).Error
	if err != nil {
		return 0, err
	}

	return earned + redeemed, nil
}

func (r *loyaltyPointsRepositoryImpl) GetPointsHistory(userID int, startDate, endDate, status string, page, limit int) ([]models.LoyaltyPoints, int64, error) {
	var history []models.LoyaltyPoints
	var totalRecords int64

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

	query.Model(&models.LoyaltyPoints{}).Count(&totalRecords)

	offset := (page - 1) * limit
	err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&history).Error
	return history, totalRecords, err
}

func (r *loyaltyPointsRepositoryImpl) RedeemPoints(userID int, points int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var earnedPoints []models.LoyaltyPoints
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("user_id = ? AND status = ?", userID, "earned").
			Order("created_at ASC").Find(&earnedPoints).
			Find(&earnedPoints).Error
		if err != nil {
			return err
		}

		totalPoints := 0
		for _, p := range earnedPoints {
			totalPoints += p.Points
		}
		if totalPoints < points {
			return errors.New("insufficient points")
		}

		remaining := points

		for _, row := range earnedPoints {
			if remaining == 0 {
			break
			}

		var toRedeem int
		if row.Points <= remaining {
			toRedeem = row.Points
		} else  {
			toRedeem = remaining
		}

		newRedemption := models.LoyaltyPoints{
			UserID: userID,
			Points: -toRedeem,
			Status: "redeemed",
			Reason: fmt.Sprintf("Redeemed from earned ID %d", row.ID),
		}
		if err := tx.Create(&newRedemption).Error; err != nil {
			return err
		}
		remaining -= toRedeem
	}
		return nil
	})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

