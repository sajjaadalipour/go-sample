package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

const UnknownError = "unknown_error"

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorHandler interface {
	Support(err error) bool
	Handle(err error) (int, ErrorResponse)
}

type ValidationErrorHandler struct {
}

func (v ValidationErrorHandler) Support(err error) bool {
	//TODO implement me
	return false
}

func (v ValidationErrorHandler) Handle(err error) (int, ErrorResponse) {
	return http.StatusBadRequest, ErrorResponse{
		Message: "X not supported!" + err.Error(),
		Code:    UnknownError,
	}
}

type UnknownErrorHandler struct {
}

func (u UnknownErrorHandler) Support(_ error) bool {
	return true
}

func (u UnknownErrorHandler) Handle(err error) (int, ErrorResponse) {
	return http.StatusInternalServerError, ErrorResponse{
		Message: "internal server err!" + err.Error(),
		Code:    UnknownError,
	}
}

func newHttpErrorHandler() func(err error, c echo.Context) {
	var handlers []ErrorHandler
	handlers = append(handlers, ValidationErrorHandler{})

	var unknownErrorHandler = UnknownErrorHandler{}

	return func(err error, c echo.Context) {
		var statusCode int
		var errorResponse ErrorResponse

		for i := range handlers {
			handler := handlers[i]
			if handler.Support(err) {
				statusCode, errorResponse = handler.Handle(err)
				return
			}
		}

		if statusCode == 0 {
			statusCode, errorResponse = unknownErrorHandler.Handle(err)
		}

		c.JSON(statusCode, errorResponse)
	}
}
