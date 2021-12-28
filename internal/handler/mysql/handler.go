package mysql

import (
	"github.com/gin-gonic/gin"
	"gitlab.insigit.com/qa/outrunner/internal/services/mysql"
	"go.uber.org/zap"
)

const (
	//componentsPath = "/api/v1/mysql/components"
	getPath    = "/api/v1/mysql/:dbName/get"
	createPath = "/api/v1/mysql/:dbName/modify"
)

type handler struct {
	services map[string]mysql.Service
	logger   *zap.Logger
}

// NewHandler - initializes and returns new handler for MySQL services.
// It doesn't register routes to serve requests relates to MySQL.
// Next step you need to call Register func available from returned handler. */
func NewHandler(s *map[string]mysql.Service, l *zap.Logger) *handler {
	return &handler{
		services: *s,
		logger:   l,
	}
}

// Register - register routes to serve requests relates to MySQL.
func (h *handler) Register(e *gin.Engine) {
	e.POST(getPath, h.getRecords())
	e.POST(createPath, h.modify())
}
