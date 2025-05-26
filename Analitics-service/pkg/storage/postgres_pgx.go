package storage

import (
	"context"
	"fmt"
	"log"
	"service/config"

	"github.com/pkg/errors"

	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       string
	SslMode  string
}

func NewDB(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DBName, cfg.Postgres.SSLMode)
	log.Println(connectionUrl)
	return pgxpool.New(ctx, connectionUrl)
}

// use to get slice of structs with needed
func SelectContext[T any](ctx context.Context, pool *pgxpool.Pool, dest *[]T, query string, args ...any) error {
	var rows pgx.Rows

	if dest == nil {
		return errors.New("destination must be a pointer to a slice, not nil")
	}
	rows, err := pool.Query(ctx, query, args...)
	defer rows.Close()
	if err != nil {
		return err
	}

	collectRows, err := pgx.CollectRows(rows, pgx.RowToStructByName[T])
	if err != nil {
		return err
	}
	*dest = collectRows

	return nil
}

// use to get single struct
func GetContext[T any](ctx context.Context, pool *pgxpool.Pool, dest *T, query string, args ...any) error {
	if dest == nil {
		return errors.New("destination must be a pointer, not nil")
	}
	row, err := pool.Query(ctx, query, args...)
	defer row.Close()
	if err != nil {
		return err
	}

	for row.Next() {
		t, err := pgx.RowToStructByName[T](row)
		if err != nil {
			return err
		}
		*dest = t
		return nil
	}
	return nil
}

// use to do some action with db, like insert, update, delete, etc. without returning any data
// when you need RETURNING id or smth else
func ExecRetFieldContext[T any](ctx context.Context, pool *pgxpool.Pool, dest *T, query string, args ...any) error {
	if dest == nil {
		return errors.New("destination must be a pointer, not nil")
	}
	err := pool.QueryRow(ctx, query, args...).Scan(dest)
	if err != nil {
		return err
	}

	return nil
}

// use to do some action with db, like insert, update, delete, etc. without returning any data
func ExecContext(ctx context.Context, pool *pgxpool.Pool, query string, args ...any) error {
	result, err := pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
