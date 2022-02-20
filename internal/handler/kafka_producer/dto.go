package kafka_producer

type dtoSend struct {
	Message map[string]interface{} `json:"message"`
}
