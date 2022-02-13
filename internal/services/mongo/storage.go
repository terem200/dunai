package mongo

import (
	"context"
	"fmt"
	"gitlab.insigit.com/qa/outrunner/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

// Mongo struct
type Mongo struct {
	config *Config
	logger logger.ILogger
	db     *mongo.Client
}

// New returns new *Mongo struct with passed config
func New(c *Config, l logger.ILogger) storage {
	return &Mongo{
		config: c,
		logger: l,
	}
}

// Open initialises new mongodb client
func (m *Mongo) Open() error {
	var maxRetries = m.config.MaxRetries
	var waitTime = m.config.WaitRetry
	var i int

	var err error

	for i < maxRetries {
		m.logger.Debug(fmt.Sprintf("Open mongodb connection. Attempt %s", strconv.Itoa(i+1)))
		err = m.open()
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

func (m *Mongo) open() error {
	var ctx = context.TODO()

	clientOptions := options.Client().ApplyURI(m.config.ConnectionURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	m.db = client
	return nil
}

// Close calls func disconnect for mongo
func (m *Mongo) Close() error {
	return m.db.Disconnect(context.TODO())
}

// Get retrieves records according to passed QueryGet
func (m *Mongo) Get(ctx context.Context, q QueryGet) ([]map[string]interface{}, error) {
	collection := m.db.Database(m.config.Database).Collection(q.Collection)

	result := make([]map[string]interface{}, 0)

	filter := bson.M{}

	if q.Query != nil {
		for k, v := range q.Query {
			if k == "_id" {
				oid, err := primitive.ObjectIDFromHex(v.(string))
				if err != nil {
					return nil, err
				}
				filter[k] = oid
			} else {
				filter[k] = v
			}
		}
	}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			m.logger.Error(err.Error())
		}
	}(cur, context.Background())

	for cur.Next(context.Background()) {
		doc := make(map[string]interface{})

		err := cur.Decode(&doc)
		if err != nil {
			return nil, err
		}

		result = append(result, doc)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates new records according to passed QueryInsert
func (m *Mongo) Create(ctx context.Context, q QueryInsert) (bool, error) {
	collection := m.db.Database(m.config.Database).Collection(q.Collection)
	docs := make([]interface{}, len(q.Query))
	for i, item := range q.Query {
		docs[i] = item
	}

	_, err := collection.InsertMany(ctx, docs)
	if err != nil {
		return false, err
	}

	return true, nil
}
