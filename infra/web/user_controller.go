package web

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"test/application"
)

type UserController struct {
	userService application.UserService
}

func NewUserController(userService application.UserService) UserController {
	return UserController{userService: userService}
}

func (controller UserController) create(c echo.Context) error {
	var user application.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := controller.userService.Create(user); err != nil {
		return err
	}

	if errs := c.JSON(http.StatusCreated, user); errs != nil {
		return errs
	}

	return nil
}
