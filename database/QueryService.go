package database

// QueryService is a function that filters database records given the column
// the minimum and maximum values of the interval
type QueryService func([]DatabaseRecord, int, int, int) []DatabaseRecord

func BasicQueryService() QueryService {
	return func(data []DatabaseRecord, col, min, max int) []DatabaseRecord {
		ret := make([]DatabaseRecord, 0)
		for _, record := range data {
			val := record.Get(col)
			if min <= val && val <= max {
				ret = append(ret, record)
			}
		}
		return ret
	}
}

func ContrainedRangeQueryService(minLimit, maxLimit int) QueryService {
	return func(data []DatabaseRecord, col, min, max int) []DatabaseRecord {
		if max-min < minLimit || max-min > maxLimit {
			return []DatabaseRecord{}
		}
		ret := make([]DatabaseRecord, 0)
		for _, record := range data {
			val := record.Get(col)
			if min <= val && val <= max {
				ret = append(ret, record)
			}
		}
		return ret
	}
}
