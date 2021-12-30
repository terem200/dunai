package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.insigit.com/qa/outrunner/pkg/logger"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	config *Config
	logger logger.ILogger
	db     *sqlx.DB
}

// New - initialize new MySQL struct with config
func New(c *Config, l logger.ILogger) storage {
	return &mysql{
		config: c,
		logger: l,
	}
}

// Open new MySQL connection using passed to New func Config
func (m *mysql) Open() error {
	var maxRetries = m.config.MaxRetries
	var waitTime = m.config.WaitRetry
	var i int

	var err error

	for i < maxRetries {
		m.logger.Debug(fmt.Sprintf("Open mysql connection. Attempt %s", strconv.Itoa(i+1)))
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

func (m *mysql) open() error {
	db, err := sqlx.Open("mysql", m.config.ConnectionURL)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	db.DB.SetMaxOpenConns(40)
	db.DB.SetMaxIdleConns(5)

	m.db = db

	return nil
}

// Close current MySQL connection
func (m *mysql) Close() error {
	return m.db.Close()
}

// Get returns records from table passed in url
// and by query passed in body
func (m *mysql) Get(ctx context.Context, stmt string) ([]map[string]interface{}, error) {
	rows, err := m.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	// It can be slow but necessary when we work with interface type
	// If you want to get more info - look at
	//         https://github.com/go-sql-driver/mysql/pull/1281
	//         https://github.com/golang/go/issues/22544
	// If you have better idea how to deal with it, let me know
	res, err := rowsToJSON(rows)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create records by passed query
func (m *mysql) Create(ctx context.Context, stmt string) (ok bool, err error) {
	if _, err := m.db.ExecContext(ctx, stmt); err != nil {
		return false, err
	}

	return true, nil
}
