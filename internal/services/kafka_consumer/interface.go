package kafka_consumer

type Message = map[string]interface{}
type Query = map[string]interface{}

type ConsumerConnection interface {
	Connect(topic string) error
	Disconnect() error
	Get() []Message
}
