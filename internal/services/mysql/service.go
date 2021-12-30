package mysql

import "context"

type service struct {
	storage storage
}

func NewService(s storage) Service {
	return &service{
		storage: s,
	}
}

func (s *service) Get(ctx context.Context, stmt string) ([]map[string]interface{}, error) {
	return s.storage.Get(ctx, stmt)
}

func (s *service) Create(ctx context.Context, stmt string) (bool, error) {
	return s.storage.Create(ctx, stmt)
}
