package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/models"
	"github.com/google/uuid"
)

func GenerateSQLQueryFromCSV(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create a sql file for inserting each author
	fileName := uuid.New().String() + "auhors_migration.sql"
	file, err = os.Create(fileName)
	if err != nil {
		log.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		author := models.NewAuthorName(record[0])

		data := fmt.Sprintf("INSERT INTO authors (name) VALUES (%s);", author.Name)

		_, err = file.WriteString(data)
		if err != nil {
			fmt.Println("Erro ao escrever no arquivo:", err)
			return
		}

	}
}
