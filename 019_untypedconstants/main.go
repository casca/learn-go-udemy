package main

import (
	"fmt"
	"time"
)

func main() {
	const min = 1 + 1     // untyped/typeless constant
	const pi = 3.14 * min // compiler doesn't complain
	fmt.Println(min, pi)

	const mi int = 42
	var i int
	i = min
	fmt.Println(i)
	// var f float64
	// f = mi
	// fmt.Println(f)

	const m = 42
	var ii int = m
	var f float64 = m
	var b byte = m
	var j int32 = m
	var r rune = m
	fmt.Println(ii, f, b, j, r)

	// const mu int32 = 1
	// max := 5 + min // same as max := int32(5) + min
	// fmt.Println(max)

	const mu = 1
	max := 5 + min // same as max:= int(5) + int(min)
	fmt.Printf("%T %v\n", max, max)

	// default types
	iii := 42 // int
	fmt.Printf("%T %v\n", iii, iii)
	fff := 3.14 // float64
	fmt.Printf("%T %v\n", fff, fff)
	bbb := true // bool
	fmt.Printf("%T %v\n", bbb, bbb)
	sss := "Hello" // string
	fmt.Printf("%T %v\n", sss, sss)
	rrr := 'A' // rune aka int32
	fmt.Printf("%T %v\n", rrr, rrr)

	// the conversion happens only when a type is needed
	iiii := 42 + 3.14 // same as float64(42 + 3.14), but the result is a typeless float value
	fmt.Printf("%T %v\n", iiii, iiii)

	// hh := true + "Hello" // incompatible values
	// fmt.Printf("%T %v\n", hh, hh)

	t := time.Second * 10
	fmt.Printf("%T %v\n", t, t)

	iiiii := 10
	// t = time.Second * i // mismatched types
	t = time.Second * time.Duration(iiiii)
	fmt.Println(t)

	const c = 10
	t = time.Second * c
	fmt.Println(t)
}
