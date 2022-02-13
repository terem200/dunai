package server

import (
	"encoding/json"
	"gitlab.insigit.com/qa/outrunner/internal/services/kafka"
	"gitlab.insigit.com/qa/outrunner/internal/services/mongo"
	"gitlab.insigit.com/qa/outrunner/internal/services/mysql"
	"io/ioutil"
)

type serverConfig struct {
	BindAddr string `json:"port"`
	LogLevel string `json:"logLevel"`
}

// Config - config for 'outRunner' server
// You can override default values for your purposes.
// Default values for server:
//					BindAddr: ":3030"
//					LogLevel: "debug"
type Config struct {
	Server        *serverConfig           `json:"server"`
	MySql         map[string]mysql.Config `json:"mySql"`
	Mongo         map[string]mongo.Config `json:"mongo"`
	KafkaConsumer map[string]kafka.Config `json:"kafkaConsumer"`
}

// NewConfig - initialize new config with default values for 'outRunner' server.
func NewConfig() *Config {
	return &Config{
		Server: &serverConfig{
			BindAddr: ":3030",
			LogLevel: "debug",
		},
	}
}

// ReadConfig
// Reads file by path specified in first argument and write values into target config
func ReadConfig(filePath string, target *Config) error {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, target); err != nil {
		return err
	}
	return nil
}
