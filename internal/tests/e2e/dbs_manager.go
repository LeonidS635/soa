package e2e

import (
	"database/sql"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

const (
	usersDBConn = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	postsDBConn = "postgres://postgres:postgres@localhost:6543/postgres?sslmode=disable"
)

type dbsManager struct {
	usersDB *sql.DB
	postsDB *sql.DB
	statsDB driver.Conn
}

func initDBsManager() (dbsManager, error) {
	m := dbsManager{}
	var err error

	m.usersDB, err = sql.Open("postgres", usersDBConn)
	if err != nil {
		return m, err
	}
	m.postsDB, err = sql.Open("postgres", postsDBConn)
	if err != nil {
		return m, err
	}
	m.statsDB, err = clickhouse.Open(
		&clickhouse.Options{
			Addr: []string{"localhost:9000"},
			Auth: clickhouse.Auth{
				Username: "clickhouse",
				Password: "clickhouse",
				Database: "clickhouse",
			},
			Protocol: clickhouse.Native,
		},
	)

	return m, err
}

func (m dbsManager) close() error {
	var err error
	if userErr := m.usersDB.Close(); userErr != nil {
		err = userErr
	}
	if postsErr := m.postsDB.Close(); postsErr != nil {
		err = postsErr
	}
	if statsErr := m.statsDB.Close(); statsErr != nil {
		err = statsErr
	}
	return err
}
