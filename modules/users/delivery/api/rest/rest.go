package rest

import (
	db "saitama/infrastructure/sqlite"
	"saitama/modules/users/repository"
	"saitama/modules/users/usecase"

	"gorm.io/gorm"
)

type module struct {
	db *gorm.DB
	uc usecase.UserUsecaseInterface
}

// uc := usecase.NewUserUsecase(db.GetDatabaseConnection(), repository.NewUserRepository(db.GetDatabaseConnection()))
func NewRestApplication() *module {
	database := db.GetDatabaseConnection()
	repo := repository.NewUserRepository(database)
	uc := usecase.NewUserUsecase(database, repo)
	return &module{
		db: database,
		uc: uc,
	}
}
