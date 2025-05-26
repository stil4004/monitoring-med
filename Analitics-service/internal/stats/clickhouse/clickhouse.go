package clickhouse

import (
	"context"
	"service/config"
	"service/internal/stats"

	//"database/sql/driver"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

const (
	insertQuery = `INSERT INTO %s (created_at, params, headers, method_name, success, error_description, hash) VALUES (?, ?, ?, ?, ?, ?, ?)`
)

type Writer struct {
	conn      driver.Conn
	tableName string
}

func NewClickhouseWriter(cfg *config.Config) (*Writer, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", cfg.ClickHouse.Host, cfg.ClickHouse.Port)},
		Auth: clickhouse.Auth{
			Database: cfg.ClickHouse.Database,
			Username: cfg.ClickHouse.Username,
			Password: cfg.ClickHouse.Password,
		},
		Debug:           cfg.ClickHouse.Debug,
		DialTimeout:     cfg.ClickHouse.DialTimeout,
		MaxOpenConns:    cfg.ClickHouse.MaxOpenConns,
		MaxIdleConns:    cfg.ClickHouse.MaxIdleConns,
		ConnMaxLifetime: cfg.ClickHouse.ConnMaxLifetime,
	})

	if err != nil {
		return nil, err
	}

	return &Writer{
		conn:      conn,
		tableName: cfg.ClickHouse.Table,
	}, nil
}

func (w *Writer) Insert(rows stats.Rows) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	batch, err := w.conn.PrepareBatch(ctx, fmt.Sprintf(insertQuery, w.tableName))
	if err != nil {
		return err
	}

	for _, v := range rows {
		err := batch.Append(
			v.Timestamp,
			v.Params,
			v.Headers,
			v.Method,
			v.Success,
			v.ErrorDescription,
			v.Hash,
		)

		if err != nil {
			return err
		}
	}

	return batch.Send()
}
