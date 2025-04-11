package handlers

import (
	"net/http"
	"time"
	"github.com/Prototype-1/loyalty-points-system/models"
	"github.com/Prototype-1/loyalty-points-system/usecases"
	"github.com/Prototype-1/loyalty-points-system/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionHandler struct {
	transactionUsecase usecase.TransactionUsecase
	db *gorm.DB
}

func NewTransactionHandler(tu usecase.TransactionUsecase, db *gorm.DB) *TransactionHandler {
	return &TransactionHandler{transactionUsecase: tu, db: db}
}

func (h *TransactionHandler) AddTransactionHandler(c *gin.Context) {
	var transaction models.Transaction

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction.UserID = userID.(int)
	transaction.TransactionDate = time.Now()

	if err := h.transactionUsecase.AddTransaction(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.LogAudit(
		h.db,
		uint(userID.(int)),
		"earn_points",
		map[string]interface{}{
			"transaction_id":   transaction.TransactionID,
			"amount":           transaction.Amount,
			"category":         transaction.Category,
			"product_code":     transaction.ProductCode,
			"transaction_date": transaction.TransactionDate.Format(time.RFC3339),
		},
	)

	c.JSON(http.StatusOK, gin.H{"message": "Transaction recorded successfully"})
}

