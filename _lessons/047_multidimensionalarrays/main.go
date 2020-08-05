package main

import "fmt"

func main() {
	// m := [2][3]int{
	// 	[3]int{5, 6, 1},
	// 	[3]int{9, 8, 4},
	// }
	m := [...][3]int{
		{5, 6, 1},
		{9, 8, 4},
	}
	fmt.Printf("%#v\n", m)
}
