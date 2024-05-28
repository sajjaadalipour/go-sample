package repository

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
	"test/usecase"
)

var Provide = fx.Provide(newSQLUserRepository)

type SQLUserRepository struct {
	db *gorm.DB
}

func newSQLUserRepository(db *gorm.DB) usecase.UserRepository {
	return SQLUserRepository{db: db}
}

func (repository SQLUserRepository) Save(user usecase.User) error {
	var result = repository.db.Create(&user)

	return result.Error
}
