package main

import "fmt"

func main() {
	// https://github.com/inancgumus/learngo/blob/master/06-variables/05-type-conversion/exercises/01-convert-and-fix/main.go
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)

	// https://github.com/inancgumus/learngo/blob/master/06-variables/05-type-conversion/exercises/02-convert-and-fix-2/main.go
	a, b = 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)

	// https://github.com/inancgumus/learngo/blob/master/06-variables/05-type-conversion/exercises/03-convert-and-fix-3/main.go
	fmt.Println(5.5)

	// https://github.com/inancgumus/learngo/blob/master/06-variables/05-type-conversion/exercises/04-convert-and-fix-4/main.go
	age := 2
	fmt.Println(7.5 + float64(age))

	// https://github.com/inancgumus/learngo/blob/master/06-variables/05-type-conversion/exercises/05-convert-and-fix-5/main.go
	min := int8(127)
	max := int16(1000)
	fmt.Println(max + int16(min))
}
