package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env(key string) string {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Panic("Failed to load env file")
	}

	return os.Getenv(key)
}
