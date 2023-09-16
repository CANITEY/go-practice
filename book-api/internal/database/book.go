package database

import (
	"book-api/api"

	_ "github.com/mattn/go-sqlite3"
)


// books api

//Add book
func AddBook(title, description string, authorID int) (error) {
	statement, err := db.Prepare("INSERT INTO books(title, description, authorID) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	if _, err := statement.Exec(title, description, authorID); err != nil {
		return err
	}
	return nil
}

//View book
func ViewBook(id int) (api.Book, error) {
	row := db.QueryRow("SELECT books.id, books.title, books.description, authors.name FROM books LEFT JOIN authors on books.authorID=authors.id where books.id=?",id)
	var result api.Book
	if err := row.Scan(&result.Id, &result.Title, &result.Description, &result.Author); err != nil {
		return api.Book{}, err
	}
	return result, nil
}

//View books
func ViewBooks() ([]api.Book, error) {
	var results []api.Book
	rows, err := db.Query("SELECT books.id, books.title, books.description, authors.name FROM books LEFT JOIN authors on books.authorID=authors.id")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var query api.Book
		rows.Scan(&query.Id, &query.Title, &query.Description, &query.Author)
		results = append(results, query)
	}
	
	return results, nil
}

//Delete book
func DeleteBook(id int) (error) {
	statement, err := db.Prepare("DELETE FROM books WHERE id=?")
	if err != nil {
		return err
	}
	
	if _, err := statement.Exec(id); err != nil {
		return err
	}
	return nil
}


