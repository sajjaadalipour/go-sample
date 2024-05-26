package repository

import (
	"gorm.io/gorm"
	"net/http"
	"test/application"
	error2 "test/common/error"
)

type SQLUserRepository struct {
	db *gorm.DB
}

func NewSQLUserRepository(db *gorm.DB) application.UserRepository {
	return SQLUserRepository{db: db}
}

func (repository SQLUserRepository) Save(user application.User) error {
	var result = repository.db.Create(&user)

	if result.Error != nil {
		return error2.AppError{
			Message: "Failed to create a user",
			Cause:   result.Error,
			Code:    "user.duplicate",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}
