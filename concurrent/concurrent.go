package concurrent

import (
	"errors"
	"parallel/utils"
	"sync"
)

// needed to populate the resulting matrix properly
type rowResult struct {
	value []float64
	row   int
}

func Multiply(a, b utils.Matrix) (utils.Matrix, error) {
	if a.Cols != b.Rows {
		return utils.Matrix{}, errors.New("invalid dimensions")
	}

	result := make([][]float64, a.Rows)
	ch := make(chan rowResult)
	var wg sync.WaitGroup

	for i := 0; i < a.Rows; i++ {
		result[i] = make([]float64, b.Cols)
		// creates green thread
		wg.Add(1)
		go func(i int) {
			aux := rowResult{}
			aux.row = i
			defer wg.Done()
			for j := 0; j < b.Cols; j++ {
				var sum float64
				for k := 0; k < a.Cols; k++ {
					sum += a.Data[i][k] * b.Data[k][j]
				}
				aux.value = append(aux.value, sum)
			}
			ch <- aux
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		result[i.row] = i.value
	}
	return utils.Matrix{Data: result, Rows: a.Rows, Cols: b.Cols}, nil
}
