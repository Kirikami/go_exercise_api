package main

import (
	"flag"
	"github.com/kirikami/go_exercise_api/config"
	"github.com/kirikami/go_exercise_api/database"
	"github.com/kirikami/go_exercise_api/server"
)

type Application struct {
	Configuration *config.Configuration
	Database      *gorm.DB
}

func (a *Application) InitConfiguration() {
	configfile := flag.String("config", "config.json", "Config for connection to database")
	flag.Parse()
	a.Configuration = config.MustNewConfig(*configfile)
}

func (a *Application) InitDatabase() {
	app.Database = database.MustNewDatabase(app.Config.DatabaseConfig)
}

func main() {
	app = Application{}
	app.InitConfiguration
	app.InitDatabase
	server.StartServer(app)
}
