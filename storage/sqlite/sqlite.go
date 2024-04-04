package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/gogapopp/Skoof/model"
	"github.com/gogapopp/Skoof/storage"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Conn *sql.DB
}

func New() (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := sql.Open("sqlite3", "storage.db")
	if err != nil {
		return &DB{}, err
	}

	err = db.Ping()
	if err != nil {
		return &DB{}, err
	}

	stmp, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS user (
		user_id SERIAL PRIMARY KEY,
		email VARCHAR(128),
		username VARCHAR(128),
		password VARCHAR(128),
		created_at TIMESTAMPTZ,
		metainfo BYTEA
	);
	`)
	if err != nil {
		return &DB{}, err
	}

	_, err = stmp.ExecContext(ctx)
	if err != nil {
		return &DB{}, err
	}

	return &DB{Conn: db}, nil
}

func (db *DB) CreateUser(ctx context.Context, user model.User) error {
	const query = "INSERT INTO user (email, username, password, created_at, metainfo) values (?1, ?2, ?3, ?4, ?5)"
	_, err := db.Conn.ExecContext(ctx, query, user.Email, user.Username, user.Password, user.CreatedAt, user.MetaInfo)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetUser(ctx context.Context, emailOrUsername, password string) (model.User, error) {
	const query = "SELECT user_id, email, username, password FROM user WHERE (email=?1 OR username=?2) AND password=?3"
	var user model.User
	rows := db.Conn.QueryRowContext(ctx, query, emailOrUsername, emailOrUsername, password)

	err := rows.Scan(&user.Email, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, storage.ErrUserNotExist
		}
	}

	return user, nil
}
