package db

import (
	"database/sql"
	_ "embed"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB   *sql.DB
	once sync.Once

	//go:embed migrations/initial_schema_01.sql
	schemaSQL string
)

func InitDB(dbpath string) error {
	var err error
	DB, err = sql.Open("sqlite3", dbpath)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	if _, err := DB.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return err
	}

	if _, err = DB.Exec(schemaSQL); err != nil {
		return err
	}

	return nil
}

func Close() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
