package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type Db struct {
	connection *sql.DB
}

func (db *Db) ConnectDB() error {
	var err error
	db.connection, err = sql.Open("sqlite3", "gomediamanager.db")
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) AddFile(m media) (sql.Result, error) {
	var error_list = []string{}
	if m.name == "" {
		error_list = append(error_list, "Filename can't be empty")
	}
	if m.location == "" {
		error_list = append(error_list, "File location can't be empty")
	}
	if m.hash == "" {
		error_list = append(error_list, "File Hash can't be empty")
	}

	if len(error_list) > 0 {
		errors := strings.Join(error_list[:], "\\n")
		return nil, fmt.Errorf(errors)
	}

	return db.connection.Exec("INSERT INTO media (name, location, type, hash) VALUES(?,?,?,?)", m.name, m.location, m.Type, m.hash)
}
