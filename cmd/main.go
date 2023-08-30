package main

import (
	"log"

	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/database"
	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/utils"
	"github.com/gofiber/fiber/v2"
)

/*
	Create a command line tool to convert a csv into sql insert query
	for each author and migrate to database
*/

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Panic("Failed to connect to database")
	}

	app := fiber.New()

	// all routes here
	listenPort := ":" + utils.Env("PORT")
	app.Listen(listenPort)
}
