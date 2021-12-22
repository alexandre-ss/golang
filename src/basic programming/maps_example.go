package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := os.Args[1:] // get all words passed as argument

	estatistics := get_stat(words)

	print_stats(estatistics)
}

func get_stat(words []string) map[string]int {
	stats := make(map[string]int)
	for _, word := range words {
		initial := strings.ToUpper(string(word[0]))
		counter, found := stats[initial]
		if found {
			stats[initial] = counter + 1
		} else {
			stats[initial] = 1
		}
	}
	return stats
}

func print_stats(stats map[string]int) {
	fmt.Println("How many words start with the same first letter?")

	for initial, counter := range stats {
		fmt.Printf("%s = %d\n", initial, counter)
	}
}
