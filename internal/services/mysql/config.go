package mysql

// Config provides options to establish connection to MySQL db
type Config struct {
	ConnectionURL string `json:"connectionUrl"`
	MaxRetries    int    `json:"maxRetries"`
	WaitRetry     int    `json:"waitRetry"`
}
