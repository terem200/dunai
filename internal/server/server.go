package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	mysqlHandler "gitlab.insigit.com/qa/outrunner/internal/handler/mysql"
	"gitlab.insigit.com/qa/outrunner/internal/services/mysql"
	"go.uber.org/zap"
)

// Server - 'outRunner' server struct
type Server struct {
	config *Config
	Engine *gin.Engine
	Logger *zap.Logger
	mySQL  map[string]mysql.Service
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
	err := s.configureMysql()
	if err != nil {
		return err
	}

	s.initRoutes()

	err = s.Engine.Run(s.config.Server.BindAddr)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) initRoutes() {
	s.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	mysql := mysqlHandler.NewHandler(&s.mySQL, s.Logger)
	mysql.Register(s.Engine)

}

func (s *Server) configureMysql() error {
	for k, v := range s.config.MySql {
		if s.mySQL == nil {
			s.mySQL = map[string]mysql.Service{}
		}

		st := mysql.New(&v)

		if err := st.Open(); err != nil {
			e := fmt.Errorf("MySql : %s, \n%w", k, err)
			return e
		}
		s.mySQL[k] = mysql.NewService(st)
		s.Logger.Info(fmt.Sprintf("MySql servise started: %s", k))
	}
	return nil
}
