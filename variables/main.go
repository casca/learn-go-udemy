package main

import "fmt"

var mightBeUnused int

func main() {
	fmt.Println(-200, -100, 0, 50, 100)

	fmt.Println(-50.5, -20.5, -0.0, 1.0, 100.56)

	fmt.Println(true, false)

	fmt.Print(
		"",          // empty string
		"hi",        // it contains 'hi'
		"nasılsın?", // "how are you?" in Turkish
	)
	fmt.Println()

	// declaration syntax
	var (
		speed, velocity int
		heat            float64

		off   bool
		brand string
	)

	fmt.Println(speed, velocity)
	fmt.Println(heat)
	fmt.Println(off)
	fmt.Printf("%q\n", brand)

	var unused int
	_ = unused

	var myAge, yourAge int
	// var yourAge int
	var temperature float64
	var success bool
	var language string
	fmt.Println(myAge, yourAge, temperature, success, language)
}
