package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	pkgErr "github.com/pkg/errors"
)

// Connect: uses to connect to the database
func Connect(dns string) (*sql.DB, error) {
	conn, err := openDB(dns)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Postgres!")
	return conn, nil
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, pkgErr.WithStack(err)
	}

	if err = db.Ping(); err != nil {
		return nil, pkgErr.WithStack(err)
	}

	return db, nil
}
