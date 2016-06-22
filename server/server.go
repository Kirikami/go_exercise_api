package server

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kirikami/go_exercise_api/config"
	mw "github.com/kirikami/go_exercise_api/middleware"
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
	server.Use(mw.UseConfig(db, appConfig))

	aut := server.Group("/auth")
	aut.GET("", routes.AutenteficationHandler)
	aut.GET("/callback", routes.ProviderCallback)

	task := server.Group("/task")
	task.Use(middleware.JWT([]byte(appConfig.SigningKey)))

	task.POST("/save", routes.SaveTaskHandler)

	task.PUT("/update/:id", routes.UpdateTaskHandler)
	task.DELETE("/delete/:id", routes.DeleteTaskHandler)
	task.GET("/get/:id", routes.GetTaskHandler)
	task.GET("/get_all_tasks", routes.GetAllTasksHendler)

	task.GET("/status", func(c echo.Context) error {
		return c.String(200, "Status ok")
	})

	server.Run(standard.New(port))
}
