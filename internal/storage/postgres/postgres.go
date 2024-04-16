package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	conn *pgxpool.Pool
}

func New(dsn string) (*repository, error) {
	const op = "postgres.postgres.New"
	ctx := context.Background()
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	err = db.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	tx, err := db.Begin(ctx)
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
	return &repository{
		conn: db,
	}, tx.Commit(ctx)
}

func (s *repository) Close() {
	s.conn.Close()
}
