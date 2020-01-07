package database

import (
	"bufio"
	"log"
	"os"
)

// Database is a wrapper structure for data and behaviour specified
type Database struct {
	data            []DatabaseRecord
	queryService    QueryService
	responseService ResponseService
}

// Query is the main functionality of the Database structure as it uses all its\
// components to obtain the response of desired format
func (db Database) Query(col, min, max int) int {
	return db.responseService(db.queryService(db.data, col, min, max))
}

// LoadDatabase is a commodity constructor for the Database structure
func LoadDatabase(dataFileName string, queryService QueryService, responseService ResponseService) Database {
	file, err := os.Open(dataFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := make([]DatabaseRecord, 0)
	for scanner.Scan() {
		data = append(data, LoadDatabaseRecord(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return Database{data, queryService, responseService}
}
