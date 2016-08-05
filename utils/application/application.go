package application

import (
	"flag"

	"github.com/jinzhu/gorm"

	"github.com/kirikami/go_exercise_api/config"
	"github.com/kirikami/go_exercise_api/database"
)

type App interface {
	InitConfiguration()
	InitDatabase()
}

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
	a.Database = database.MustNewDatabase(a.Configuration.DatabaseConfig)
}
