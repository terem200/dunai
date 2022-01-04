package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"gitlab.insigit.com/qa/outrunner/pkg/logger"
	"strconv"
	"time"
)

// Consumer struct
type Consumer struct {
	config     *Config
	logger     logger.ILogger
	connection *kafka.Reader
	messages   []message
}

// New returns new *Consumer struct with passed config
func New(c *Config, l logger.ILogger) consumerConnection {
	return &Consumer{
		config: c,
		logger: l,
	}
}

func (c *Consumer) Connect() error {
	var maxRetries = c.config.MaxRetries
	var waitTime = c.config.WaitRetry
	var i int

	var err error

	for i < maxRetries {
		c.logger.Debug(fmt.Sprintf("Establish kafka consumer connection. Attempt %s", strconv.Itoa(i+1)))
		err = c.connect()
		if err != nil {
			i++
			time.Sleep(time.Duration(waitTime) * time.Millisecond)
			continue
		}
		break
	}

	if err != nil {
		return err
	}

	go c.read(context.Background())
	//if err != nil {
	//	return err
	//}

	return nil
}

func (c *Consumer) connect() error {
	// TODO
	var topic = "my-topic-test"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{c.config.ConnectionURL},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	c.connection = r
	return nil
}

func (c *Consumer) Disconnect() error {
	return c.connection.Close()
}

func (c *Consumer) read(cxt context.Context) {

	for {
		m, err := c.connection.ReadMessage(cxt)
		if err != nil {
			break
		}
		c.logger.Debug(m.Topic, string(m.Value))
	}
}

func (c *Consumer) Get() []message {
	return c.messages
}
