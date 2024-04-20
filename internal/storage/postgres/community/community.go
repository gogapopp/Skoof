package community

import "github.com/jackc/pgx/v5/pgxpool"

type communityStorage struct {
	conn *pgxpool.Pool
}

func NewCommunityStorage(conn *pgxpool.Pool) *communityStorage {
	return &communityStorage{
		conn: conn,
	}
}
