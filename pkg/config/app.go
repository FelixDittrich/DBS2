package config

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func Connect() {
	d, err := sql.Open("sqlite3", "../pkg/config/database.db")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *sql.DB {
	return db
}
