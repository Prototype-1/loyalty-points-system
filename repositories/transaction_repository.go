package repository

import (
	"github.com/Prototype-1/loyalty-points-system/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(tx *models.Transaction) error
}

type transactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepositoryImpl{db: db}
}

func (r *transactionRepositoryImpl) CreateTransaction(tx *models.Transaction) error {
	return r.db.Create(tx).Error
}
