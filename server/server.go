package server

import (
	"fmt"
	"github.com/kirikami/go_exercise_api/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func StartServer(app routes.ApiV1Handler) {
	port := fmt.Sprintf(":%d", app.Config.ListenAddress)
	server := echo.New()

	server.Use(middleware.Recover())
	server.Use(middleware.Logger())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderAccessControlAllowOrigin},
	}))

	api := routes.ApiV1Handler{app.DB, app.Config}

	v1 := server.Group("/v1")

	v1.GET("/tasks", api.GetAllTasksHendler, middleware.JWT([]byte(app.Config.SigningKey)))

	aut := v1.Group("/auth")
	aut.GET("", api.AutenteficationHandler)
	aut.GET("/callback", api.ProviderCallback)

	task := v1.Group("/task")
	task.Use(middleware.JWT([]byte(app.Config.SigningKey)))

	task.POST("", api.SaveTaskHandler)

	task.PUT("/:id", api.UpdateTaskHandler)
	task.DELETE("/:id", api.DeleteTaskHandler)
	task.GET("/:id", api.GetTaskHandler)

	task.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "Status ok")
	})

	server.Run(standard.New(port))
}
