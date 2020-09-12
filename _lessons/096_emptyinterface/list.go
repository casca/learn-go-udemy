package main

import "fmt"

type printer interface {
	print()
	// discount(ratio float64)
}

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery. 🚛")
		return
	}

	for _, it := range l {
		// fmt.Printf("(%-10T) --> ", it)
		it.print()
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
		if it, ok := it.(discounter); ok {
			it.discount(ratio)
		}
	}
}
