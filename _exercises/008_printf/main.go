package main

import (
	"fmt"
	"os"
)

func main() {
	// https://github.com/inancgumus/learngo/blob/master/07-printf/exercises/01-print-your-age/main.go
	fmt.Printf("I'm %d years old.\n", 36)

	// https://github.com/inancgumus/learngo/blob/master/07-printf/exercises/02-print-your-name-and-lastname/main.go
	f := "My name is %s and my lastname is %s.\n"
	fmt.Printf(f, "John", "Doe")

	// https://github.com/inancgumus/learngo/blob/master/07-printf/exercises/03-false-claims/main.go
	tr := false
	fmt.Printf("These are %t claims.\n", tr)

	// https://github.com/inancgumus/learngo/blob/master/07-printf/exercises/04-print-the-temperature/main.go
	fmt.Printf("Temperature if %.1f degrees.\n", 29.5)

	// https://github.com/inancgumus/learngo/blob/master/07-printf/exercises/05-double-quotes/main.go
	fmt.Printf("%q\n", "hello world")

	// https://github.com/inancgumus/learngo/blob/master/07-printf/exercises/06-print-the-type/main.go
	fmt.Printf("Type of %d is %[1]T\n", 3)

	// https://github.com/inancgumus/learngo/blob/master/07-printf/exercises/07-print-the-type-2/main.go
	fmt.Printf("Type of %.2f is %[1]T\n", 3.14)

	// https://github.com/inancgumus/learngo/blob/master/07-printf/exercises/08-print-the-type-3/main.go
	fmt.Printf("Type of %s is %[1]T\n", "hello")

	// https: //github.com/inancgumus/learngo/blob/master/07-printf/exercises/09-print-the-type-4/main.go
	fmt.Printf("Type of %t is %[1]T\n", true)

	// https://github.com/inancgumus/learngo/blob/master/07-printf/exercises/10-print-your-fullname/main.go
	fmt.Printf("Your name is %s and your latname is %s\n", os.Args[1], os.Args[2])
}
