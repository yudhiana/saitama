package repository

import (
	"context"
	"saitama/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(ctx context.Context, username string, email string) (uint64, error) {
	db := u.db.WithContext(ctx)
	user := models.User{
		Name:  username,
		Email: email,
	}
	result := db.Create(&user)
	return user.ID, result.Error
}

func (u *userRepository) GetUser(ctx context.Context, id uint64) (*models.User, error) {
	db := u.db.WithContext(ctx)
	var user models.User
	result := db.First(&user, id)
	return &user, result.Error
}
