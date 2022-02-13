package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	kafkaGo "github.com/segmentio/kafka-go"
	"gitlab.insigit.com/qa/outrunner/internal/handler/kafka_consumer"
	"gitlab.insigit.com/qa/outrunner/pkg/logger"
	"strconv"
	"time"
)

// Consumer struct
type Consumer struct {
	config     *Config
	logger     logger.ILogger
	connection *kafkaGo.Reader
	messages   []kafka_consumer.Message
}

// New returns new *Consumer struct with passed config
func New(c *Config, l logger.ILogger) kafka_consumer.ConsumerConnection {
	return &Consumer{
		config: c,
		logger: l,
	}
}

func (c *Consumer) Connect(topic string) error {
	var maxRetries = c.config.MaxRetries
	var waitTime = c.config.WaitRetry
	var i int

	var err error

	for i < maxRetries {
		c.logger.Debug(fmt.Sprintf("Establish kafka consumer connection. Attempt %s", strconv.Itoa(i+1)))
		err = c.connect(topic)
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

	return nil
}

func (c *Consumer) connect(topic string) error {
	r := kafkaGo.NewReader(kafkaGo.ReaderConfig{
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
			c.logger.Debug("READING MESSAGE ERROR::", err.Error())
			break
		}
		c.logger.Debug(m.Topic, string(m.Value))

		var parsed kafka_consumer.Message
		err = json.Unmarshal(m.Value, &parsed)
		if err != nil {
			c.logger.Debug("PARSING MESSAGE ERROR::", err.Error())
		}

		c.messages = append(c.messages, parsed)
		c.logger.Debug("MESSAGES", c.messages)
	}
}

func (c *Consumer) Get() []kafka_consumer.Message {
	return c.messages
}
