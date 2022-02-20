package kafka_producer

import (
	"context"
	"gitlab.insigit.com/qa/outrunner/internal/handler/kafka_producer"
)

type producerService struct {
	conn ProducerConnection
}

func NewService(c ProducerConnection) kafka_producer.ProducerService {
	return &producerService{
		conn: c,
	}
}

func (cs *producerService) Send(ctx context.Context, m Message) error {
	err := cs.conn.Send(ctx, m)
	if err != nil {
		return err
	}

	return nil
}
