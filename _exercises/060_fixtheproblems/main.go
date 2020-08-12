package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// var names []string
	// names := []string{}
	names := []string{
		"Einstein",
		"Shepard",
		"Tesla",
	}

	var books []string = []string{
		"Stay Golden",
		"Fire",
		"Kafka's Revenge",
	}
	sort.Strings(books)

	nums := [...]int{5, 1, 7, 3, 8, 2, 6, 9}
	sort.Ints(nums[:]) // array to slice!!

	fmt.Printf("%q\n", strings.Join(names, " and "))
	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
}
