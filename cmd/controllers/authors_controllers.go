package controllers

import (
	"database/sql"
	"log"

	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/models"
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
