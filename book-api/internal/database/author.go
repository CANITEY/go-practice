package database

import (
	"book-api/api"
	"log"

	_ "github.com/mattn/go-sqlite3"
)



// Author API

//Add author function
func AddAuthor(name string) (error) {
	statement, err := db.Prepare("INSERT INTO authors(name) VALUES(?)")
	if err != nil {
		return err
	}

	if _, err := statement.Exec(name); err != nil {
		return err
	}

	log.Printf("Author: %v, added successfully\n", name)
	return nil
}

//View author function
func ViewAuthor(id int) (api.Author, error) {
	rows, err := db.Query("SELECT authors.id, authors.name, books.title from authors INNER JOIN books ON authors.id=books.authorID where authors.id=?", id)
	if err != nil {
		return api.Author{}, err
	}

	var name string
	var books []string
	var idResult int

	for rows.Next() {
		var title string
		err := rows.Scan(&idResult, &name, &title)
		if err != nil {
			return api.Author{}, err
		}
		books = append(books, title)
	}

	result := api.Author{
		Id: idResult,
		Name: name,
		Books: books,
	}

	return result, nil
}

//Delete author function
func DeleteAuthor(id int) (error) {
	statement, err := db.Prepare("DELETE FROM authors WHERE id=?")
	if err != nil {
		return err
	}
	if _, err := statement.Exec(id); err != nil {
		return err
	}
	return nil
}

//View all authors function
func ViewAuthors() ([]api.Author, error) {
	rows, err := db.Query("SELECT authors.id, authors.name, books.title from authors LEFT JOIN books ON authors.id=books.authorID")
	if err != nil {
		return []api.Author{}, err
	}

	var results []api.Author
	var id int
	var authorName string
	var query api.Author
	query.Id = 1

	for rows.Next() {
		var title interface{}

		err := rows.Scan(&id, &authorName, &title)
		if title == nil {
			title = ""
		}
		if err != nil {
			return []api.Author{}, err
		}
		if id == query.Id {
			query.Name = authorName
			query.Books = append(query.Books, title.(string))
		} else if id != query.Id {
			results = append(results, query)
			query.Id = id
			query.Name = authorName
			query.Books = nil
			query.Books = append(query.Books, title.(string))
		}
	}
	results = append(results, query)

	return results, nil
}

