package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.ReadFile("data.dat")
	if err != nil {
		panic(err)
	}
	A, C, T, G := 0, 0, 0, 0
	for i := 0; i < len(file); i++ {
		if file[i] == 65 {
			A += 1
		} else if file[i] == 67 {
			C += 1
		} else if file[i] == 71 {
			G += 1
		} else if file[i] == 84 {
			T += 1
		}
	}
	fmt.Printf("A = %d\nC = %d\nG = %d\nT = %d\n", A, C, G, T)
	elapsed := time.Since(start)

	log.Printf("The counter took %s", elapsed)
}
