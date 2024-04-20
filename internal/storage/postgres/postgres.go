package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(dsn string) (*pgxpool.Pool, error) {
	const op = "postgres.postgres.New"
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer func() {
		err = tx.Rollback(ctx)
	}()
	_, err = tx.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS users (
		user_id SERIAL PRIMARY KEY,
		email TEXT,
		username TEXT,
		password_hash TEXT,
		created_at TIMESTAMP,
		role TEXT,
		metainfo TEXT
	);
	CREATE UNIQUE INDEX IF NOT EXISTS username_idx ON users(username);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return conn, tx.Commit(ctx)
}
