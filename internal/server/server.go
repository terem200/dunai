package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	mongoHandler "gitlab.insigit.com/qa/outrunner/internal/handler/mongo"
	mysqlHandler "gitlab.insigit.com/qa/outrunner/internal/handler/mysql"
	"gitlab.insigit.com/qa/outrunner/internal/services/kafka"
	"gitlab.insigit.com/qa/outrunner/internal/services/mongo"
	"gitlab.insigit.com/qa/outrunner/internal/services/mysql"
	"gitlab.insigit.com/qa/outrunner/pkg/logger"
)

// Server - 'outRunner' server struct
type Server struct {
	config         *Config
	Engine         *gin.Engine
	Logger         logger.ILogger
	mySQL          map[string]mysql.Service
	mongo          map[string]mongo.Service
	kafkaConsumers map[string]kafka.ConsumerService
}

// New - initialize new connector server
func New(config *Config, logger logger.ILogger) *Server {
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

	err = s.configureMongo()
	if err != nil {
		return err
	}

	err = s.configureKafkaConsumers()
	if err != nil {
		return err
	}

	s.initRoutes()

	err = s.Engine.Run(s.config.Server.BindAddr)
	if err != nil {
		return err
	}

	s.Logger.Info("Server started successfully")
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

	mongo := mongoHandler.NewHandler(&s.mongo, s.Logger)
	mongo.Register(s.Engine)
}

func (s *Server) configureMysql() error {
	for k, v := range s.config.MySql {
		if s.mySQL == nil {
			s.mySQL = map[string]mysql.Service{}
		}

		st := mysql.New(&v, s.Logger)

		if err := st.Open(); err != nil {
			e := fmt.Errorf("MySql : %s, \n%w", k, err)
			return e
		}
		s.mySQL[k] = mysql.NewService(st)
		s.Logger.Info(fmt.Sprintf("MySql servise started: %s", k))
	}
	return nil
}

func (s *Server) configureMongo() error {
	for k, v := range s.config.Mongo {
		if s.mongo == nil {
			s.mongo = map[string]mongo.Service{}
		}

		st := mongo.New(&v, s.Logger)

		if err := st.Open(); err != nil {
			e := fmt.Errorf("Mongo : %s, \n%w", k, err)
			return e
		}
		s.mongo[k] = mongo.NewService(st)
		s.Logger.Info(fmt.Sprintf("Mongo servise started: %s", k))
	}
	return nil
}

func (s *Server) configureKafkaConsumers() error {
	for k, v := range s.config.KafkaConsumer {
		if s.kafkaConsumers == nil {
			s.kafkaConsumers = map[string]kafka.ConsumerService{}
		}

		st := kafka.New(&v, s.Logger)

		if err := st.Connect(); err != nil {
			e := fmt.Errorf("kafka consumer : %s, \n%w", k, err)
			return e
		}
		s.kafkaConsumers[k] = kafka.NewService(st)
		s.Logger.Info(fmt.Sprintf("kafka consumer started: %s", k))
	}
	return nil
}
