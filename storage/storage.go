package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres",
			"user=postgres password=E5p1n0z4% dbname=curso_go host=localhost sslmode=disable")
		if err != nil {
			log.Fatalf("can´t open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("cant´t do ping: %v", err)
		}

		fmt.Println("Connected to Postgres")
	})
}

// Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}

func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}