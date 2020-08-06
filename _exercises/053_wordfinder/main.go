package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {
	query := os.Args[1:]
	if len(query) == 0 {
		fmt.Println("Please give me a word to search.")
		return
	}

	filter := [5]string{"and", "or", "the", "since", "very"}

	words := strings.Fields(strings.ToLower(corpus))

queries:
	for _, q := range query {
		q = strings.ToLower(q)

		for _, f := range filter {
			if q == f {
				continue queries
			}
		}

		for i, w := range words {
			if q == w {
				fmt.Printf("  #%-2d: %q\n", i+1, w)
				continue queries
			}
		}
	}
}
