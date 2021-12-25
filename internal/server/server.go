package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Server - 'outRunner' server struct
type Server struct {
	config *Config
	Engine *gin.Engine
	Logger *zap.Logger
}

// New - initialize new connector server
func New(config *Config, logger *zap.Logger) *Server {
	return &Server{
		config: config,
		Logger: logger,
		Engine: gin.Default(),
	}
}

// Run outRunner server
func (s *Server) Run() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.initRoutes()

	err := s.Engine.Run(s.config.Server.BindAddr)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) configureLogger() error {
	return nil
}

func (s *Server) initRoutes() {
	//...
}
