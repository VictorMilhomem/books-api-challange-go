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

func GenerateSQLQueryFromCSV(filepath string, db *sql.DB) error {
	csvFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	// create a sql file for inserting each author
	sqlFileName := "cmd/migrations/" + uuid.New().String() + "authors_migration.sql"
	sqlFile, err := readFileAndCreateSQL(reader, sqlFileName)
	if err != nil {
		log.Fatal("Failed to create sql file ", sqlFile)
	}

	err = executeSQLFromFile(db, sqlFileName)

	if err != nil {
		log.Fatal("Migration Failed")
	}
	return nil
}

func readFileAndCreateSQL(csvReader *csv.Reader, sqlFileName string) (*os.File, error) {
	sqlFile, err := os.Create(sqlFileName)
	if err != nil {
		log.Println("Erro ao criar o arquivo:", err)
		return nil, err
	}
	defer sqlFile.Close()

	line := 0
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if line > 0 {
			author := models.NewAuthorName(record[0])

			data := fmt.Sprintf("INSERT INTO authors (name) VALUES ('%s');\n", author.Name)

			_, err = sqlFile.WriteString(data)
			if err != nil {
				fmt.Println("Erro ao escrever no arquivo:", err)
				return nil, err
			}
		}
		line++
	}
	return sqlFile, nil
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
