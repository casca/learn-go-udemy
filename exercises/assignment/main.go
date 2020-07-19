package main

import (
	"fmt"
	"path"
)

func main() {
	// https://github.com/inancgumus/learngo/blob/master/06-variables/04-assignment/exercises/01-make-it-blue/main.go
	color := "green"
	color = "blue"
	fmt.Println(color)

	// https://github.com/inancgumus/learngo/blob/master/06-variables/04-assignment/exercises/02-vars-to-vars/main.go
	color = "green"
	color = "dark " + color
	fmt.Println(color)

	// https://github.com/inancgumus/learngo/blob/master/06-variables/04-assignment/exercises/03-assign-with-expressions/main.go
	n := 0.
	n = 3.14 * 2
	fmt.Println(n)

	// https://github.com/inancgumus/learngo/blob/master/06-variables/04-assignment/exercises/04-find-the-rectangle-perimeter/main.go
	var (
		perimeter     int
		width, height = 5, 6
	)
	perimeter = 2 * (width + height)
	fmt.Println(perimeter)

	// https://github.com/inancgumus/learngo/blob/master/06-variables/04-assignment/exercises/05-multi-assign/main.go
	var (
		lang    string
		version int
	)
	lang, version = "go", 2
	fmt.Println(lang, "version", version)

	// https://github.com/inancgumus/learngo/blob/master/06-variables/04-assignment/exercises/06-multi-assign-2/main.go
	var (
		planet string
		isTrue bool
		temp   float64
	)
	planet, isTrue, temp = "Mars", true, 19.5
	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")

	// https://github.com/inancgumus/learngo/blob/master/06-variables/04-assignment/exercises/07-multi-short-func/main.go
	_, b := multi()
	fmt.Println(b)

	// https://github.com/inancgumus/learngo/blob/master/06-variables/04-assignment/exercises/08-swapper/main.go
	color, color2 := "red", "blue"
	color, color2 = "orange", "green"
	fmt.Println(color, color2)

	// https://github.com/inancgumus/learngo/blob/master/06-variables/04-assignment/exercises/09-swapper-2/main.go
	red, blue := "red", "blue"
	red, blue = blue, red
	fmt.Println(red, blue)

	// https://github.com/inancgumus/learngo/blob/master/06-variables/04-assignment/exercises/10-discard-the-file/main.go
	dir, _ := path.Split("secret/file.txt")
	fmt.Println(dir)
}

func multi() (int, int) {
	return 5, 4
}
