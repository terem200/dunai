package kafka_producer

type Config struct {
	BrokerURL  string `json:"brokerUrl"`
	MaxRetries int    `json:"maxRetries"`
	WaitRetry  int    `json:"waitRetry"`
}
