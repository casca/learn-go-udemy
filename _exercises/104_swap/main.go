package main

import "fmt"

func main() {
	a, b := 1.0, 2.0
	swapVal(&a, &b)
	fmt.Printf("a=%v\tb=%v\n", a, b)

	pa, pb := &a, &b
	pa, pb = swapAddr(pa, pb)
	fmt.Printf("pa: %p — pb: %p\n", pa, pb)
	fmt.Printf("pa: %g — pb: %g\n", *pa, *pb)
}

func swapVal(f1, f2 *float64) {
	*f1, *f2 = *f2, *f1
}

func swapAddr(a, b *float64) (*float64, *float64) {
	return b, a
}
