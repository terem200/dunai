package kafka

import (
	"context"
	"gitlab.insigit.com/qa/outrunner/internal/handler/kafka_consumer"
)

type consumerService struct {
	conn kafka_consumer.ConsumerConnection
}

func NewService(c kafka_consumer.ConsumerConnection) kafka_consumer.ConsumerService {
	return &consumerService{
		conn: c,
	}
}

func (cs *consumerService) Get(ctx context.Context, q kafka_consumer.Query) ([]kafka_consumer.Message, error) {
	messages := make([]kafka_consumer.Message, 0)

	km := cs.conn.Get()

	if len(q) == 0 {
		return km, nil
	}

	for _, m := range km {
		for k := range q {
			if m[k] == q[k] {
				messages = append(messages, m)
			}
		}
	}
	return messages, nil
}
