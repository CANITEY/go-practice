package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db, err = sql.Open("sqlite3", "books.db")

func init() {
	if err := createTable(); err != nil {
		log.Panic(err)
	}
}
func createTable() (error) {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS authors(id INTEGER, name varchar, PRIMARY KEY(id AUTOINCREMENT))")
	if err != nil {
		return err
	}

	if _, err := statement.Exec(); err != nil {
		return err
	}
	log.Println("authors table created successfully")
	statement, err = db.Prepare("CREATE TABLE IF NOT EXISTS books(id INTEGER, title varchar, description varchar, authorID int, FOREIGN KEY(authorID) REFERENCES authors(id), PRIMARY KEY(id AUTOINCREMENT))")
	if err != nil {
		return err
	}

	if _, err := statement.Exec(); err != nil {
		return err
	}
	log.Println("books table created successfully")
	return nil
}



