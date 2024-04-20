package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/gogapopp/Skoof/internal/model"
	"github.com/gogapopp/Skoof/internal/storage"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type authStorage struct {
	conn *pgxpool.Pool
}

func NewAuthStorage(conn *pgxpool.Pool) *authStorage {
	return &authStorage{
		conn: conn,
	}
}

func (s *authStorage) SignUp(ctx context.Context, user model.SignUpUser) error {
	const (
		op    = "postgres.auth.SignUp"
		query = "INSERT INTO users (email, username, password_hash, created_at, role) VALUES ($1, $2, $3, $4, $5);" // TODO: add metainfo
	)
	_, err := s.conn.Exec(ctx, query, user.Email, user.Username, user.PasswordHash, user.CreatedAt, user.Role)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return fmt.Errorf("%s: %w", op, storage.ErrUserExists)
			}
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *authStorage) SignIn(ctx context.Context, user model.SignInUser) (int, string, error) {
	const (
		op    = "postgres.auth.SignIn"
		query = "SELECT user_id, role FROM users WHERE username=$1 AND password_hash=$2;"
	)
	var (
		userID int
		role   string
	)
	row := s.conn.QueryRow(ctx, query, user.Username, user.PasswordHash)
	err := row.Scan(&userID, &role)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, "", fmt.Errorf("%s: %w", op, storage.ErrUserNotExist)
		}
		return 0, "", fmt.Errorf("%s: %w", op, err)
	}
	return userID, role, nil
}
