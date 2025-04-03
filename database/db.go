package database

import (
	"fmt"
	"log"
	"github.com/Prototype-1/loyalty-points-system/config"
	"github.com/Prototype-1/loyalty-points-system/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db

	err = db.AutoMigrate(
		&models.User{},
		&models.Transaction{},
		&models.Session{},
		&models.PointsHistory{},
		&models.LoyaltyPoints{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migrated successfully!")
}

