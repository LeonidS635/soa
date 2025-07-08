package service

import "context"

func clearTables(ctx context.Context) error {
	return db.Exec(ctx, `TRUNCATE TABLE daily_stats`)
}
