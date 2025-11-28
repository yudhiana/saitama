package repository

import (
	"context"
	commonModels "saitama/internal/common/models"
	"saitama/modules/users/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, username string, email string) (uint64, error)
	GetUser(ctx context.Context, id uint64) (*models.User, error)
	GetUserList(ctx context.Context, option *commonModels.QueryOption) ([]models.User, error)
}
