package utils

import (
	"github.com/labstack/echo"
)

type AppErrorInterface interface {
	NewAppError()
}

type AppError struct {
	UserMessage string
}

type SendErrorMessage struct {
	Message string `json:"message"`
}

func SendError(code int, c echo.Context, err error, msg string) {
	message := SendErrorMessage{msg}
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}

	if !c.Response().Committed() {
		c.JSON(code, message)
	}
	return
}

func (a *AppError) NewAppError(err string) {
	a.UserMessage = err
}
