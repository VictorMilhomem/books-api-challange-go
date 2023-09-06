package models

import "github.com/google/uuid"

type Book struct {
	Id              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Edition         int32     `json:"edition"`
	PublicationYear string    `json:"publication_year"`
	Authors         *[]Author `json:"authors"`
}

type BookDB struct {
	Id              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Edition         int32     `json:"edition"`
	PublicationYear string    `json:"publication_year"`
}

func NewBook(bookDB *BookDB, authors *[]Author) Book {
	return Book{
		Id:              bookDB.Id,
		Name:            bookDB.Name,
		Edition:         bookDB.Edition,
		PublicationYear: bookDB.PublicationYear,
		Authors:         authors,
	}
}
