package mongo

import (
	"github.com/gin-gonic/gin"
	"gitlab.insigit.com/qa/outrunner/internal/services/mongo"
	"gitlab.insigit.com/qa/outrunner/pkg/logger"
)

const (
	//componentsPath = "/api/v1/mysql/components"
	getPath    = "/mongo/:dbName/get"
	createPath = "/mongo/:dbName/modify"
)

type handler struct {
	services map[string]mongo.Service
	logger   logger.ILogger
}

// NewHandler - initializes and returns new handler for Mongo services.
// It doesn't register routes to serve requests relates to MySQL.
// Next step you need to call Register func available from returned handler. */
func NewHandler(s *map[string]mongo.Service, l logger.ILogger) *handler {
	return &handler{
		services: *s,
		logger:   l,
	}
}

// Register - register routes to serve requests relates to Mongo.
func (h *handler) Register(e *gin.Engine) {
	e.POST(getPath, h.getRecords())
	e.POST(createPath, h.modify())
}
