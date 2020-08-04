package main

import "fmt"

func main() {
	// goto loop // goto loop jumps over declaration of i
	var i int

loop:
	if i < 3 {
		fmt.Println("Looping")
		i++
		goto loop
	}
	fmt.Println("Done")
}
