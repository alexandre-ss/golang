package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	entry_count := os.Args[1:]
	numbers := make([]int, len(entry_count))

	for index, number := range entry_count {
		number_item, err := strconv.Atoi(number)

		if err != nil {
			fmt.Printf("%s is not a valid number!\n", number)
			os.Exit(1)
		}
		numbers[index] = number_item
	}
	fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	numbers_aux := make([]int, len(numbers))
	copy(numbers_aux, numbers)

	reference_number_index := len(numbers_aux) / 2
	reference_number := numbers[reference_number_index]
	numbers_aux = append(numbers_aux[:reference_number_index], numbers_aux[reference_number_index+1:]...)

	smaller, bigger := partition(numbers_aux, reference_number)

	return append(
		append(quicksort(smaller), reference_number),
		quicksort(bigger)...)
}

func partition(numbers []int, reference_number int) (smaller []int, bigger []int) {
	for _, number := range numbers {
		if number <= reference_number {
			smaller = append(smaller, number)
		} else {
			bigger = append(bigger, number)
		}
	}
	return smaller, bigger
}
