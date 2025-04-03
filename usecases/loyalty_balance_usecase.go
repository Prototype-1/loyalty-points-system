package usecase

import (
	"github.com/Prototype-1/loyalty-points-system/models"
	"github.com/Prototype-1/loyalty-points-system/repositories"
)

type LoyaltyPointsUsecase interface {
	GetUserPointsBalance(userID int) (int, error)
	GetUserPointsHistory(userID int) ([]models.LoyaltyPoints, error)
	RedeemUserPoints(userID int, points int) error 
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

func (u *loyaltyPointsUsecaseImpl) GetUserPointsHistory(userID int) ([]models.LoyaltyPoints, error) {
	return u.pointsRepo.GetPointsHistory(userID)
}

func (u *loyaltyPointsUsecaseImpl) RedeemUserPoints(userID int, points int) error {
	return u.pointsRepo.RedeemPoints(userID, points)
}
