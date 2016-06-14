package database

import (
	"errors"
	"fmt"
	//log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kirikami/go_exercise_api/config"
)

var (
	ErrDbConnect = errors.New("Failed connect to database")
)

func NewDatabase(c config.DatabaseConfig) (*gorm.DB, error) {
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", c.Username, c.Password, c.DatabaseUri, c.Port, c.DBName)
	db, err := gorm.Open("mysql", dbConnection)

	if err != nil {
		return nil, ErrDbConnect
	}

	return db, nil

}

func MustNewDatabase(c config.DatabaseConfig) *gorm.DB {
	db, err := NewDatabase(c)

	if err != nil {
		//log.Fatalf("Connection problem: %s", err)
	}

	return db

}
