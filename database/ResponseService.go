package database

import "math/rand"

// ResponseService is a function that converts database records to a long string
type ResponseService func([]DatabaseRecord) int

func BasicResponseService() ResponseService {
	return func(data []DatabaseRecord) int {
		return len(data)
	}
}

func SpuriousVolumesResponseService(x int) ResponseService {
	return func(data []DatabaseRecord) int {
		return len(data) + x
	}
}

func RandomSpuriousVolumesResponseService(r rand.Rand, n int) ResponseService {
	return func(data []DatabaseRecord) int {
		return len(data) + r.Intn(n)
	}
}
