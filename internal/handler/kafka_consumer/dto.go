package kafka_consumer

type dtoGet struct {
	Topic string                 `json:"topic"`
	Query map[string]interface{} `json:"query"`
}
