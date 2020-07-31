package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}

	scale := os.Args[1]
	var desc string
	switch scale {
	case "micro":
		desc = "less than 2.0"
	case "very minor":
		desc = "2.0-2.9"
	case "minor":
		desc = "3.0-3.9"
	case "light":
		desc = "4.0-4.9"
	case "moderate":
		desc = "5.0-5.9"
	case "strong":
		desc = "6.0-6.9"
	case "major":
		desc = "7.0-7.9"
	case "great":
		desc = "8.0-9.9"
	case "massive":
		desc = "10.0 or more"
	default:
		desc = "unknown"
	}
	fmt.Printf("%s richter scale is %s\n", scale, desc)
}
