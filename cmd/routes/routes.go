package routes

import (
	"database/sql"

	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/handlers"
	"github.com/gofiber/fiber/v2"
)

func Routes(db *sql.DB, app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!!")
	})

	app.Get("/authors", handlers.FetchAllAuthorsHandler(db).Func)
	app.Get("/authors/:id", handlers.FetchAuthorByIdHandler(db).Func)
}
