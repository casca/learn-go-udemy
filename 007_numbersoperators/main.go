package main

import "fmt"

func main() {
	var (
		myAge   = 30
		yourAge = 35
		average float64
	)

	average = float64(myAge+yourAge) / 2
	fmt.Println(average)

	ratio := 1.0 / 10
	fmt.Printf("%.60f\n", ratio)
}
