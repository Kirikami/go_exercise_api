package database

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kirikami/go_exercise_api/config"
)

type GormLogger struct{}

var (
	ErrDbConnect = errors.New("Failed connect to database")
)

func NewDatabase(c config.DatabaseConfig) (*gorm.DB, error) {
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", c.Username, c.Password, c.DatabaseUri, c.Port, c.DBName)
	db, err := gorm.Open("mysql", dbConnection)
	db.LogMode(true)
	db.SetLogger(&GormLogger{})
	db.SingularTable(true)

	if err != nil {
		return nil, ErrDbConnect
	}

	return db, nil
}

func MustNewDatabase(c config.DatabaseConfig) *gorm.DB {
	db, err := NewDatabase(c)

	if err != nil {
		panic(err)
	}

	return db
}

func (*GormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		log.WithFields(log.Fields{"module": "gorm", "type": "sql"}).Print(v[3])
	}
	if v[0] == "log" {
		log.WithFields(log.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}
