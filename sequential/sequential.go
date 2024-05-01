package sequential

import (
	"errors"
	"parallel/utils"
)

func Multiply(a, b utils.Matrix) (utils.Matrix, error) {
	if a.Cols != b.Rows {
		return utils.Matrix{}, errors.New("invalid dimensions")
	}

	result := make([][]float64, a.Rows)
	for i := range a.Rows {
		result[i] = make([]float64, b.Rows)
		for j := range b.Cols {
			for k := range a.Cols {
				result[i][j] += a.Data[i][k] * b.Data[k][j]
			}
		}
	}

	return utils.Matrix{Data: result, Rows: a.Rows, Cols: b.Cols}, nil
}
