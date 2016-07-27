package task

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/kirikami/go_exercise_api/config"
	u "github.com/kirimaki/go_exercise_api/utils"
)

type ApiV1Handler struct {
	Configuration *config.Configuration
	Database      *gorm.DB
}

var (
	ErrInvalidTaskId          string = "Id should be numeric"
	ErrInternalDatabase       string = "Database internal error"
	ErrIncorrectData          string = "Incorrect data or format of data sent"
	StatusUnprocessableEntity int    = 422
)

func SendBadRequest(userError u.AppError, internalError error) error {
	NewErrorWIthStatus(http.StatusBadRequest, userError, internalError)
}

func EchoErrorHandler(err error, c *echo.Context) {

	appError, ok := err.(AppErrorWithStatus)
	if !ok {
		panic("Invalid response error type")
	}
}
