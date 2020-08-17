package main

import "fmt"

func main() {
	ages, oldCap := []int{1}, 1.

	for len(ages) < 5e5 {
		ages = append(ages, 1)

		c := cap(ages)
		if float64(c) != oldCap {
			fmt.Printf("len:%-10d cap:%-10g growth:%.2f\n", len(ages), c, float64(c)/oldCap)
		}
		oldCap = float64(c)
	}

}
