package database

import "strings"

// ResponseService is a function that converts database records to a long string
type ResponseService func([]DatabaseRecord) string

func BasicResponseService(data []DatabaseRecord) string {
	var str strings.Builder

	for _, record := range data {
		str.WriteString(record.ToString())
	}

	return str.String()
}
