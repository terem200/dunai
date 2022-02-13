package mongo

import "context"

//Service declares an interface for work with Mongo storage
type Service interface {
	Get(ctx context.Context, query QueryGet) ([]map[string]interface{}, error)
	Create(ctx context.Context, query QueryInsert) (bool, error)
}

type storage interface {
	Open() error
	Close() error
	Get(ctx context.Context, query QueryGet) ([]map[string]interface{}, error)
	Create(ctx context.Context, query QueryInsert) (bool, error)
}

// QueryGet describes to perform retrieving mongo records
type QueryGet struct {
	Collection string                 `json:"collection"`
	Query      map[string]interface{} `json:"query"`
}

// QueryInsert describes interface to insert mongo records
type QueryInsert struct {
	Collection string                   `json:"collection"`
	Query      []map[string]interface{} `json:"query"`
}
