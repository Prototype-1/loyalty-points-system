package repository

import (
	"errors"
	"gorm.io/gorm"
	"github.com/Prototype-1/loyalty-points-system/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (r *userRepositoryImpl) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
