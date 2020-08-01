package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("%q\n", os.Args[i])
	}

	for _, v := range os.Args {
		fmt.Printf("%q\n", v)
	}

	words := strings.Fields("lazy cat jupms again and again and again")
	for j := 0; j < len(words); j++ {
		fmt.Printf("#%-2d: %q\n", j+1, words[j])
	}

	var (
		i int
		v string
	)
	for i, v = range words {
		fmt.Printf("#%-2d: %q\n", i+1, v)
	}
	fmt.Printf("Last value of i = %d and v = %q\n", i, v)
}
