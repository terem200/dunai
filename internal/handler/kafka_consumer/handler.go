package kafka_consumer

import (
	"github.com/gin-gonic/gin"
	"gitlab.insigit.com/qa/outrunner/internal/handler"
	"gitlab.insigit.com/qa/outrunner/internal/services/kafka"
	"gitlab.insigit.com/qa/outrunner/pkg/logger"
)

const (
	getPath    = "/kafka/:topic/get"
	createPath = "/kafka/:topic/send"
)

type kafkaHandler struct {
	servicesConsumer map[string]kafka.ConsumerService
	servicesProducer map[string]kafka.ProducerService
	logger           logger.ILogger
}

func New(
	sc map[string]kafka.ConsumerService,
	sp map[string]kafka.ProducerService,
	l logger.ILogger) handler.IHandler {
	return &kafkaHandler{
		servicesConsumer: sc,
		servicesProducer: sp,
		logger:           l,
	}
}

func (h kafkaHandler) Register(e *gin.Engine) {
	e.POST(getPath, h.getRecords())
}
