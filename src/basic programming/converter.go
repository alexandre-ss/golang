// celsius to fahrenheit example to train the golang sintaxe
// to run:
// go run converter.go <value> <unit of measurement>
//
// for example: go run converter.go 19 celsius
// expected output: 19.00 celsius = 66.20 fahrenheit

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Missing arguments, expected more than 3!")
		os.Exit(1)
	}

	origin_unit := os.Args[len(os.Args)-1]
	origin_value := os.Args[1 : len(os.Args)-1]

	var destination_unit string

	if origin_unit == "celsius" {
		destination_unit = "fahrenheit"
	} else {
		fmt.Printf("%s is not a valid unit of measurement",
			origin_unit)
		os.Exit(1)
	}

	for index, value := range origin_value {
		origin_value, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Printf("The value %s in the %dth position is not valid!\n",
				value, index)
			os.Exit(1)
		}
		var destination_value float64
		if origin_unit == "celsius" {
			destination_value = origin_value*1.8 + 32
		} else {
			destination_value = origin_value / 1.60934
		}
		fmt.Printf("%.2f %s = %.2f %s\n",
			origin_value, origin_unit, destination_value, destination_unit)
	}
}
