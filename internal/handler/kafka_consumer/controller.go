package kafka_consumer

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	customResponse "gitlab.insigit.com/qa/outrunner/internal/handler"
	"io/ioutil"
)

func (h *kafkaHandler) getRecords() gin.HandlerFunc {
	return func(c *gin.Context) {
		topic := c.Param("topic")
		if h.servicesConsumer[topic] == nil {
			e := fmt.Errorf("kafka cosumer connection for topic '%s' not configured", topic)
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

		var dto dtoGet
		if err := json.Unmarshal(b, &dto); err != nil {
			customResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		data, err := h.servicesConsumer[topic].Get(c, dto.Query)
		if err != nil {
			customResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		customResponse.SuccessData(c, data)
	}
}
