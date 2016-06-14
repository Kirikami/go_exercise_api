package routes

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
)

func init() {
	goth.UseProviders(
		facebook.New("183514925379752", "868f22a3890a3267906df8408d44e6ff", "http://localhost:1223/auth/facebook/callback"),
	)
}

func ProviderCallback(c echo.Context) error {

	res := c.Response().(*standard.Response).ResponseWriter
	req := c.Request().(*standard.Request).Request
	_, err := gothic.CompleteUserAuth(res, req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway)
	}

	return c.String(http.StatusOK, "ProviderCallback")
}

func HomePageHandler(c echo.Context) error {
	t, _ := template.New("index").Parse(indexTemplate)
	t.Execute(c.Response().(*standard.Response).ResponseWriter, nil)
	return c.String(http.StatusOK, "Welcome!")
}

func AutenteficationHandler(c echo.Context) error {
	gothic.BeginAuthHandler(c.Response().(*standard.Response).ResponseWriter, c.Request().(*standard.Request).Request)
	return c.String(http.StatusOK, "Autentefication with facebook")
}

var indexTemplate = `<p><a href="/auth?provider=facebook">Log in with Facebook</a></p>`
