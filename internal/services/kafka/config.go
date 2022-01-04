package kafka

type Config struct {
	ConnectionURL string `json:"connectionUrl"`
	MaxRetries    int    `json:"maxRetries"`
	WaitRetry     int    `json:"waitRetry"`
}
