package server

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	error2 "test/common/error"
)

func GlobalHttpErrorHandler() func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
		var appErr error2.AppError
		if errors.As(err, &appErr) {
			err := c.JSON(appErr.Status, map[string]interface{}{
				"message": appErr.Message,
				"code":    appErr.Code,
			})

			if err != nil {
				unknownError(c)
			}
			return
		}
		unknownError(c)
	}
}

func unknownError(c echo.Context) {
	_ = c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"message": "internal server error!",
		"code":    "unknown_error",
	})
}
