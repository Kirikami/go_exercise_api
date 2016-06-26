package main

import (
	"flag"
	"github.com/kirikami/go_exercise_api/config"
	"github.com/kirikami/go_exercise_api/database"
	"github.com/kirikami/go_exercise_api/routes"
	"github.com/kirikami/go_exercise_api/server"
)

var app routes.ApiV1Handler

func init() {
	initConfig()
	initDatabase()
}

func initConfig() {
	configfile := flag.String("config", "config.json", "Config for connection to database")
	flag.Parse()
	app.Config = config.MustNewConfig(*configfile)
}

func initDatabase() {
	app.DB = database.MustNewDatabase(app.Config.DatabaseConfig)
}

func main() {
	server.StartServer(app)
}
