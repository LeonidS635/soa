package impl

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgPostsStorageImpl struct {
	db *pgxpool.Pool
}

func NewPgPostsStorageImpl(pool *pgxpool.Pool) (*PgPostsStorageImpl, error) {
	return &PgPostsStorageImpl{pool}, nil
}

func (pg *PgPostsStorageImpl) Close() {
	pg.db.Close()
}
