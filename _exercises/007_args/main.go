package main

import (
	"fmt"
	"os"
)

func main() {
	count := len(os.Args) - 1
	fmt.Printf("There are %d names.\n", count)

	fmt.Println("Path:", os.Args[0])

	fmt.Printf("Let me guess your name... %s?\n", os.Args[1])

	fmt.Printf("There are %d people.\n", len(os.Args)-1)
	for _, n := range os.Args[1:] {
		fmt.Printf("Hello %s!\n", n)
	}
	fmt.Println("Nice to meet you all.")
}
