package main

import "fmt"

func main() {
	blue := [3]int{6, 9, 3}
	red := blue
	blue[0] = 1
	fmt.Println(blue, red)
}
