package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// https://github.com/inancgumus/learngo/blob/master/11-if/exercises/06-odd-even/main.go

	if len(os.Args) != 2 {
		fmt.Println("Pick a number")
		return
	}
	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a number\n", os.Args[1])
		return
	}
	if number%2 == 0 {
		fmt.Printf("%d is an even number\n", number)
	} else {
		fmt.Printf("%d is an odd number\n", number)
	}
}
