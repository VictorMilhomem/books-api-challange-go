package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/utils"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db_string_config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		utils.Env("HOST"), utils.Env("USER"), utils.Env("PASSWORD"), utils.Env("DB"), utils.Env("PORT"))

	db, err := sql.Open("postgres", db_string_config)
	if err != nil {
		log.Panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Connected!")

	return db, nil
}
