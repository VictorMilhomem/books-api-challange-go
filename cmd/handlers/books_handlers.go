package handlers

import (
	"database/sql"

	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/controllers"
	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/models"
	"github.com/gofiber/fiber/v2"
)

type BookHandlers struct {
	db *sql.DB
}

func NewBookHandler(db *sql.DB) *BookHandlers {
	return &BookHandlers{
		db: db,
	}
}

func (b *BookHandlers) FetchAllBooksHandler() *Handler {
	handler := NewHandler(b.db, func(c *fiber.Ctx) error {
		var authors []models.Author
		var allbooks []models.Book

		books, err := controllers.FetchAllBooks(b.db)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Could not fetch books")
		}
		for _, book := range books {

			authors, err = controllers.FetchBookAuthors(b.db, book.Name)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON("Could not fetch book authors")
			}
			temp_book := models.NewBook(&book, &authors)
			allbooks = append(allbooks, temp_book)
		}

		return c.Status(fiber.StatusOK).JSON(allbooks)
	})

	return handler
}
