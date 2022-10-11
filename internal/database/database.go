package database

import (
	"database/sql"
	"log"
)

// Initialize our database and create our tables.
func InitDatabase(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./"+name+".db?mode=memory")
	if err != nil {
		log.Fatal(err)
	}

	// Drop and create the tables we need.
	db.Exec(`
		DROP TABLE IF EXISTS books;
		CREATE TABLE books (id INTEGER PRIMARY KEY, title TEXT, author TEXT);
	`)

	return db, err
}
