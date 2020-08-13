package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := "2 4 6 1 3 5"
	ix := strings.Fields(numbers)

	fmt.Println("string", numbers)
	fmt.Println("slice", ix)
	fmt.Println("evens", ix[:3])
	fmt.Println("odds", ix[3:])
	fmt.Println("middle", ix[2:4])
	fmt.Println("first 2", ix[:2])
	fmt.Println("last 2", ix[len(ix)-2:])
	fmt.Println("evens last 1", ix[:3][2:])
	fmt.Println("odds last 2", ix[3:6][1:])

}
