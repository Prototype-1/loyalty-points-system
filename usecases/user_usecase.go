package usecase

import (
	"errors"
	"github.com/Prototype-1/loyalty-points-system/models"
	"github.com/Prototype-1/loyalty-points-system/repositories"
	"github.com/Prototype-1/loyalty-points-system/utils"
	"log"
)

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserUsecase interface {
	Signup(username, email, password string) error
	Login(email, password string) (*AuthResponse, error)
	Logout(token string, userID int) error
}

type userUsecaseImpl struct {
	userRepo repository.UserRepository
	sessionRepo repository.SessionRepository
}

func NewUserUsecase(userRepo repository.UserRepository, sessionRepo repository.SessionRepository) UserUsecase {
	return &userUsecaseImpl{userRepo, sessionRepo}
}

func (u *userUsecaseImpl) Signup(username, email, password string) error {
	_, err := u.userRepo.GetUserByEmail(email)
	if err == nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		PasswordHash: hashedPassword,
	}

	return u.userRepo.CreateUser(user)
}

func (u *userUsecaseImpl) Login(email, password string) (*AuthResponse, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.ComparePassword(user.PasswordHash, password) {
		return nil, errors.New("invalid email or password")
	}

	accessToken, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	session := &models.Session{
		UserID:       user.ID,
		Token:        accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    utils.GetTokenExpiryTime(),
	}

	err = u.sessionRepo.CreateSession(session)
	if err != nil {
		return nil, errors.New("failed to save session")
	}

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *userUsecaseImpl) Logout(token string, userID int) error {
	if token == "" || userID == 0 {
		return errors.New("invalid token or user ID")
	}

	err := u.sessionRepo.DeleteSessionByTokenAndUserID(token, userID)
	if err != nil {
		log.Println("Logout Error:", err)
		return err
	}

	return nil
}
