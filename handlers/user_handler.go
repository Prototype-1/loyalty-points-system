package handlers

import (
	"net/http"
	"strings"
	"github.com/Prototype-1/loyalty-points-system/usecases"
	"github.com/Prototype-1/loyalty-points-system/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
	db *gorm.DB
}

func NewUserHandler(userUsecase usecase.UserUsecase, db *gorm.DB) *UserHandler {
	return &UserHandler{userUsecase, db}
}

func (h *UserHandler) SignupHandler(c *gin.Context) {
	var request struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userUsecase.Signup(request.Username, request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	utils.LogAudit(h.db, 0, "signup", map[string]interface{}{
		"username": request.Username,
		"email":    request.Email,
	})

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *UserHandler) LoginHandler(c *gin.Context) {
	var request struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.userUsecase.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userIDInt, _ := utils.ExtractUserIDFromToken(tokens.AccessToken)
	userID := uint(userIDInt)
	utils.LogAudit(h.db, userID, "login", map[string]interface{}{
		"email": request.Email,
	})

	c.JSON(http.StatusOK, gin.H{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}

func (h *UserHandler) LogoutHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	userID, err := utils.ExtractUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	uid := uint(userID)

	err = h.userUsecase.Logout(token, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.LogAudit(h.db, uid, "logout", map[string]interface{}{
		"message": "User logged out",
	})

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}


