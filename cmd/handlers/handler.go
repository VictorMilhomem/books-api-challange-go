package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type cbck func(c *fiber.Ctx) error

type Handler struct {
	DB   *sql.DB
	Func cbck
}

func NewHandler(db *sql.DB, fn cbck) *Handler {
	return &Handler{
		DB:   db,
		Func: fn,
	}
}
