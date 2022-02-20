package kafka_producer

import (
	"github.com/gin-gonic/gin"
	"gitlab.insigit.com/qa/outrunner/internal/handler"
	"gitlab.insigit.com/qa/outrunner/pkg/logger"
)

const (
	createPath = "/kafka/:topic/send"
)

type kafkaProducerHandler struct {
	servicesProducer map[string]ProducerService
	logger           logger.ILogger
}

func NewHandler(sp map[string]ProducerService, l logger.ILogger) handler.IHandler {
	return &kafkaProducerHandler{
		servicesProducer: sp,
		logger:           l,
	}
}

func (h kafkaProducerHandler) Register(e *gin.Engine) {
	e.POST(createPath, h.writeRecord())
}
