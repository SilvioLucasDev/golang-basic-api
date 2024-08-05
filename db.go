package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}

	if db == nil {
		log.Fatal("db nil")
	}

	createTable(db)
	return db
}

func createTable(db *sql.DB) {
	sqlTable := `
		CREATE TABLE IF NOT EXISTS tasks(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL
		);
	`

	_, err := db.Exec(sqlTable)

	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}
}
