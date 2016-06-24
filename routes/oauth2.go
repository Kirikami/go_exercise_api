package routes

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/kirikami/go_exercise_api/config"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
)

func (h ApiV1Handler) ProviderCallback(c echo.Context) error {

	res := c.Response().(*standard.Response).ResponseWriter
	req := c.Request().(*standard.Request).Request
	_, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["admin"] = true
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(h.Config.SigningKey))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"token": t})
}

func (h ApiV1Handler) AutenteficationHandler(c echo.Context) error {
	fb := h.Config.FacebookConfig
	goth.UseProviders(
		facebook.New(fb.ID, fb.Key, fb.CallbackAddress),
	)
	gothic.BeginAuthHandler(c.Response().(*standard.Response).ResponseWriter, c.Request().(*standard.Request).Request)
	return c.NoContent(http.StatusNoContent)
}
