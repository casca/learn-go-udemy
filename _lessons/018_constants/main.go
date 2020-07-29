package main

import (
	"fmt"
)

func main() {
	const meters int = 100

	cm := 100
	m := cm / meters
	fmt.Printf("%dcm is %dm\n", cm, m)

	cm = 200
	m = cm / meters
	fmt.Printf("%dcm is %dm\n", cm, m)

	// total, numberOf := 5, 0
	// const total int = 5
	// const numberOf int = 0
	// fmt.Println(total / numberOf) // division by zero

	// const max int = 100
	// max = 100

	// const max int = math.Pow10(2) // const initializer math.Pow10(2) is not a constant

	// n := 2
	// const max int = n // const initializer n is not a constant

	const max int = len("Hello") // valid initializer
	fmt.Println(max)

	// const min int = 1
	const min = 1 + 1
	// const pi float64 = 3.14
	const pi = 3.14 * min
	// const version string = "2.0.1"
	const version = "2.0.1" + "-beta"
	// const debug bool = true
	const debug = !true
	fmt.Println(min, pi, version, debug)

	// const mi, ma int = 1, 1000
	const (
		mi int = 1 + 1
		ma
	)
	fmt.Println(mi, ma)
}
