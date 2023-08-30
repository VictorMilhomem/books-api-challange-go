package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB  *sql.DB
	err error
)

func env(key string) string {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Panic("Failed to load env file")
	}

	return os.Getenv(key)
}

func ConnectDB() {
	db_string_config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env("HOST"), env("USER"), env("PASSWORD"), env("DB"), env("PORT"))

	DB, err = sql.Open("postgres", db_string_config)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}
}
