package handler

import "github.com/gin-gonic/gin"

type IHandler interface {
	Register(e *gin.Engine)
}
