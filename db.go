package main

import (
	"database/sql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "gomediamanager.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
