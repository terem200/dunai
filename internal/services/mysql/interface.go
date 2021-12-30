package mysql

import "context"

//Service declares an interface for work with MySQL storage
type Service interface {
	Get(ctx context.Context, stmt string) ([]map[string]interface{}, error)
	Create(ctx context.Context, stmt string) (bool, error)
}

type storage interface {
	Open() error
	Close() error
	Get(ctx context.Context, stmt string) ([]map[string]interface{}, error)
	Create(ctx context.Context, stmt string) (bool, error)
}
