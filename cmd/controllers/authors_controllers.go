package controllers

import (
	"database/sql"
	"log"

	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/models"
	"github.com/google/uuid"
)

func FetchAllAuthors(db *sql.DB) ([]models.Author, error) {
	var authors []models.Author

	rows, err := db.Query("SELECT id, name FROM authors;")
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

func FetchAuthorById(db *sql.DB, id uuid.UUID) (models.Author, error) {
	var author models.Author

	err := db.QueryRow("SELECT * FROM authors WHERE id = $1", id).Scan(&author.Id, &author.Name)
	if err != nil {
		log.Println("Error querying author")
		return author, err
	}

	return author, nil
}
