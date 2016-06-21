package main

import (
	"flag"
	"fmt"
	"github.com/kirikami/go_exercise_api/config"
	"github.com/kirikami/go_exercise_api/database"
	mw "github.com/kirikami/go_exercise_api/middleware"
	"github.com/kirikami/go_exercise_api/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"

	_ "github.com/uber-common/zap"
)

var (
	conf     *config.Configuration
	DBConfig config.DatabaseConfig
	port     string
)

func init() {
	configfile := flag.String("config", "config.json", "Config for connection to database")
	conf = config.MustNewConfig(*configfile)
	DBConfig = conf.DatabaseConfig
	port = fmt.Sprintf(":%d", conf.ListenAddress)
}

func main() {
	db := database.MustNewDatabase(DBConfig)

	server := echo.New()

	server.Use(middleware.Recover())
	server.Use(middleware.Logger())
	server.Use(mw.UseConfig(db, conf))

	aut := server.Group("/auth")
	aut.GET("", routes.AutenteficationHandler)
	aut.GET("/callback", routes.ProviderCallback)

	task := server.Group("/task")
	task.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(conf.SigningKey),
		TokenLookup: "Authorization: Bearer" + echo.HeaderAuthorization,
	}))

	task.GET("/write", routes.WriteTaskHandler)

	task.POST("/save", routes.SaveTaskHandler)

	task.PUT("/update/:id", routes.UpdateTaskHandler)
	task.DELETE("/delete/:id", routes.DeleteTaskHandler)
	task.GET("/get/:id", routes.GetTaskHandler)

	r := server.Group("/status")
	r.Use(middleware.JWT([]byte(conf.SigningKey)))
	server.Run(standard.New(port))
}
