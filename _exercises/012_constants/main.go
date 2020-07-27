package main

import (
	"fmt"
	"time"
)

func main() {
	// https://github.com/inancgumus/learngo/blob/master/10-constants/exercises/01-minutes-in-weeks/main.go
	const (
		minsPerDay = 60 * 24
		weekDays   = 7
	)
	fmt.Println(minsPerDay * weekDays * 2)

	// https://github.com/inancgumus/learngo/blob/master/10-constants/exercises/02-remove-the-magic/main.go
	const (
		hoursInDay   = 24
		daysInWeek   = 7
		hoursPerWeek = hoursInDay * daysInWeek
	)
	fmt.Printf("There are %d hours in %d weeks.\n",
		hoursPerWeek*5, 5)

	// https://github.com/inancgumus/learngo/blob/master/10-constants/exercises/03-constant-length/main.go
	const (
		home   = "Milky Way Galaxy"
		length = len(home)
	)
	fmt.Printf("There are %d characters inside %q\n", length, home)

	// https://github.com/inancgumus/learngo/blob/master/10-constants/exercises/04-tau/main.go
	const (
		pi  = 3.14159265358979323846264
		tau = pi * 2
	)
	fmt.Printf("%g\n", tau)

	// https://github.com/inancgumus/learngo/blob/master/10-constants/exercises/05-area/main.go
	const (
		width  = 25
		height = width * 2
	)
	fmt.Printf("area = %d\n", width*height)

	// https://github.com/inancgumus/learngo/blob/master/10-constants/exercises/06-no-conversions-allowed/main.go
	const later = 10
	hours, _ := time.ParseDuration("1h")
	fmt.Printf("%s later...\n", hours*later)

	// https://github.com/inancgumus/learngo/blob/master/10-constants/exercises/07-iota-months/main.go
	const (
		Nov = 11 - iota
		Oct
		Sep
	)
	fmt.Println(Sep, Oct, Nov)

	// https://github.com/inancgumus/learngo/blob/master/10-constants/exercises/08-iota-months-2/main.go
	const (
		_ = iota
		Jan
		Feb
		Mar
	)
	fmt.Println(Jan, Feb, Mar)

	// https://github.com/inancgumus/learngo/blob/master/10-constants/exercises/09-iota-seasons/main.go
	// const (
	// 	Winter = 12 - (9*iota)%12
	// 	Spring
	// 	Summer
	// 	Fall
	// )
	const (
		Spring = (iota + 1) * 3
		Summer
		Fall
		Winter
	)
	fmt.Println(Winter, Spring, Summer, Fall)
}
