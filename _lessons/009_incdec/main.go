package main

import "fmt"

func main() {
	var n int

	// increment
	n = n + 1
	fmt.Println(n)

	n += 1
	fmt.Println(n)

	n++
	fmt.Println(n)

	// decrement
	n = 10

	n = n - 1
	fmt.Println(n)

	n -= 1
	fmt.Println(n)

	n--
	fmt.Println(n)

	// ---
	var counter int
	counter++
	counter++
	counter++
	counter--
	fmt.Printf("There are %d line(s) in the file\n", counter)

	// counter = counter++ + counter-- // not allowed
}
