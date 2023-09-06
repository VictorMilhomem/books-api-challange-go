package handlers

import (
	"database/sql"

	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/controllers"
	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthorHandler struct {
	db *sql.DB
}

func NewAuthorHandler(db *sql.DB) *AuthorHandler {
	return &AuthorHandler{
		db: db,
	}
}

func (a AuthorHandler) FetchAllAuthorsHandler() *Handler {
	handler := NewHandler(a.db, func(c *fiber.Ctx) error {
		authors, err := controllers.FetchAllAuthors(a.db)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Could not fetch authors")
		}

		return c.Status(fiber.StatusOK).JSON(authors)
	})

	return handler
}

func (a AuthorHandler) FetchAuthorByIdHandler() *Handler {
	handler := NewHandler(a.db, func(c *fiber.Ctx) error {
		author, err := controllers.FetchAuthorById(a.db, uuid.MustParse(c.Params("id")))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Could not fetch author")
		}

		return c.Status(fiber.StatusOK).JSON(author)
	})

	return handler
}

func (a AuthorHandler) CreateAuthorHandler() *Handler {
	handler := NewHandler(a.db, func(c *fiber.Ctx) error {
		var author models.AuthorName
		err := c.BodyParser(&author)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		err = controllers.CreateAuthor(a.db, &author)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
		return c.Status(fiber.StatusCreated).JSON(author)
	})

	return handler
}

func (a AuthorHandler) DeleteAuthorByIdHandler() *Handler {
	handler := NewHandler(a.db, func(c *fiber.Ctx) error {
		id := uuid.MustParse(c.Params("id"))

		err := controllers.DeleteAuthorById(a.db, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON("Author deleted")
	})

	return handler
}
