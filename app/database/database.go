package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func (db *Database) Connect(connectionStr string) error {
	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return err
	}
	db.Conn = conn
	return nil
}
