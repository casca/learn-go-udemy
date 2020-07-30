package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.Itoa(42)
	fmt.Println(s)

	// n, err := strconv.Atoi(os.Args[1])
	n, err := strconv.Atoi("42")
	fmt.Println("Converted number    :", n)
	fmt.Println("Returned error value:", err)

	age := "21"
	n, err = strconv.Atoi(age)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("SUCCESS: Converted %q to %d.\n", age, n)
}
