package main

import "fmt"

var unused string

func main() {
	var anInt int
	var aBool bool
	var aFloat64 float64
	var aString string

	fmt.Println(anInt, aBool, aFloat64)
	fmt.Printf("s (%T): %q\n", aString, aString)

	// invalid identifiers
	//  var 3speed int
	//  var !speed int
	//  var spe?ed int
	//  var var int
	//  var func int
	//  var package int

	var (
		a int
		b int8
		c int16
		d int32
		e int64
		f float32
		g float64
		h complex64
		i complex128
		j bool
		k string
		l rune
		m byte
	)

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m)

	var (
		active int
		delta  int
	)
	fmt.Println(active, delta)

	var firstName, lastName string
	fmt.Printf("%q %q\n", firstName, lastName)

	var isLiquid bool
	_ = isLiquid

	// fmt.Println(tooSoon)
	// var tooSoon string
}
