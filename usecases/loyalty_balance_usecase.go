package usecase

import (
	"github.com/Prototype-1/loyalty-points-system/models"
	"github.com/Prototype-1/loyalty-points-system/repositories"
)

type LoyaltyPointsUsecase interface {
	GetUserPointsBalance(userID int) (int, error)
	GetUserPointsHistory(userID int, startDate, endDate, pointType string, page, limit int) ([]models.LoyaltyPoints, int64, error) 
	RedeemUserPoints(userID int, points int) (int, error)
}

type loyaltyPointsUsecaseImpl struct {
	pointsRepo repository.LoyaltyPointsRepository
}

func NewLoyaltyPointsUsecase(pointsRepo repository.LoyaltyPointsRepository) LoyaltyPointsUsecase {
	return &loyaltyPointsUsecaseImpl{pointsRepo}
}

func (u *loyaltyPointsUsecaseImpl) GetUserPointsBalance(userID int) (int, error) {
	return u.pointsRepo.GetPointsBalance(userID)
}

func (u *loyaltyPointsUsecaseImpl) GetUserPointsHistory(userID int, startDate, endDate, pointType string, page, limit int) ([]models.LoyaltyPoints, int64, error) {
	return u.pointsRepo.GetPointsHistory(userID, startDate, endDate, pointType, page, limit)
}

func (u *loyaltyPointsUsecaseImpl) RedeemUserPoints(userID int, points int) (int, error) {
	err := u.pointsRepo.RedeemPoints(userID, points)
	if err != nil {
		return 0, err
	}

	balance, err := u.pointsRepo.GetPointsBalance(userID)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

