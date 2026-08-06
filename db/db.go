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
	// TODO: maybe remove the once.Do() since this should be called only once always
	once.Do(func() {
		DB, err = sql.Open("sqlite3", dbpath)
		if err != nil {
			return
		}

		if err = DB.Ping(); err != nil {
			return
		}

		if _, err = DB.Exec(schemaSQL); err != nil {
			return
		}
	})

	return err
}

func close() {
	if err := DB.Close(); err != nil {
		log.Fatal(err)
	}
}
