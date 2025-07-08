package e2e

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

func (m dbsManager) clearUsersDB(ctx context.Context) error {
	_, err := m.usersDB.ExecContext(ctx, `TRUNCATE TABLE profiles`)
	if err != nil {
		return fmt.Errorf("could not truncate table profiles: %w", err)
	}

	_, err = m.usersDB.ExecContext(ctx, `TRUNCATE TABLE users`)
	if err != nil {
		return fmt.Errorf("could not truncate table users: %w", err)
	}

	return nil
}

func (m dbsManager) clearPostsDB(ctx context.Context) error {
	_, err := m.postsDB.ExecContext(ctx, `TRUNCATE TABLE posts`)
	if err != nil {
		return fmt.Errorf("could not create truncate posts: %w", err)
	}

	_, err = m.postsDB.ExecContext(ctx, `TRUNCATE TABLE posts_info`)
	if err != nil {
		return fmt.Errorf("could not create truncate posts_info: %w", err)
	}

	return nil
}

func (m dbsManager) clearStatsDB(ctx context.Context) error {
	err := m.statsDB.Exec(ctx, `TRUNCATE TABLE daily_stats`)
	if err != nil {
		return fmt.Errorf("could not truncate table daily_stats: %w", err)
	}
	return nil
}

func (m dbsManager) clearDB(ctx context.Context) error {
	if err := m.clearUsersDB(ctx); err != nil {
		return err
	}
	if err := m.clearPostsDB(ctx); err != nil {
		return err
	}
	if err := m.clearStatsDB(ctx); err != nil {
		return err
	}
	return nil
}
