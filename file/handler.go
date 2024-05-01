package file

import (
	"fmt"
	"os"
)

func WriteMatrix(m [][]float64, filename string) {
	file, err := os.Create("output/" + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// index, row itself
	for _, row := range m {
		for _, value := range row {
			if _, err := fmt.Fprintf(file, "%v ", value); err != nil {
				panic(err)
			}
		}
		if _, err := fmt.Fprintln(file); err != nil {
			panic(err)
		}
	}
}
