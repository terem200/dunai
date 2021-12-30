package main

import (
	"fmt"
	"gitlab.insigit.com/qa/outrunner/internal/server"
	"gitlab.insigit.com/qa/outrunner/pkg/logger"
	"log"
)

const configPath = "config/config.json"

func main() {
	serverCfg := server.NewConfig()
	if err := server.ReadConfig(configPath, serverCfg); err != nil {
		log.Fatal(fmt.Sprintf("Config read failed...%s", err.Error()))
	}

	logger := logger.New(serverCfg.Server.LogLevel)

	s := server.New(serverCfg, logger)

	err := s.Run()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Server start failed... %s", err.Error()))
	}
}
