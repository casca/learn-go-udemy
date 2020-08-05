package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		names     [3]string
		distances [5]int
		data      [5]uint8
		ratios    [1]float64
		alives    [4]bool
		zero      [0]uint8
	)

	names[0] = "Einstein"
	names[1] = "Tesla"
	names[2] = "Shepard"

	distances[0] = 50
	distances[1] = 40
	distances[2] = 75
	distances[3] = 30
	distances[4] = 125

	data[0] = 'H'
	data[1] = 'E'
	data[2] = 'L'
	data[3] = 'L'
	data[4] = 'O'

	ratios[0] = 3.14145

	alives[0] = true
	alives[1] = false
	alives[2] = true
	alives[3] = false

	// zero[0] = "BOMB!"
	_ = zero

	separator := strings.Repeat("=", 20)

	fmt.Println("names")
	fmt.Println(separator)
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d] = %q\n", i, names[i])
	}
	fmt.Println()

	fmt.Println("distances")
	fmt.Println(separator)
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d] = %v\n", i, distances[i])
	}
	fmt.Println()

	fmt.Println("data")
	fmt.Println(separator)
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d] = %v\n", i, data[i])
	}
	fmt.Println()

	fmt.Println("ratios")
	fmt.Println(separator)
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d] = %v\n", i, ratios[i])
	}
	fmt.Println()

	fmt.Println("alives")
	fmt.Println(separator)
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d] = %v\n", i, alives[i])
	}
	fmt.Println()

	fmt.Println("zero")
	fmt.Println(separator)
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d] = %v\n", i, zero[i])
	}
	fmt.Println()
}
