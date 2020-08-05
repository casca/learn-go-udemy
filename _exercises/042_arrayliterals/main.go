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

	names = [3]string{"foo", "bar", "baz"}
	distances = [5]int{1, 4, 27, 64, 125}
	data = [5]uint8{0x2f, 0x11, 0x0, 0x2, 0x32}
	ratios = [1]float64{3.14}
	alives = [4]bool{false, true, true, false}
	zero = [0]uint8{}

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
