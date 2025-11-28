package repository

import (
	"context"

	commonModels "saitama/internal/common/models"
	"saitama/modules/users/models"

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

func (u *userRepository) GetUserList(ctx context.Context, option *commonModels.QueryOption) ([]models.User, error) {
	db := u.db.WithContext(ctx)
	var users []models.User
	var limit int = 10
	var offset int = 0
	var order string = "id DESC"

	if option != nil {
		// Apply pagination if provided
		limit = option.Limit
		if option.Order != "" {
			order = option.Order
		}
		offset = (option.Page - 1) * limit
	}

	query := db.Model(&models.User{}).Order(order)
	query = query.Limit(limit).Offset(offset)
	result := query.Find(&users)
	return users, result.Error
}
