package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	// nums = append(nums, 4)
	// nums = append(nums, 9)
	// nums = append(nums, 4, 9)

	tens := []int{12, 13}
	nums = append(nums, tens...)
	fmt.Println(nums)
}
