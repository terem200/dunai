package kafka_consumer

import (
	"context"
)

type Message = map[string]interface{}
type Query = map[string]interface{}

type ConsumerConnection interface {
	Connect(topic string) error
	Disconnect() error
	Get() []Message
}
type ProducerConnection interface {
	Connect() error
	Disconnect() error
	Send() (bool, error)
}

type ConsumerService interface {
	Get(ctx context.Context, query map[string]interface{}) ([]map[string]interface{}, error)
}

type ProducerService interface {
	Send(ctx context.Context, query map[string]interface{}) (bool, error)
}
