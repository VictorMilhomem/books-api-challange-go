package handlers

import (
	"database/sql"

	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/controllers"
	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func FetchAllAuthorsHandler(db *sql.DB) *Handler {
	handler := NewHandler(db, func(c *fiber.Ctx) error {
		authors, err := controllers.FetchAllAuthors(db)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Could not fetch authors")
		}

		return c.Status(fiber.StatusOK).JSON(authors)
	})

	return handler
}

func FetchAuthorByIdHandler(db *sql.DB) *Handler {
	handler := NewHandler(db, func(c *fiber.Ctx) error {
		author, err := controllers.FetchAuthorById(db, uuid.MustParse(c.Params("id")))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Could not fetch author")
		}

		return c.Status(fiber.StatusOK).JSON(author)
	})

	return handler
}

func CreateAuthorHandler(db *sql.DB) *Handler {
	handler := NewHandler(db, func(c *fiber.Ctx) error {
		var author models.AuthorName
		err := c.BodyParser(&author)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		err = controllers.CreateAuthor(db, &author)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
		return c.Status(fiber.StatusCreated).JSON(author)
	})

	return handler
}
