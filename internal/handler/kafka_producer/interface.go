package kafka_producer

import (
	"context"
)

type ProducerService interface {
	Send(ctx context.Context, message map[string]interface{}) error
}
