package kafka_producer

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	customResponse "gitlab.insigit.com/qa/outrunner/internal/handler"
	"io/ioutil"
)

func (h *kafkaProducerHandler) writeRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		topic := c.Param("topic")
		if h.servicesProducer[topic] == nil {
			e := fmt.Errorf("kafka producer connection for topic '%s' not configured", topic)
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

		var dto dtoSend
		if err := json.Unmarshal(b, &dto); err != nil {
			customResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		err = h.servicesProducer[topic].Send(c, dto.Message)
		if err != nil {
			customResponse.InternalErr(c, err)
			h.logger.Error(err.Error())
			return
		}

		customResponse.SuccessData(c, "ok")
	}
}
