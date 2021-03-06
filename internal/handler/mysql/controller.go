package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	customResponse "gitlab.insigit.com/qa/outrunner/internal/handler"
	"io/ioutil"
)

func (h *mySqlHandler) getRecords() gin.HandlerFunc {
	return func(c *gin.Context) {
		dbName := c.Param("dbName")
		if h.services[dbName] == nil {
			e := fmt.Errorf("mysql connection for '%s' not configured", dbName)
			customResponse.RequestErr(c, e)
			h.logger.Error(e.Error())
			return
		}

		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			customResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		var dto dto
		if err := json.Unmarshal(b, &dto); err != nil {
			customResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		data, err := h.services[dbName].Get(c, dto.Query)
		if err != nil {
			customResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		customResponse.SuccessData(c, data)
	}
}

func (h *mySqlHandler) modify() gin.HandlerFunc {
	return func(c *gin.Context) {
		dbName := c.Param("dbName")
		if h.services[dbName] == nil {
			e := fmt.Errorf("mysql connection for '%s' not configured", dbName)
			customResponse.RequestErr(c, e)
			h.logger.Error(e.Error())
			return
		}

		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			customResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		var dto dto
		if err := json.Unmarshal(b, &dto); err != nil {
			customResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		_, err = h.services[dbName].Create(c, dto.Query)
		if err != nil {
			customResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		customResponse.SuccessOK(c)
	}
}
