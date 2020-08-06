package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 5 {
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	}

	var sum float64
	var count float64
	for _, a := range args {
		n, err := strconv.ParseFloat(a, 64)
		if err != nil {
			continue
		}
		sum += n
		count++
	}

	fmt.Printf("  Your numbers: %v\n", args)
	fmt.Printf("  Average: %g\n", sum/count)
}
