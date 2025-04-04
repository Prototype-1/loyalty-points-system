package handlers

import (
	"net/http"
	"strconv"
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

// func (h *LoyaltyPointsHandler) GetPointsHistoryHandler(c *gin.Context) {
// 	userID, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 		return
// 	}

// 	startDate := c.Query("start_date") // YYYY-MM-DD
// 	endDate := c.Query("end_date") // YYY-MM-DD
// 	pointType := c.Query("type") // the status of point

// 	history, err := h.pointsUsecase.GetUserPointsHistory(userID.(int), startDate, endDate, pointType)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch points history"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"user_id": userID, "history": history})
// }

func (h *LoyaltyPointsHandler) GetPointsHistoryHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	startDate := c.Query("start_date") 
	endDate := c.Query("end_date")
	pointType := c.Query("type")

	// Extract pagination params
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	history, total, err := h.pointsUsecase.GetUserPointsHistory(userID.(int), startDate, endDate, pointType, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch points history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":       userID,
		"history":       history,
		"total_records": total,
		"page":          page,
		"limit":         limit,
	})
}

func (h *LoyaltyPointsHandler) RedeemPointsHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req struct {
		Points int `json:"points" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.pointsUsecase.RedeemUserPoints(userID.(int), req.Points)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Points redeemed successfully", "redeemed_points": req.Points})
}
