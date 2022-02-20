package kafka_producer

import (
	"context"
	"encoding/json"
	"fmt"
	kafkaGo "github.com/segmentio/kafka-go"
	"gitlab.insigit.com/qa/outrunner/pkg/logger"
	"strconv"
	"time"
)

// Producer struct
type Producer struct {
	config     *Config
	logger     logger.ILogger
	connection *kafkaGo.Writer
}

// New returns new *Producer struct with passed config
func New(c *Config, l logger.ILogger) ProducerConnection {
	return &Producer{
		config: c,
		logger: l,
	}
}

func (p *Producer) Connect(topic string) error {
	var maxRetries = p.config.MaxRetries
	var waitTime = p.config.WaitRetry
	var i int

	var err error

	for i < maxRetries {
		p.logger.Debug(fmt.Sprintf("Establish kafka producer connection. Attempt %s", strconv.Itoa(i+1)))
		err = p.connect(topic)
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

	return nil
}

func (p *Producer) connect(topic string) error {
	r := kafkaGo.Writer{
		Addr:     kafkaGo.TCP(p.config.BrokerURL),
		Topic:    topic,
		Balancer: &kafkaGo.LeastBytes{},
	}

	p.connection = &r
	return nil
}

func (p *Producer) Disconnect() error {
	return p.connection.Close()
}

func (p *Producer) Send(cxt context.Context, message Message) error {

	m, err := json.Marshal(message)
	if err != nil {
		p.logger.Error("Failed to marshal message", message)
		return err
	}

	err = p.connection.WriteMessages(cxt,
		kafkaGo.Message{
			Key:   []byte("DUNAI"),
			Value: m,
		},
	)
	if err != nil {
		p.logger.Debug("SEND MESSAGE ERROR::", err.Error())
		return err
	}
	p.logger.Debug("Message SEND", message)
	return nil
}
