package middleware

import (
	"net/http"
	"strings"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/Prototype-1/loyalty-points-system/database"
	"github.com/Prototype-1/loyalty-points-system/models"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		secretKey := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token claims"})
			c.Abort()
			return
		}

		userID := int(claims["user_id"].(float64))

		db := database.GetDB()
		var session models.Session
		if err := db.Where("user_id = ? AND token = ?", userID, tokenString).First(&session).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session expired. Please log in again."})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
