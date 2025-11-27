package service

import (
	"context"
	"saitama/internal/models"
)

type UserServiceRepository interface {
	CreateUser(ctx context.Context, username string, email string) (uint64, error)
	GetUser(ctx context.Context, id uint64) (*models.User, error)
}
