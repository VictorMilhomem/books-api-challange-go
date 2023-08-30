package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func env(key string) string {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Panic("Failed to load env file")
	}

	return os.Getenv(key)
}

func ConnectDB() (*sql.DB, error) {
	db_string_config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env("HOST"), env("USER"), env("PASSWORD"), env("DB"), env("PORT"))

	db, err := sql.Open("postgres", db_string_config)

	if err != nil {
		log.Panic(err.Error())
	} else {
		log.Println("Connected!")
	}

	return db, err
}
