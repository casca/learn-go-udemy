package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// https://github.com/inancgumus/learngo/blob/master/12-switch/exercises/01-richter-scale/main.go
	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}

	r, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}

	var desc string
	switch {
	case r < 2.0:
		desc = "micro"
	case r >= 2.0 && r < 3.0:
		desc = "very minor"
	case r >= 3.0 && r < 4.0:
		desc = "minor"
	case r >= 4.0 && r < 5.0:
		desc = "light"
	case r >= 5.0 && r < 6.0:
		desc = "moderate"
	case r >= 6.0 && r < 7.0:
		desc = "strong"
	case r >= 7.0 && r < 8.0:
		desc = "major"
	case r >= 8.0 && r < 10.0:
		desc = "great"
	default:
		desc = "massive"
	}
	fmt.Printf("%.2f is %s\n", r, desc)

}
