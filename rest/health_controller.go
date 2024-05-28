package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthController struct {
}

func newHealthController() HealthController {
	return HealthController{}
}

func (HealthController) getHealth(c echo.Context) error {
	if err := c.JSON(
		http.StatusCreated,
		map[string]string{"message": "ok"},
	); err != nil {
		return err
	}

	return nil
}
