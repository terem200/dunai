package mongo

// Config provides options to establish connection to mongo db
type Config struct {
	ConnectionURL string `json:"connectionUrl"`
	Database      string `json:"database"`
	MaxRetries    int    `json:"maxRetries"`
	WaitRetry     int    `json:"waitRetry"`
}
