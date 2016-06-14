package routes

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kirikami/go_api_with_jwt/config"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
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
