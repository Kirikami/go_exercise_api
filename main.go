package main

import (
	"flag"
	"github.com/jinzhu/gorm"
	"github.com/kirikami/go_exercise_api/config"
	"github.com/kirikami/go_exercise_api/database"
	"github.com/kirikami/go_exercise_api/server"
)

var appConfig *config.Configuration
var db *gorm.DB

func init() {
	initConfig()
	initDatabase()
}

func initConfig() {
	configfile := flag.String("config", "config.json", "Config for connection to database")
	appConfig = config.MustNewConfig(*configfile)
}

func initDatabase() {
	db = database.MustNewDatabase(appConfig.DatabaseConfig)
}

func main() {
	server.StartServer(db, appConfig)
}
