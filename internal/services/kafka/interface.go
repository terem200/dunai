package kafka

import "context"

type consumerConnection interface {
	Connect() error
	Disconnect() error
	Get() ([]message, error)
}
type producerConnection interface {
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
