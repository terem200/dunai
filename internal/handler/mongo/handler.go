package mongo

import (
	"github.com/gin-gonic/gin"
	handler "gitlab.insigit.com/qa/outrunner/internal/handler"
	"gitlab.insigit.com/qa/outrunner/internal/services/mongo"
	"gitlab.insigit.com/qa/outrunner/pkg/logger"
)

const (
	//componentsPath = "/api/v1/mysql/components"
	getPath    = "/mongo/:dbName/get"
	createPath = "/mongo/:dbName/modify"
)

type mongoHandler struct {
	services map[string]mongo.Service
	logger   logger.ILogger
}

// NewHandler - initializes and returns new handler for Mongo services.
// It doesn't register routes to serve requests relates to MySQL.
// Next step you need to call Register func available from returned handler. */
func NewHandler(s *map[string]mongo.Service, l logger.ILogger) handler.IHandler {
	return &mongoHandler{
		services: *s,
		logger:   l,
	}
}

// Register - register routes to serve requests relates to Mongo.
func (h *mongoHandler) Register(e *gin.Engine) {
	e.POST(getPath, h.getRecords())
	e.POST(createPath, h.modify())
}
