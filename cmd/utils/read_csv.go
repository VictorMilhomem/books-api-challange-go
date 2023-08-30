package utils

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/VictorMilhomem/go-bakcend-challenge/cmd/models"
	"github.com/google/uuid"
)

func GenerateSQLQueryFromCSV(filepath string, db *sql.DB) {
	csvFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	// create a sql file for inserting each author
	sqlFileName := "./cmd/migrations/" + uuid.New().String() + "auhors_migration.sql"
	sqlFile, err := os.Create(sqlFileName)
	if err != nil {
		log.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer sqlFile.Close()
	err = readFileAndCreateSQL(reader, sqlFile)

	if err != nil {
		log.Fatal("Failed to create sql file")
	}

	err = executeSQLFromFile(db, sqlFileName)

	if err != nil {
		log.Fatal("Migration Failed")
	}
}

func readFileAndCreateSQL(csvReader *csv.Reader, sqlFile *os.File) error {
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		author := models.NewAuthorName(record[0])

		data := fmt.Sprintf("INSERT INTO authors (name) VALUES (%s);\n", author.Name)

		_, err = sqlFile.WriteString(data)
		if err != nil {
			fmt.Println("Erro ao escrever no arquivo:", err)
			return err
		}
	}
	return nil
}

func executeSQLFromFile(db *sql.DB, filePath string) error {
	sqlContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlContent))
	if err != nil {
		return err
	}

	return nil
}
