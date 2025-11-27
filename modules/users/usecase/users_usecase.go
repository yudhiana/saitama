package usecase

import (
	"context"

	"saitama/modules/users/models"
	"saitama/modules/users/repository"

	"gorm.io/gorm"
)

type userUsecase struct {
	db   *gorm.DB
	repo repository.UserRepository
}

func NewUserUsecase(db *gorm.DB, repo repository.UserRepository) userUsecaseInterface {
	return &userUsecase{
		db:   db,
		repo: repo,
	}
}

func (uc *userUsecase) CreateUser(ctx context.Context, name, email string) (id uint64, err error) {
	id, err = uc.repo.CreateUser(ctx, name, email)
	if err != nil {
		return
	}
	return id, nil
}

func (uc *userUsecase) GetUser(ctx context.Context, id uint64) (*models.User, error) {
	userData, err := uc.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return userData, nil
}
