package main

import (
	"fmt"
	"strings"
)

func main() {
	var books [3]string
	books[0] = "Kafka's Revenge"
	books[1] = "Fasola"
	books[2] = "Pizzazz"

	upper := books
	for i := 0; i < len(upper); i++ {
		upper[i] = strings.ToUpper(upper[i])
	}

	lower := books
	for i := 0; i < len(lower); i++ {
		lower[i] = strings.ToLower(lower[i])
	}

	fmt.Printf("lower = %#v\n", lower)
	fmt.Printf("upper = %#v\n", upper)

}
