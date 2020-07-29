package main

import "fmt"

func main() {
	speed := 100
	force := 2.5
	// speed = speed * force
	// speed = speed * int(force)
	speed = int(float64(speed) * force)
	fmt.Println(speed)
	fmt.Println(force, int(force))

	var apple int
	var orange int32
	apple = int(orange)
	orange = int32(apple)
	fmt.Println(apple, orange)

	// isDelicious := bool(orange)
	orange = 65
	color := string(orange)
	fmt.Println(color)
}
