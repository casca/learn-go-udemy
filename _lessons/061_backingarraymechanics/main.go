package main

import "github.com/inancgumus/prettyslice"

func main() {
	prettyslice.PrintBacking = true

	var nums []int
	prettyslice.Show("no backing array", nums)

	nums = append(nums, 1, 3)
	prettyslice.Show("allocates", nums)

	nums = append(nums, 2)
	prettyslice.Show("free capacity", nums)

	nums = append(nums, 4)
	prettyslice.Show("no allocation", nums)

	nums = append(nums, nums[2:]...)
	prettyslice.Show("nums <- nums[2:]", nums)

	nums = append(nums[:2], 7, 9)
	prettyslice.Show("nums[:2] <- 7, 9", nums)

	nums = nums[:6]
	prettyslice.Show("extend", nums)
}
