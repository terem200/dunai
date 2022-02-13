package kafka_consumer

import (
	"github.com/gin-gonic/gin"
	"gitlab.insigit.com/qa/outrunner/internal/handler"
	"gitlab.insigit.com/qa/outrunner/pkg/logger"
)

const (
	getPath    = "/kafka/:topic/get"
	createPath = "/kafka/:topic/send"
)

type kafkaHandler struct {
	servicesConsumer map[string]ConsumerService
	logger           logger.ILogger
}

func NewHandler(
	sc map[string]ConsumerService,
	l logger.ILogger) handler.IHandler {
	return &kafkaHandler{
		servicesConsumer: sc,
		logger:           l,
	}
}

func (h kafkaHandler) Register(e *gin.Engine) {
	e.POST(getPath, h.getRecords())
}
