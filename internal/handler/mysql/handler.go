package mysql

import (
	"github.com/gin-gonic/gin"
	"gitlab.insigit.com/qa/outrunner/internal/handler"
	"gitlab.insigit.com/qa/outrunner/internal/services/mysql"
	"gitlab.insigit.com/qa/outrunner/pkg/logger"
)

const (
	//componentsPath = "/api/v1/mysql/components"
	getPath    = "/mysql/:dbName/get"
	createPath = "/mysql/:dbName/modify"
)

type mySqlHandler struct {
	services map[string]mysql.Service
	logger   logger.ILogger
}

// NewHandler - initializes and returns new mySqlHandler for MySQL services.
// It doesn't register routes to serve requests relates to MySQL.
// Next step you need to call Register func available from returned mySqlHandler. */
func NewHandler(s *map[string]mysql.Service, l logger.ILogger) handler.IHandler {
	return &mySqlHandler{
		services: *s,
		logger:   l,
	}
}

// Register - register routes to serve requests relates to MySQL.
func (h *mySqlHandler) Register(e *gin.Engine) {
	e.POST(getPath, h.getRecords())
	e.POST(createPath, h.modify())
}
