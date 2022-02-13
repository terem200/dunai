package mongo

import "context"

type service struct {
	storage storage
}

// NewService - initialises new mongodb service
func NewService(s storage) Service {
	return &service{
		storage: s,
	}
}

func (s *service) Get(ctx context.Context, q QueryGet) ([]map[string]interface{}, error) {
	return s.storage.Get(ctx, q)
}

func (s *service) Create(ctx context.Context, q QueryInsert) (bool, error) {
	return s.storage.Create(ctx, q)
}
