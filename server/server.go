package server

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kirikami/go_exercise_api/config"
	"github.com/kirikami/go_exercise_api/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func StartServer(db *gorm.DB, appConfig *config.Configuration) {
	port := fmt.Sprintf(":%d", appConfig.ListenAddress)
	server := echo.New()

	server.Use(middleware.Recover())
	server.Use(middleware.Logger())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderAccessControlAllowOrigin},
	}))

	api := routes.ApiV1Handler{db, appConfig}

	v1 := server.Group("/v1")

	aut := v1.Group("/auth")
	aut.GET("", api.AutenteficationHandler)
	aut.GET("/callback", api.ProviderCallback)

	task := v1.Group("/task")
	task.Use(middleware.JWT([]byte(appConfig.SigningKey)))

	task.POST("/save", api.SaveTaskHandler)

	task.PUT("/update/:id", api.UpdateTaskHandler)
	task.DELETE("/delete/:id", api.DeleteTaskHandler)
	task.GET("/get/:id", api.GetTaskHandler)
	task.GET("/get_all_tasks", api.GetAllTasksHendler)

	task.GET("/status", func(c echo.Context) error {
		return c.String(200, "Status ok")
	})

	server.Run(standard.New(port))
}
