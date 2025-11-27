package service

import (
	"context"
	"saitama/internal/models"
	"saitama/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserServiceRepository {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, username string, email string) (uint64, error) {
	return s.repo.CreateUser(ctx, username, email)
}

func (s *UserService) GetUser(ctx context.Context, id uint64) (*models.User, error) {
	return s.repo.GetUser(ctx, id)
}
