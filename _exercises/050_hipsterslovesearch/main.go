package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	if len(os.Args) != 2 {
		fmt.Println("Tell me a book title")
		return
	}

	query := strings.ToLower(os.Args[1])

	fmt.Println("Search Results:")
	var found bool
	for _, b := range books {
		if strings.Contains(strings.ToLower(b), query) {
			found = true
			fmt.Printf("+ %s\n", b)
		}
	}
	if !found {
		fmt.Printf("We don't have the book: %q\n", query)
	}
}
