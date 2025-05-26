package postgres

import (
	"context"
	"fmt"
	"service/config"
	"service/internal/stats"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	insertQuery = `INSERT INTO %s (created_at, params, headers, method_name, success, error_description, hash, resp_body_on_error) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
)

type Writer struct {
	db    *pgxpool.Pool
	query string
}

func NewPostgresWriter(cfg *config.Config, db *pgxpool.Pool) (*Writer, error) {
	return &Writer{
		db:    db,
		query: fmt.Sprintf(insertQuery, cfg.Postgres.Table),
	}, nil
}

func (w *Writer) Insert(rows stats.Rows) error {
	tx, err := w.db.BeginTx(context.TODO(), pgx.TxOptions{})
	if err != nil {
		return err
	}

	for _, v := range rows {
		_, err := tx.Exec(
			context.TODO(),
			w.query,
			v.Timestamp,
			v.Params,
			v.Headers,
			v.Method,
			v.Success,
			v.ErrorDescription,
			v.Hash,
			v.BodyOnError,
		)

		if err != nil {
			return err
		}
	}

	return tx.Commit(context.TODO())
}
