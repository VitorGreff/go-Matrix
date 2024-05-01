package concurrent

import (
	"errors"
	"parallel/utils"
	"sync"
)

// needed to populate the resulting matrix properly
type Result struct {
	value float64
	row   int
	col   int
}

func Multiply(a, b utils.Matrix) (utils.Matrix, error) {
	if a.Cols != b.Rows {
		return utils.Matrix{}, errors.New("invalid dimensions")
	}

	result := make([][]float64, a.Rows)
	ch := make(chan Result)
	var wg sync.WaitGroup

	for i := range a.Rows {
		result[i] = make([]float64, b.Cols)
		for j := range b.Cols {
			// create green thread
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				var sum float64
				for k := range a.Cols {
					sum += a.Data[i][k] * b.Data[k][j]
				}
				ch <- Result{value: sum, row: i, col: j}
			}(i, j)
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		result[i.row][i.col] = i.value
	}
	return utils.Matrix{Data: result, Rows: a.Rows, Cols: b.Cols}, nil
}
