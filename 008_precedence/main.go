package main

import "fmt"

func main() {
	fmt.Println(2 + 2*4/2) // same as 2 + ((2 * 4) / 2)

	fmt.Println(1 + 5 - 3*10/2)
	fmt.Println((1 + 5 - 3) * (10 / 2))

	celsius := 35.
	fahrenheit := (9*celsius + 160) / 5
	fmt.Printf("%g C is %g F\n", celsius, fahrenheit)
}
