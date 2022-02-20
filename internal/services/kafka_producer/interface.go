package kafka_producer

import "context"

type Message = map[string]interface{}

type ProducerConnection interface {
	Connect(topic string) error
	Disconnect() error
	Send(cxt context.Context, message Message) error
}
