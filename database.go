package main

import (
	"strconv"
	"strings"	
	"bufio"
    "fmt"
    "log"
    "os"
)

const (
	// ID column of DatabaseRecord
	ID = iota  // 0

	// AGE column of DatabaseRecord
	AGE = iota  // 1

	// DGX column of DatabaseRecord
	DGX = iota  // 2
)

// DatabaseRecord is a wrapper structure for data records from the database;
// can be replaced with a simple byte slice
type DatabaseRecord struct {
	ID int
	Name string
	Age int
	Dgx int
}

// Get the requested field of the structure
func (record DatabaseRecord) Get(col int) int {
	switch col {
	case ID:
		return record.ID
	case AGE:
		return record.Age
	case DGX:
		return record.Dgx
	default:
		return -1
	}
}

// ToString serializes the object to a human readable format
func (record DatabaseRecord) ToString() string {
	return fmt.Sprintf("[%d-%s-%d-%d]", record.ID, record.Name, record.Age, record.Dgx)
}

// LoadDatabaseRecord is a commodity constructor for the DatabaseRecord structure
func LoadDatabaseRecord(line string) DatabaseRecord {
	values := strings.Split(line, ",")
	id, err := strconv.ParseInt(values[0], 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	name := values[1]

	age, err := strconv.ParseInt(values[2], 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	dgx, err := strconv.ParseInt(values[3], 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	return DatabaseRecord{int(id), name, int(age), int(dgx)}
}

// QueryService is a function that filters database records given the column
// the minimum and maximum values of the interval
type QueryService func([]DatabaseRecord, int, int, int) []DatabaseRecord

// ResponseService is a function that converts database records to a long string
type ResponseService func([]DatabaseRecord) string

// Database is a wrapper structure for data and behaviour specified
type Database struct {
	data []DatabaseRecord
	queryService QueryService
	responseService ResponseService
}

// Query is the main functionality of the Database structure as it uses all its\
// components to obtain the response of desired format
func (db Database) Query(col, min, max int) string {
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

func main() {
	if len(os.Args) < 2 {
        fmt.Println("Usage: PROGRAM database_file")
        return
    }

	queryService := func(data []DatabaseRecord, col, min, max int) []DatabaseRecord {
		ret := make([]DatabaseRecord, 0)
		for _, record := range data {
			val := record.Get(col)
			if min <= val && val <= max {
				ret = append(ret, record)
			}
		}
		return ret
	}

	responseService := func(data []DatabaseRecord) string {
		var str strings.Builder

		for _, record := range data {
			str.WriteString(record.ToString())
		}
	
		return str.String()
	}

	db := LoadDatabase(os.Args[1], queryService, responseService)

	fmt.Println(db.Query(AGE, 30, 32))
}