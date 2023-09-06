package controllers

import (
	"database/sql"
	"log"

	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/models"
)

func FetchBookAuthors(db *sql.DB, book_name string) ([]models.Author, error) {
	var authors []models.Author

	rows, err := db.Query(`
	SELECT *
	FROM authors
	JOIN book_authors ON authors.id = book_authors.author_id
	JOIN books ON book_authors.book_id = books.id
    WHERE books.name LIKE '$1';`, book_name)
	if err != nil {
		log.Println("Error querying data ", err)
		return nil, err
	}

	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.Id, &author.Name); err != nil {
			log.Println("Error scanning model ", err)
			return nil, err
		}
		log.Println(author)
		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error fetching authors", err)
		return nil, err
	}

	return authors, nil
}

func FetchAllBooks(db *sql.DB) ([]models.BookDB, error) {
	var books []models.BookDB

	rows, err := db.Query("SELECT * FROM books;")
	if err != nil {
		log.Println("Error fetching the books")
		return nil, err
	}

	for rows.Next() {
		for rows.Next() {
			var book models.BookDB
			if err := rows.Scan(&book.Id, &book.Name, &book.Edition, &book.PublicationYear); err != nil {
				log.Println("Error scanning model ", err)
				return nil, err
			}
			log.Println(book)
			books = append(books, book)
		}

		if err := rows.Err(); err != nil {
			log.Println("Error fetching books", err)
			return nil, err
		}
	}
	return books, nil
}
