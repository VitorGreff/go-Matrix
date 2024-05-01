package main

import (
	"flag"
	"log"
	"parallel/concurrent"
	"parallel/file"
	"parallel/sequential"
	"parallel/utils"
	"runtime"
	"time"
)

func main() {
	var (
		rowN int
		colN int
		tN   int
	)
	flag.IntVar(&rowN, "r", 800, "defines the number of rows of each matrix")
	flag.IntVar(&colN, "c", 800, "defines the number of columns of each matrix")
	flag.IntVar(&tN, "t", 16, "defines the maximum number of threads used by the concurrent function")
	flag.Parse()

	runtime.GOMAXPROCS(tN)

	m1 := utils.CreateRandomMatrix(rowN, colN)
	m2 := utils.CreateRandomMatrix(rowN, colN)

	start := time.Now()
	m3, err := sequential.Multiply(m1, m2)
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	file.WriteMatrix(m3.Data, "sequential.txt")
	log.Printf("Sequential time -> %f seconds", elapsed.Seconds())

	start = time.Now()
	m4, err := concurrent.Multiply(m1, m2)
	if err != nil {
		panic(err)
	}
	elapsed = time.Since(start)
	file.WriteMatrix(m4.Data, "concurrent.txt")
	log.Printf("Concurrent time -> %f seconds", elapsed.Seconds())
}
