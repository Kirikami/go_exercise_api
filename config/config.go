package config

import (
	"encoding/json"
	"errors"
	//log "github.com/Sirupsen/logrus"
	"os"
)

type Configuration struct {
	ListenAddress  int
	DatabaseConfig DatabaseConfig
	SigningKey     string
}
type DatabaseConfig struct {
	DatabaseUri string
	Username    string
	Password    string
	Port        int
	DBName      string
}

var (
	ErrCantFindConfig = errors.New("Config file is missing")
)

func NewConfig(configfile string) (*Configuration, error) {
	if _, err := os.Stat(configfile); os.IsNotExist(err) {
		return nil, ErrCantFindConfig
	}

	configuration := Configuration{}

	file, _ := os.Open(configfile)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}

	return &configuration, nil

}

func MustNewConfig(configfile string) *Configuration {
	config, err := NewConfig(configfile)

	if err != nil {
		panic(err)
	}

	return config

}
