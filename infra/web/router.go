package web

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(
	router *echo.Echo,
	userController UserController,
	healthController HealthController,
) {
	router.POST(USER, userController.create)
	router.GET(HEALTH, healthController.getHealth)
}
