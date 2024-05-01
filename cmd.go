package main

import (
	"log"
	"parallel/concurrent"
	"parallel/sequential"
	"parallel/utils"
)

func main() {
	m1 := utils.CreateRandomMatrix(2, 2)
	m2 := utils.CreateRandomMatrix(2, 2)
	m3, _ := sequential.Multiply(m1, m2)
	m4, _ := concurrent.Multiply(m1, m2)

	for r := range m3.Rows {
		log.Println(m3.Data[r])
	}
	log.Println()
	for r := range m4.Rows {
		log.Println(m4.Data[r])
	}
}
