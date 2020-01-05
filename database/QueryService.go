package database

// QueryService is a function that filters database records given the column
// the minimum and maximum values of the interval
type QueryService func([]DatabaseRecord, int, int, int) []DatabaseRecord

func BasicQueryService(data []DatabaseRecord, col, min, max int) []DatabaseRecord {
	ret := make([]DatabaseRecord, 0)
	for _, record := range data {
		val := record.Get(col)
		if min <= val && val <= max {
			ret = append(ret, record)
		}
	}
	return ret
}
