package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kirikami/go_exercise_api/config"
	"github.com/kirikami/go_exercise_api/database"
	"github.com/kirikami/go_exercise_api/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"

	_ "github.com/uber-common/zap"
)

var c *config.Configuration
var DBConfig config.DatabaseConfig

func init() {
	configfile := flag.String("config", "config.json", "Config for connection to database")
	c = config.MustNewConfig(configfile)
	DBConfig = c.DatabaseConfig
}

func UseConfig(handler echo.HandlerFunc, db *gorm.DB, config *config.Configuration) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) (err error) {
		c.Set("db", db)
		c.Set("config", config)
		return handler(c)
	})
}

func main() {
	db := database.MustNewDatabase(DBConfig)

	server := echo.New()

	server.Use(UseConfig(db, c))
	server.Static("/", "assets")

	server.GET("/login", routes.Login)

	server.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(c.SigningKey),
		TokenLookup: "Authorization: Bearer" + echo.HeaderAuthorization,
	}))
	server.Use(middleware.Recover())
	server.Use(middleware.Logger())

	server.Get("/", routes.HomePageHandler)

	aut := server.Group("/auth")
	aut.Get("", routes.AutenteficationHandler)
	aut.Get("/:provider/callback", routes.ProviderCallback)

	task := server.Group("/task")
	task.Get("/write", routes.WriteTaskHandler)

	//task.Post("/save", UseDB(routes.SaveTaskHandler, db))

	//task.Get("/update/:id", UseDB(routes.UpdateTaskHandler, db))
	//task.Get("/delete/:id", UseDB(routes.DeleteTaskHandler, db))
	//task.Get("/get/:id", UseDB(routes.GetTaskHandler, db))

	r := server.Group("/status")
	r.Use(middleware.JWT([]byte(c.SigningKey)))
	server.Run(standard.New(":" + string(c.ListenAddress)))
	//	server.Run(standard.New(":1223"))
}
