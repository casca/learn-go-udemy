package main

import "fmt"

var mightBeUnused int

// sfe := true
var sfe = true

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

	// initializing variables - short declaration
	// var safe bool = true
	safe := true
	fmt.Println(safe)

	// var name string = "Carl"
	// var name = "Carl"
	name := "Carl"

	// var isScientist bool = true
	// var isScientist = true
	isScientist := true

	// var age int = 62
	// var age = 62
	age := 62

	// var degree float64 = 5.
	// var degree = 5.
	degree := 5.

	fmt.Println(name, isScientist, age, degree)

	sf, spd := true, 50
	fmt.Println(sf, spd)

	nm, lstnm := "Nikola", "Tesla"
	fmt.Println(nm, lstnm)

	birth, death := 1856, 1943
	fmt.Println(birth, death)

	on, off := true, false
	fmt.Println(on, off)

	d, r, h := 10.55, 30.5, 20.
	fmt.Println(d, r, h)

	nFiles, valid, msg := 10, true, "hello"
	fmt.Println(nFiles, valid, msg)

	// redeclaration
	var se bool
	se, sd := true, 50 // at least one of the variables should be a new variable
	fmt.Println(se, sd)

	name, ae := "Marie", 66
	fmt.Println(name, ae)

	// name = "Albert" // assignment
	// bh := 1879      // declaration
	name, bh := "Albert", 1879 // redeclaration
	fmt.Println(name, bh)

	var (
		// related:
		video string

		// closely related:
		duration int
		current  int
	)

	fmt.Println(video, duration, current)
}
