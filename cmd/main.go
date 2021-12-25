package main

import (
	"gitlab.insigit.com/qa/outrunner/internal/server"
	"go.uber.org/zap"
)

const configPath = "config/config.json"

func main() {
	logger, _ := zap.NewProduction()
	serverCfg := server.NewConfig()

	if err := server.ReadConfig(configPath, serverCfg); err != nil {
		logger.Fatal("Config read failed...", zap.String("Error : ", err.Error()))
	}

	s := server.New(serverCfg, logger)

	err := s.Run()
	if err != nil {
		logger.Fatal("Server start failed...", zap.String("Error : ", err.Error()))
	}
}
