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

	fmt.Println("sum:", 3+2)
	fmt.Println("sum:", 2+3.5)
	fmt.Println("dif:", 3-1)
	fmt.Println("dif:", 3-0.5)
	fmt.Println("prod:", 4*5)
	fmt.Println("prod:", 5*2.5)
	fmt.Println("quot:", 8/2)
	fmt.Println("quot:", 8/1.5)

	fmt.Println("rem:", 8%3)
	// fmt.Println("rem:", 8.0%3)
	f := 8.0
	fmt.Println("rem:", int(f)%3)

	var ratio2 float64 = 3 / 2 // ratio2 = float64(int(3)/int(2))
	fmt.Println(ratio2)
	fmt.Println("ratio3", 3./2) // or float(3)/2 or 3/2.0
}
