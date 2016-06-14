package config

import (
	"encoding/json"
	"errors"
	"fmt"
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
		fmt.Print("parsing config file", err.Error())
		//log.Fatal(err)
		return nil, err
	}

	fmt.Print(configuration)
	return &configuration, nil

}

func MustNewConfig(configfile *string) *Configuration {
	config, err := NewConfig(*configfile)

	if err != nil {
		//log.Fatalf("Cant parse config file: %s", err)
	}

	return config

}
