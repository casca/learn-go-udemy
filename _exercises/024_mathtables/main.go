package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	usage     = "Usage: [op=*/+-] [size]"
	validOps  = "%*/+-"
	invalidOp = -1
)

func main() {
	switch l := len(os.Args); {
	case l == 1:
		fmt.Println("Size is missing")
		fallthrough
	case l == 0 || l > 3:
		fmt.Println(usage)
		return
	}

	op := os.Args[1]
	if strings.IndexAny(op, validOps) == invalidOp {
		fmt.Println("Invalid operator.")
		fmt.Println("Valid ops one of:", validOps)
		return
	}
	fmt.Println("fo")

	size, err := strconv.Atoi(os.Args[2])
	if err != nil || size < 0 {
		fmt.Println("Wrong size")
		return
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= size; j++ {
			var res int
			switch op {
			case "*":
				res = i * j
			case "%":
				if j != 0 {
					res = i % j
				}
			case "/":
				if j != 0 {
					res = i / j
				}
			case "+":
				res = i + j
			case "-":
				res = i - j
			}
			fmt.Printf("%5d", res)
		}
		fmt.Println()
	}
}
