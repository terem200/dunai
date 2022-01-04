package kafka

import "context"

type message = map[string]interface{}
type query = map[string]interface{}

type consumerService struct {
	conn consumerConnection
}

func NewService(c consumerConnection) ConsumerService {
	return &consumerService{
		conn: c,
	}
}

func (cs *consumerService) Get(ctx context.Context, q query) ([]message, error) {
	messages := make([]message, 0)

	km, err := cs.conn.Get()
	if err != nil {
		return nil, err
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
