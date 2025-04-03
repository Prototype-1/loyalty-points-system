package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/loyalty-points-system/handlers"
	"github.com/Prototype-1/loyalty-points-system/repositories"
	"github.com/Prototype-1/loyalty-points-system/usecases"
	"github.com/Prototype-1/loyalty-points-system/database"
)

func SetupRoutes(r *gin.Engine) {
	db := database.GetDB()

	userRepo := repository.NewUserRepository(db)
	sessionRepo := repository.NewSessionRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, sessionRepo)
	userHandler := handlers.NewUserHandler(userUsecase)

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signup", userHandler.SignupHandler)
		authRoutes.POST("/login", userHandler.LoginHandler)
		authRoutes.POST("/logout", userHandler.LogoutHandler)
	}
}
