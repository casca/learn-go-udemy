package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// https://github.com/inancgumus/learngo/blob/master/09-go-type-system/exercises/01-optimal-types/main.go
	// an english letter (search web for: ascii control code)
	// byte

	// a non-english letter (search web for: unicode codepoint)
	// rune

	// a year in 4 digits like 2040
	// uint16

	// a month in 2 digits: 1 to 12
	// uint8

	// the speed of the light
	// uint32

	// angle of a circle
	// float32

	// PI number: 3.141592653589793
	// float64

	// https://github.com/inancgumus/learngo/blob/master/09-go-type-system/exercises/02-the-type-problem/main.go
	var (
		width  uint16
		height uint16
	)

	width, height = 255, 265
	width += 10
	fmt.Printf("width: %d height: %d\n", width, height)
	fmt.Println("are they equal?", width == height)

	// https://github.com/inancgumus/learngo/blob/master/09-go-type-system/exercises/03-parse-arg-numbers/main.go
	// go run main.go 50 25000 2000000 50000000000000000 00000100
	// val, _ := strconv.ParseInt(os.Args[1], 10, 8)
	// fmt.Println("int8 value is:", int8(val))

	// val, _ = strconv.ParseInt(os.Args[2], 10, 16)
	// fmt.Println("int16 value is:", int16(val))

	// val, _ = strconv.ParseInt(os.Args[3], 10, 32)
	// fmt.Println("int32 value is:", int32(val))

	// val, _ = strconv.ParseInt(os.Args[4], 10, 64)
	// fmt.Println("int64 value is:", int64(val))

	// val, _ = strconv.ParseInt(os.Args[5], 2, 8)
	// fmt.Printf("%s is: %d\n", os.Args[5], int8(val))

	// https://github.com/inancgumus/learngo/blob/master/09-go-type-system/exercises/04-time-multiplier/main.go
	// go run main.go 2
	// t, _ := time.ParseDuration("1h30m")
	// mult, _ := strconv.ParseInt(os.Args[1], 10, 64)
	// t *= time.Duration(mult)
	// fmt.Println(t)

	// https://github.com/inancgumus/learngo/blob/master/09-go-type-system/exercises/05-refactor-feet-to-meter/main.go
	type Feet float64
	type Meters float64
	var (
		feet   Feet
		meters Meters
	)
	val, _ := strconv.ParseFloat(os.Args[1], 64)
	feet = Feet(val)
	meters = Meters(feet) * 0.3048
	fmt.Printf("%g feet is %g meters.\n", feet, meters)

	// https://github.com/inancgumus/learngo/blob/master/09-go-type-system/exercises/06-convert-the-types/main.go
	type (
		Temperature float64
		Celsius     Temperature
		Fahrenheit  Temperature
	)

	var (
		celsius       Celsius     = 15.5
		fahr          Fahrenheit  = 59.9
		celsiusDegree Temperature = 10.5
		fahrDegree    Temperature = 2.5
		factor                    = 2.
	)

	celsius *= Celsius(float64(celsiusDegree) * factor)
	fahr *= Fahrenheit(float64(fahrDegree) * factor)

	fmt.Println(celsius, fahr)
}
