package routes

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"

	"github.com/dgrijalva/jwt-go"
	"github.com/kirikami/go_exercise_api/config"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
)

var (
	appId        = "183514925379752"
	appKey       = "868f22a3890a3267906df8408d44e6ff"
	callbackAddr = "http://localhost:1223/auth/callback?provider=facebook"
)

func init() {
	goth.UseProviders(
		facebook.New(appId, appKey, callbackAddr),
	)
}

func ProviderCallback(c echo.Context) error {

	res := c.Response().(*standard.Response).ResponseWriter
	req := c.Request().(*standard.Request).Request
	_, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway)
	}

	conf := c.Get("config").(*config.Configuration)
	SigningKey := conf.SigningKey
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["admin"] = true
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(SigningKey))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"token": t})
}

func AutenteficationHandler(c echo.Context) error {
	gothic.BeginAuthHandler(c.Response().(*standard.Response).ResponseWriter, c.Request().(*standard.Request).Request)
	return c.String(http.StatusOK, "Autentefication with facebook")
}
