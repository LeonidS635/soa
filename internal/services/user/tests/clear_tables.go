package tests

import "context"

func clearTables(ctx context.Context) error {
	if _, err := db.Exec(ctx, `TRUNCATE TABLE users`); err != nil {
		return err
	}
	if _, err := db.Exec(ctx, `TRUNCATE TABLE profiles`); err != nil {
		return err
	}
	return nil
}
