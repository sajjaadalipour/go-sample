package repository

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (userRepository UserRepository) Save(user User) error {
	var result = userRepository.db.Create(&user)

	return result.Error
}
