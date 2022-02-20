package kafka_consumer

import (
	"context"
)

type ConsumerService interface {
	Get(ctx context.Context, query map[string]interface{}) ([]map[string]interface{}, error)
}
