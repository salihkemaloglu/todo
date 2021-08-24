package pgu

import (
	"database/sql"

	"github.com/heetch/sqalx"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// OpenPostgres creates new sql.DB with pgx driver.
func OpenPostgres(url string) (*sql.DB, error) {
	connConfig, err := pgx.ParseConfig(url)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse environment for db")
	}

	return stdlib.OpenDB(*connConfig), nil
}

// MustOpen ensures pg is ready
func MustOpen(url string) (sqalx.Node, error) {
	db, err := OpenPostgres(url)
	if err != nil {
		return nil, err
	}
	node, err := sqalx.New(sqlx.NewDb(db, "pgx"))
	if err != nil {
		return nil, err
	}
	return node, nil
}
