package main

import "fmt"

func main() {
	speed := 10
	fast := speed >= 80
	slow := speed < 20

	fmt.Printf("%T\n", fast)

	fmt.Printf("going fast? %t\n", fast)
	fmt.Printf("going slow? %t\n", slow)

	fmt.Printf("is it 100mph? %t\n", speed == 100)
	fmt.Printf("is it not 100mph? %t\n", speed != 100)

}
