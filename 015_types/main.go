package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		width  uint8 = 255 // max value
		height       = 255
	)

	width++

	fmt.Println(width)
	if int(width) < height {
		fmt.Println("height is greater")
	}

	// int8(math.MaxInt8 + 1)
	var n int8 = math.MaxInt8
	fmt.Println("max int8     :", n)
	fmt.Println("max int8 + 1 :", n+1)

	n = math.MinInt8
	fmt.Println("max int8     :", n)
	fmt.Println("max int8 - 1 :", n-1)

	var un uint8
	fmt.Println("max uint8     :", un)
	fmt.Println("max uint8 - 1 :", un-1)

	f := float32(math.MaxFloat32)
	fmt.Println("max float     :", f*1.2)
}
