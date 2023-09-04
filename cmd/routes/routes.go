package routes

import (
	"database/sql"

	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(db *sql.DB, app *fiber.App) {
	app.Get("/authors", func(c *fiber.Ctx) error {
		authors, err := controllers.FetchAllAuthors(db)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Could not fetch authors")
		}

		return c.Status(fiber.StatusOK).JSON(authors)
	})
}
