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

func (db *Db) AddFile(m Media) (sql.Result, error) {
	var error_list = []string{}
	if m.Name == "" {
		error_list = append(error_list, "Filename can't be empty")
	}
	if m.Location == "" {
		error_list = append(error_list, "File location can't be empty")
	}
	if m.Hash == "" {
		error_list = append(error_list, "File Hash can't be empty")
	}

	if len(error_list) > 0 {
		errors := strings.Join(error_list[:], "\\n")
		return nil, fmt.Errorf(errors)
	}
	isNewFile, _ := db.isNewFile(m)
	if isNewFile {
		return db.connection.Exec("INSERT INTO media (name, location, type, hash) VALUES(?,?,?,?)", m.Name, m.Location, m.Type, m.Hash)
	} else {
		return db.UpdateFile(m)
	}

}

func (db *Db) GetFile(m Media) (Media, error) {
	var fileResult = Media{}
	err := db.connection.QueryRow("SELECT * FROM media WHERE hash=?", m.Hash).Scan(&fileResult.Id, &fileResult.Name, &fileResult.Location, &fileResult.Type, &fileResult.Hash)
	if err != nil {
		return fileResult, err
	}
	return fileResult, err
}

func (db *Db) UpdateFile(m Media) (sql.Result, error) {
	return db.connection.Exec("UPDATE media SET name = ?, location = ?, hash = ? WHERE id = ?", m.Name, m.Location, m.Hash, m.Id)
}

func (db *Db) isNewFile(m Media) (bool, error) {
	media, err := db.GetFile(m)
	if err != sql.ErrNoRows {
		return false, err
	} else if media.Id != 0 {
		return false, err
	}

	return true, err
}
