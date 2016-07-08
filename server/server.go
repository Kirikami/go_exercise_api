package server

import (
	"fmt"
	mw "github.com/kirikami/go_exercise_api/middleware"
	t "github.com/kirikami/go_exercise_api/server/handlers/task"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func StartServer(app t.ApiV1Handler) {
	server := echo.New()

	server.Use(middleware.Recover())
	server.Use(mw.Logger())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderAccessControlAllowOrigin},
	}))

	api := t.ApiV1Handler{app.DB, app.Config}

	server.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Status ok")
	})

	v1 := server.Group("/v1")

	v1.GET("/tasks", api.GetAllTasksHendler, mw.JWT(app.Config.SigningKey))

	aut := v1.Group("/auth")
	aut.GET("", api.AutenteficationHandler)
	aut.GET("/callback", api.ProviderCallback)

	task := v1.Group("/task")
	task.Use(mw.JWT(app.Config.SigningKey))

	task.POST("", api.SaveTaskHandler)

	task.PUT("/:id", api.UpdateTaskHandler)
	task.DELETE("/:id", api.DeleteTaskHandler)
	task.GET("/:id", api.GetTaskHandler)

	server.Run(standard.New(fmt.Sprintf(":%d", app.Config.ListenAddress)))
}
