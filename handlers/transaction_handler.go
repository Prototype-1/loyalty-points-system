package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/loyalty-points-system/models"
	"github.com/Prototype-1/loyalty-points-system/usecases"
)

type TransactionHandler struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionHandler(tu usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{transactionUsecase: tu}
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

	c.JSON(http.StatusOK, gin.H{"message": "Transaction recorded successfully"})
}

