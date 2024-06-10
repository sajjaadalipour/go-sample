package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
	repository "test/repository/user"
	"test/usecase"
)

type UserController struct {
	userService usecase.UserUseCase
}

func newUserController(userService usecase.UserUseCase) UserController {
	return UserController{userService: userService}
}

func (controller UserController) create(c echo.Context) error {
	var dto UserDto
	if err := (&echo.DefaultBinder{}).BindBody(c, &dto); err != nil {
		return err
	}

	err := c.Validate(dto)
	if err != nil {
		return err
	}

	if err := controller.userService.Create(dto.ToUser()); err != nil {
		return err
	}

	if errs := c.JSON(http.StatusCreated, dto); errs != nil {
		return errs
	}

	return nil
}

// DTOs

type UserDto struct {
	Id    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (u *UserDto) ToUser() repository.User {
	return repository.User{
		Name:  u.Name,
		Email: u.Email,
	}
}
