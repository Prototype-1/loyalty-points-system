package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/loyalty-points-system/usecases"
)

type LoyaltyPointsHandler struct {
	pointsUsecase usecase.LoyaltyPointsUsecase
}

func NewLoyaltyPointsHandler(pointsUsecase usecase.LoyaltyPointsUsecase) *LoyaltyPointsHandler {
	return &LoyaltyPointsHandler{pointsUsecase}
}

func (h *LoyaltyPointsHandler) GetPointsBalanceHandler(c *gin.Context) {
	userID, exists := c.Get("userID") 
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	points, err := h.pointsUsecase.GetUserPointsBalance(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch points balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "points_balance": points})
}

func (h *LoyaltyPointsHandler) GetPointsHistoryHandler(c *gin.Context) {
	userID, exists := c.Get("userID") 
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	history, err := h.pointsUsecase.GetUserPointsHistory(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch points history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "history": history})
}
