package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/gogapopp/Skoof/internal/model"
	"github.com/gogapopp/Skoof/internal/storage"
	_ "github.com/mattn/go-sqlite3"
)

type repository struct {
	conn *sql.DB
}

func New() (*repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := sql.Open("sqlite3", "storage.db")
	if err != nil {
		return &repository{}, err
	}

	err = db.Ping()
	if err != nil {
		return &repository{}, err
	}

	stmp, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS user (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		email VARCHAR(128),
		username VARCHAR(128),
		password VARCHAR(128),
		created_at TIMESTAMPTZ,
		metainfo BYTEA
	);
	`)
	if err != nil {
		return &repository{}, err
	}

	_, err = stmp.ExecContext(ctx)
	if err != nil {
		return &repository{}, err
	}

	return &repository{conn: db}, nil
}

func (db *repository) CreateUser(ctx context.Context, user model.User) error {
	const query = "INSERT INTO user (email, username, password, created_at, metainfo) values (?1, ?2, ?3, ?4, ?5)"
	_, err := db.conn.ExecContext(ctx, query, user.Email, user.Username, user.Password, user.CreatedAt, user.MetaInfo)
	if err != nil {
		return err
	}
	return nil
}

func (db *repository) GetUser(ctx context.Context, emailOrUsername, password string) (model.User, error) {
	const query = "SELECT user_id, email, username, password FROM user WHERE (email=?1 OR username=?2) AND password=?3"
	var user model.User
	rows := db.conn.QueryRowContext(ctx, query, emailOrUsername, emailOrUsername, password)

	err := rows.Scan(&user.UserID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, storage.ErrUserNotExist
		}
	}

	return user, nil
}

func (db *repository) Close() error {
	return db.conn.Close()
}
