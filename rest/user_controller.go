package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"test/usecase"
)

type UserController struct {
	userService usecase.UserService
}

func newUserController(userService usecase.UserService) UserController {
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
