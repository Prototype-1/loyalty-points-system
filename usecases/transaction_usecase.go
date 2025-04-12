package usecase

import (
	"errors"
	//"time"
	"github.com/Prototype-1/loyalty-points-system/models"
	"github.com/Prototype-1/loyalty-points-system/repositories"
)

type TransactionUsecase interface {
	AddTransaction(tx *models.Transaction) error
}

type transactionUsecaseImpl struct {
	transactionRepo repository.TransactionRepository
	loyaltyRepo     repository.LoyaltyRepository
}

func NewTransactionUsecase(tr repository.TransactionRepository, lr repository.LoyaltyRepository) TransactionUsecase {
	return &transactionUsecaseImpl{
		transactionRepo: tr,
		loyaltyRepo:     lr,
	}
}

func (u *transactionUsecaseImpl) AddTransaction(tx *models.Transaction) error {
	categoryMultipliers := map[string]float64{
		"electronics": 0.1,
		"groceries":   0.2,
		"fashion":   0.2,
	}

	multiplier, exists := categoryMultipliers[tx.Category]
	if !exists {
		return errors.New("invalid category")
	}

	pointsEarned := int(float64(tx.Amount) * multiplier)

	err := u.transactionRepo.CreateTransaction(tx)
	if err != nil {
		return err
	}

	loyaltyPoints := &models.LoyaltyPoints{
		UserID: tx.UserID,
		Points: pointsEarned,
		Status: "earned",
		Reason: "Earned via transaction in " + tx.Category,
	}

	return u.loyaltyRepo.AddLoyaltyPoints(loyaltyPoints)
}
