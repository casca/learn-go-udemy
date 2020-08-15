package main

import "fmt"

func main() {
	var foo []int // cap = len = 0
	c1 := cap(foo)
	l1 := len(foo)

	ages := []int{35, 15, 25} // cap = len = 3
	c2 := cap(ages)
	l2 := len(ages)

	fmt.Println(c1, c2)
	fmt.Println(l1, l2)

	t := ages[0:0] // cap = 3, len = 0
	c3 := cap(t)
	l3 := len(t)
	fmt.Println(c3, l3)

	ages = ages[0:0]
	// ages = ages[0:3]
	ages = ages[0:cap(ages)] // cap = 3, len = 0
	// ages = ages[0 : cap(ages)+1]
	fmt.Println(ages)

	ages = ages[1:3] // cap = len = 2
}
