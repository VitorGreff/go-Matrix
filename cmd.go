package main

import (
	"flag"
	"log"
	"parallel/concurrent"
	"parallel/file"
	"parallel/sequential"
	"parallel/utils"
	"time"
)

func main() {
	var (
		rowN int
		colN int
	)
	flag.IntVar(&rowN, "r", 800, "defines the number of rows of each matrix")
	flag.IntVar(&colN, "c", 800, "defines the number of columns of each matrix")
	flag.Parse()

	m1 := utils.CreateRandomMatrix(rowN, colN)
	m2 := utils.CreateRandomMatrix(rowN, colN)

	start := time.Now()
	m3, err := sequential.Multiply(m1, m2)
	if err != nil {
		panic(err)
	}
	file.WriteMatrix(m3.Data, "1.txt")
	elapsed := time.Since(start)
	log.Printf("Sequential time -> %f seconds", elapsed.Seconds())

	start = time.Now()
	m4, err := sequential.Multiply(m1, m2)
	if err != nil {
		panic(err)
	}
	file.WriteMatrix(m4.Data, "2.txt")
	concurrent.Multiply(m1, m2)
	elapsed = time.Since(start)
	log.Printf("Concurrent time -> %f seconds", elapsed.Seconds())
}
