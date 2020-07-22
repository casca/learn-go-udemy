package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/01-numbers/exercises/01-do-some-calculations/main.go
	fmt.Println(50 + 25)
	fmt.Println(50 - 15.5)
	fmt.Println(50 * 0.5)
	fmt.Println(50 / 0.5)
	fmt.Println(25 % 3)
	fmt.Println(-(5 + 2))

	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/01-numbers/exercises/02-fix-the-float/main.go
	x := 5. / 2
	fmt.Println(x)

	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/01-numbers/exercises/03-precedence/main.go
	// This expression should print 20
	fmt.Println(10 + 5 - (5 - 10))

	// This expression should print -16
	fmt.Println(-10 + 0.5 - (1 + 5.5))

	// This expression should print -25
	fmt.Println(5 + 10*(2-5))

	// This expression should print 0.5
	fmt.Println(0.5 * (2 - 1))

	// This expression should print 24
	fmt.Println((3+1)/2*10 + 4)

	// This expression should print 15
	fmt.Println(10 / 2 * (10 % 7))

	// This expression should print 40
	// Note that you may need to use floats to solve this
	fmt.Println(100 / (5. / 2))

	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/01-numbers/exercises/04-incdecs/main.go
	counter, factor := 45, 0.5
	counter++
	counter++
	counter++
	counter++
	counter++
	factor--
	factor--
	fmt.Println(float64(counter) * factor)

	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/01-numbers/exercises/05-manipulate-a-counter/main.go
	var c int

	c++
	c--
	c += 5
	c *= 10
	c /= 5

	fmt.Println(c)

	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/01-numbers/exercises/06-simplify-the-assignments/main.go
	width, height := 10, 2

	width++
	width += height
	width--
	width -= height
	width *= 20
	width /= 25
	width %= 5

	fmt.Println(width)

	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/01-numbers/exercises/07-circle-area/main.go
	var (
		radius = 10.
		area   float64
	)
	area = math.Pi * radius * radius
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)

	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/01-numbers/exercises/08-sphere-area/main.go
	arg := os.Args[1]
	radius, _ = strconv.ParseFloat(arg, 64)
	area = 4 * math.Pi * radius * radius
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)

	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/01-numbers/exercises/09-sphere-volume/main.go
	var vol float64
	arg = os.Args[1]
	radius, _ = strconv.ParseFloat(arg, 64)
	vol = 4. / 3 * math.Pi * radius * radius * radius
	fmt.Printf("radius: %g -> volume: %.2f\n", radius, vol)
}
