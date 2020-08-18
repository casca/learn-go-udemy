package main

import (
	"fmt"

	"github.com/inancgumus/prettyslice"
)

func main() {
	sliceable := []byte{'f', 'u', 'l', 'l'}
	fmt.Println(sliceable[0:3:3])

	prettyslice.PrintBacking = true

	nums := []int{1, 3, 2, 4}
	odds := append(nums[:2:2], 5, 7)
	// odds = append(odds, 5, 7)
	evens := append(nums[2:4], 6, 8)

	prettyslice.Show("nums", nums)
	prettyslice.Show("odds", odds)
	prettyslice.Show("evens", evens)
}
