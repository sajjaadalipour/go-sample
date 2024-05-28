package rest

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

const User = "/users"
const Health = "/health"

func RegisterRoutes(
	echo *echo.Echo,
	userController UserController,
	healthController HealthController,
) {
	echo.POST(User, userController.create)
	echo.GET(Health, healthController.getHealth)
}

var Provide = fx.Provide(
	newHealthController,
	newUserController,
)
