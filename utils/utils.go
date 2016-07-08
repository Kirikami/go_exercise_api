package utils

import (
	"github.com/kirikami/go_exercise_api/config"
	"github.com/kirikami/go_exercise_api/database"
	"strconv"
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

func ParseIdInt64FromString(s string) (int64, error) {
	result, err := strconv.ParseInt(s, 10, 64)

	return result, err
}
