package e2e

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

func (m dbsManager) fillUsersDB(ctx context.Context) error {
	schemaProfiles := `
	CREATE TABLE profiles
	(
		user_id     int  not null,
		name        text default '',
		surname     text default '',
		age         int default 0,
		email       text not null,
		phone       text default '',
		city        text default '',
		description text default '',
		birthdate   timestamptz default now(),
		updated_at  timestamptz default now(),
	
		primary key (user_id)
	)`

	schemaUsers := `
	CREATE TABLE users
	(
		user_id       serial not null,
		username      text   not null,
		password_hash bytea  not null,
		email         text   not null,
		created_at    timestamptz default now(),
	
		primary key (user_id)
	)`

	_, err := m.usersDB.ExecContext(ctx, schemaProfiles)
	if err != nil {
		return fmt.Errorf("could not create table profiles: %w", err)
	}

	_, err = m.usersDB.ExecContext(ctx, schemaUsers)
	if err != nil {
		return fmt.Errorf("could not create table users: %w", err)
	}

	return nil
}

func (m dbsManager) fillPostsDB(ctx context.Context) error {
	schemaPosts := `
	CREATE TABLE IF NOT EXISTS posts (
	    post_id    serial not null,

    	is_private boolean default false,
    	title      text    default '',
    	tags       text[]  default array []::text[],
    	text       text    default '',

    	primary key (post_id)
	)`

	schemaPostsInfo := `
	CREATE TABLE IF NOT EXISTS posts_info (
	    post_id    int not null,
    	author_id  int not null,

    	created_at timestamptz default now(),
    	updated_at timestamptz default now(),

    	primary key (post_id)
	)`

	_, err := m.postsDB.ExecContext(ctx, schemaPosts)
	if err != nil {
		return fmt.Errorf("could not create table posts: %w", err)
	}

	_, err = m.postsDB.ExecContext(ctx, schemaPostsInfo)
	if err != nil {
		return fmt.Errorf("could not create table posts_info: %w", err)
	}

	return nil
}

func (m dbsManager) fillStatsDB(ctx context.Context) error {
	schemaStats := `
	CREATE TABLE daily_stats
	(
		post_id   Int32,
		author_id Int32,
		post_date DATETIME,
		views     UInt8 DEFAULT 0,
		likes     UInt8 DEFAULT 0,
		comments  UInt8 DEFAULT 0
	)
	ENGINE = MergeTree
	PRIMARY KEY post_id`

	err := m.statsDB.Exec(ctx, schemaStats)
	if err != nil {
		return fmt.Errorf("could not create table daily_stats: %w", err)
	}
	return nil
}

func (m dbsManager) fillDB(ctx context.Context) error {
	if err := m.fillUsersDB(ctx); err != nil {
		return err
	}
	if err := m.fillPostsDB(ctx); err != nil {
		return err
	}
	if err := m.fillStatsDB(ctx); err != nil {
		return err
	}
	return nil
}
