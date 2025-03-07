package impl

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgUserStorageImpl struct {
	db *pgxpool.Pool
}

func NewPgUserStorageImpl(ctx context.Context, pool *pgxpool.Pool) (*PgUserStorageImpl, error) {
	return &PgUserStorageImpl{pool}, nil
}

func (pg *PgUserStorageImpl) Close() {
	pg.db.Close()
}
