package utils

import (
	"math/rand"
)

type Matrix struct {
	Data [][]float64
	Rows int
	Cols int
}

func CreateRandomMatrix(rowN, colN int) Matrix {
	result := make([][]float64, rowN)

	for r := range rowN {
		result[r] = make([]float64, colN)
		for c := range colN {
			result[r][c] = rand.Float64()
		}
	}

	return Matrix{Data: result, Rows: rowN, Cols: colN}
}
