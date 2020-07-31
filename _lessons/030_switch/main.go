package main

import "fmt"

func main() {
	city := "Lyon"

	switch city {
	case "Tokyo":
		fmt.Println("Japan")
	default:
		fmt.Println("Where?")
	case "Paris", "Lyon":
		fmt.Println("France")
	}

	// switch true { // true by default
	// }

	i := 10
	switch {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}

	i = 142
	switch {
	case i > 100:
		fmt.Print("big ")
		fallthrough
	case i > 0:
		fmt.Print("positive ")
		fallthrough
	default:
		fmt.Println("number")
	}

	// short switch
	switch i := 10; {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}
