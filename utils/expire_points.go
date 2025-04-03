package utils

import (
	"log"
	"time"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"github.com/Prototype-1/loyalty-points-system/models"
)

//Is to make it expire points older than 1 year
func ExpireOldPoints(db *gorm.DB) {
	oneYearAgo := time.Now().AddDate(-1, 0, 0)

	var expiredPoints []models.LoyaltyPoints
	if err := db.Where("created_at <= ? AND status = ?", oneYearAgo, "earned").Find(&expiredPoints).Error; err != nil {
		log.Println("Error fetching old points:", err)
		return
	}

	for _, points := range expiredPoints {
		if err := db.Model(&models.LoyaltyPoints{}).
			Where("user_id = ?", points.UserID).
			Update("points_balance", gorm.Expr("points_balance - ?", points.Points)).Error; err != nil {
			log.Println("Error updating user balance:", err)
			continue
		}

		// Mark points as expired
		points.Status = "expired"
		points.CreatedAt = time.Now()
		if err := db.Save(&points).Error; err != nil {
			log.Println("Error updating points status:", err)
		}

		log.Printf("Expired %d points for user %d\n", points.Points, points.UserID)
	}
}

// Executes every day at midnight
func ScheduleExpirationJob(db *gorm.DB) {
	c := cron.New()
	_, err := c.AddFunc("@daily", func() {
		log.Println("Running points expiration job...")
		ExpireOldPoints(db)
		log.Println("Points expiration job completed.")
	})

	if err != nil {
		log.Fatal("Failed to schedule cron job:", err)
	}

	c.Start()
	log.Println("Points expiration cron job scheduled to run daily at midnight.")
}
