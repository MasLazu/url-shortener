package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase() *Database {
	db, err := sql.Open("postgres", os.Getenv("POSTGRE_URI"))
	if err != nil {
		panic(err)
	}

	return &Database{
		Conn: db,
	}
}

func (db *Database) Close() {
	if err := db.Conn.Close(); err != nil {
		panic(err)
	}
}
