package rest

import (
	"test/usecase"
)

type UserDto struct {
	Id    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (u *UserDto) ToUser() usecase.User {
	return usecase.User{
		Name:  u.Name,
		Email: u.Email,
	}
}
