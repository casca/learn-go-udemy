package main

import "fmt"

type list []*product

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery. 🚛")
		return
	}

	for _, p := range l {
		// fmt.Printf("(%-10T) --> ", p)
		p.print()
	}
}

func (l list) discount(ratio float64) {
	type discounter interface{ discount(float64) }

	// for _, it := range l {
	// 	g, ok := it.(discounter)
	// 	// fmt.Printf("%T game? %v\n", it, ok)
	// 	if !ok {
	// 		continue
	// 	}
	// 	g.discount(ratio)
	// }

	for _, it := range l {
		// if it, ok := it.(discounter); ok {
		it.discount(ratio)
		// }
	}
}
