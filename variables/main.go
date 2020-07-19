package main

import (
	"fmt"
	"time"
)

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

	assignment1()
	assignment2()
	assignment3()
}

func assignment1() {
	fmt.Println("-- assignment 1")
	var speed int
	fmt.Println(speed)

	speed = 100
	fmt.Println(speed)

	speed = speed - 25
	fmt.Println(speed)

	// speed = "100"

	// var running bool
	// running = 1

	var force float64
	// speed = force

	force = 1
	fmt.Println(force)
}

func assignment2() {
	fmt.Println("-- assignment 2")
	var (
		name   string
		age    int
		famous bool
	)

	name = "Newton"
	age = 84
	famous = true
	fmt.Println(name, age, famous)

	name = "Somebody"
	age = 20
	famous = false
	fmt.Println(name, age, famous)

	// prevName := name
	var prevName string
	prevName = name

	name = "Einstein"

	fmt.Println("previous name:", prevName)
	fmt.Println("current name:", name)

	var (
		speed int
		now   time.Time
	)

	speed, now = 100, time.Now()
	fmt.Println(speed, now)
}

func assignment3() {
	fmt.Println("-- assignment 3")
	var (
		speed     = 100
		prevSpeed = 50
	)

	speed, prevSpeed = prevSpeed, speed
	fmt.Println(speed, prevSpeed)

}
