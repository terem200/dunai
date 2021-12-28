package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	sendResponse "gitlab.insigit.com/qa/outrunner/internal/handler"
	"io/ioutil"
)

func (h *handler) getRecords() gin.HandlerFunc {
	return func(c *gin.Context) {
		dbName := c.Param("dbName")
		if h.services[dbName] == nil {
			e := fmt.Errorf("mysql connection for '%s' not configured", dbName)
			sendResponse.RequestErr(c, e)
			h.logger.Error(e.Error())
			return
		}

		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			sendResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		var dto dto
		if err := json.Unmarshal(b, &dto); err != nil {
			sendResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		data, err := h.services[dbName].Get(c, dto.Query)
		if err != nil {
			sendResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		sendResponse.SuccessData(c, data)
	}
}

func (h *handler) modify() gin.HandlerFunc {
	return func(c *gin.Context) {
		dbName := c.Param("dbName")
		if h.services[dbName] == nil {
			e := fmt.Errorf("mysql connection for '%s' not configured", dbName)
			sendResponse.RequestErr(c, e)
			h.logger.Error(e.Error())
			return
		}

		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			sendResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		var dto dto
		if err := json.Unmarshal(b, &dto); err != nil {
			sendResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		_, err = h.services[dbName].Create(c, dto.Query)
		if err != nil {
			sendResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		sendResponse.SuccessOK(c)
	}
}
