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

	authorHandler := handlers.NewAuthorHandler(db)

	app.Get("/authors", authorHandler.FetchAllAuthorsHandler().Func)
	app.Get("/authors/:id", authorHandler.FetchAuthorByIdHandler().Func)
	app.Post("/authors", authorHandler.CreateAuthorHandler().Func)
	app.Delete("/authors/:id", authorHandler.DeleteAuthorByIdHandler().Func)

	bookHandler := handlers.NewBookHandler(db)
	app.Get("/books", bookHandler.FetchAllBooksHandler().Func)
}
