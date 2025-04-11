package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/loyalty-points-system/handlers"
	"github.com/Prototype-1/loyalty-points-system/repositories"
	"github.com/Prototype-1/loyalty-points-system/usecases"
	"github.com/Prototype-1/loyalty-points-system/database"
	"github.com/Prototype-1/loyalty-points-system/middleware"
)

func SetupRoutes(r *gin.Engine) {
	db := database.GetDB()

	userRepo := repository.NewUserRepository(db)
	sessionRepo := repository.NewSessionRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, sessionRepo)
	userHandler := handlers.NewUserHandler(userUsecase, db)

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signup", userHandler.SignupHandler)
		authRoutes.POST("/login", userHandler.LoginHandler)
		authRoutes.POST("/logout", userHandler.LogoutHandler)
	}

	transactionRepo := repository.NewTransactionRepository(db)
	loyaltyRepo := repository.NewLoyaltyRepository(db)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo, loyaltyRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionUsecase, db)

	transactionRoutes := r.Group("/transactions")
	transactionRoutes.Use(middleware.AuthMiddleware()) 
	{
		transactionRoutes.POST("/add", transactionHandler.AddTransactionHandler)
	}

	pointsRepo := repository.NewLoyaltyPointsRepository(db)
	pointsUsecase := usecase.NewLoyaltyPointsUsecase(pointsRepo)
	pointsHandler := handlers.NewLoyaltyPointsHandler(pointsUsecase, db)

	pointsRoutes := r.Group("/points")
	pointsRoutes.Use(middleware.AuthMiddleware()) 
	{
		pointsRoutes.GET("/balance", pointsHandler.GetPointsBalanceHandler)
		pointsRoutes.GET("/history", pointsHandler.GetPointsHistoryHandler)
		pointsRoutes.POST("/redeem", pointsHandler.RedeemPointsHandler)
	}
}