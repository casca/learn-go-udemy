package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("give me numbers")
		return
	}

	args := os.Args[1:]
	for _, arg := range args {
		n, err := strconv.Atoi(arg)
		if err != nil || !isPrime(n) {
			continue
		}
		fmt.Printf("%d ", n)

	}
	fmt.Println()
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
