package database

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	// ID column of DatabaseRecord
	ID = iota // 0

	// AGE column of DatabaseRecord
	AGE = iota // 1

	// DGX column of DatabaseRecord
	DGX = iota // 2
)

// DatabaseRecord is a wrapper structure for data records from the database;
// can be replaced with a simple byte slice
type DatabaseRecord struct {
	ID   int
	Name string
	Age  int
	Dgx  int
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
